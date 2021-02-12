import {Component, OnDestroy, OnInit, ViewChild} from '@angular/core';
import {Storage} from '@ionic/storage';
import * as Leaflet from 'leaflet';
import {LatLngLiteral, LeafletMouseEvent} from 'leaflet';
import {DataService} from '../../services/data.service';
import {IGpsCoordinate} from '../../building-objects-if';
import {ActivatedRoute, Router} from '@angular/router';
import {IRouteLocation, NavigationModel} from '../navigationModel';
import {CampusViewModel} from '../campusViewModel';
import {Platform} from '@ionic/angular';
import {Geolocation} from '@ionic-native/geolocation/ngx';
import {HttpErrorResponse} from '@angular/common/http';
import {OpenStreetMapService, OsmRoute} from '../../services/osm/open-street-map.service';
import {SearchInputComponent} from '../search-input/search-input.component';
import {NavigationInstructionSlidesComponent} from '../navigation-instruction-slides/navigation-instruction-slides.component';
import {INavigationInstruction} from '../navigation-instruction-slides/navigation-instruction-if';
import {LastOpenStreetMapCenterPersistenceService} from '../../services/LastOpenStreetMapCenterPersistence.service';
import {Plugins} from '@capacitor/core';
import {NavDrawerManagerComponent, NavDrawerState} from '../nav-drawer-manager/nav-drawer-manager.component';

const { Keyboard } = Plugins;


const iconRetinaUrl = 'leaflet/marker-icon-2x.png';
const iconUrl = 'leaflet/marker-icon.png';
const shadowUrl = 'leaflet/marker-shadow.png';
const iconDefault = Leaflet.icon({
  iconRetinaUrl,
  iconUrl,
  shadowUrl,
  iconSize: [25, 41],
  iconAnchor: [12, 41],
  popupAnchor: [1, -34],
  tooltipAnchor: [16, -28],
  shadowSize: [41, 41]
});
Leaflet.Marker.prototype.options.icon = iconDefault;

@Component({
  selector: 'app-map-page',
  templateUrl: './map-page.component.html',
  styleUrls: ['./map-page.component.scss'],
})
export class MapPageComponent implements OnInit, OnDestroy {

  constructor(
      private dataService: DataService,
      private  route: ActivatedRoute,
      private router: Router,
      public model: NavigationModel,
      private storage: Storage,
      private geolocation: Geolocation,
      private openStreetMapService: OpenStreetMapService,
      private lastOpenStreetMapCenterPersistence: LastOpenStreetMapCenterPersistenceService,
      private platform: Platform
  ) {
     this.isHybridPlatform = this.platform.is('hybrid');
  }

  public static readonly MyLocation = 'My Location';

  private readonly isHybridPlatform: boolean;
  map: Leaflet.Map;
  private searchMarker: Leaflet.Marker[] = [];
  private routes: Leaflet.Polyline[] = [];
  public progressIsVisible = false;

  @ViewChild('searchInput') searchInput: SearchInputComponent;
  @ViewChild('navSlides') navSlides: NavigationInstructionSlidesComponent;
  @ViewChild('drawerManager') drawerManager: NavDrawerManagerComponent;
  errorMessage: string;
  private currentPositionMarker: Leaflet.Marker = null;
  private isInitialized = false;
  private isSubscribed = false;

  private readonly DEFAULT_ZOOM = 17;
  private readonly MAX_ZOOM = 18;
  private readonly MIN_ZOOM = 14;

  private static convertToLeafLetCoordinates(body: IGpsCoordinate[]) {
    const leafletBody: LatLngLiteral[] = [];
    for (const coordinate of body){
      leafletBody.push({lat: coordinate.Latitude, lng: coordinate.Longitude});
    }

    return leafletBody;
  }

  async ngOnInit() {

    if (!this.model.availableCampus || this.model.availableCampus.length === 0) {
      const campus = await this.dataService.get_campus_search().toPromise();
      for (const c of campus) {
        this.model.availableCampus.push(new CampusViewModel(c));
      }
    }
  }

  async ionViewDidEnter() {

    if (!this.isInitialized) {
      await this.initializeMap(this.router);
      this.isInitialized = true;
    }

    try {
      const position = await this.geolocation.getCurrentPosition();
      if (this.currentPositionMarker === null){
        this.currentPositionMarker = Leaflet.marker([position.coords.latitude, position.coords.longitude]).
        addTo(this.map).
        bindPopup('Position');
      } else {
        this.currentPositionMarker.setLatLng([position.coords.latitude, position.coords.longitude]);
      }
    }
    catch (error) {
      console.log('Error getting location', error);
    }

    if (this.isSubscribed === false) {
      this.isSubscribed = true;
      this.route.queryParams.subscribe(async params => {

        if (params === null || params === undefined) {
          return;
        }

        if (params.location !== null && params.location !== undefined) {
          await this.execSearch(params.location); // this will show also the location drawer
          this.clearRoutes();
        } else if (params.destination !== null && params.destination !== undefined) {
          this.drawerManager.SetState(NavDrawerState.RouteView);
          await this.execDirectSearch(params.destination);
          await this.routeToLatestSearchResult();
        } else {
          // assuming initial state
          await this.drawerManager.SetState(NavDrawerState.SearchView);
          this.clearSearchMarkers();
          this.clearRoutes();
        }

      });
    }
  }

  private async initializeMap(router: Router) {
    const osmBounds = await this.openStreetMapService.GetBounds();
    const southWest = Leaflet.latLng(osmBounds.SouthWest.Lat, osmBounds.SouthWest.Lng);
    const northEast = Leaflet.latLng(osmBounds.NorthEast.Lat, osmBounds.NorthEast.Lng);

    const bounds = Leaflet.latLngBounds(southWest, northEast);
    console.log(osmBounds, bounds);

    // maxZoom for leaflet map is 18
    this.map = Leaflet.map('leafletMap', {
      maxBounds: bounds,
      maxZoom: this.MAX_ZOOM,
      minZoom: this.MIN_ZOOM,
      zoomControl: !this.platform.is('hybrid')
    });

    this.map.on('moveend', event => {
      LastOpenStreetMapCenterPersistenceService.persist(this.storage, {
        center: event.target.getCenter(),
        zoom: event.target.getZoom()
      });
    });

    await this.lastOpenStreetMapCenterPersistence.load(this.storage, this.map, this.DEFAULT_ZOOM);

    Leaflet.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
      attribution: 'edupala.com © Angular LeafLet',
    }).addTo(this.map);

    if (this.platform.is('hybrid')) {
      await this.delay(500);
      this.map.invalidateSize();
    }

    const buildings = await this.dataService.get_buildings_search().toPromise();

    function onPolygonClick(event: LeafletMouseEvent) {

      router.navigate(['tabs/navigation'],
          {queryParams: {location: event.target.options.className}})
          .then(() => console.log(event.latlng, event.target.options.className));
    }

    if (buildings !== null) {
      for (const building of buildings) {
        if (building.edges.Body !== null && building.edges.Body.length > 0) {
          Leaflet.polygon(MapPageComponent.convertToLeafLetCoordinates(building.edges.Body),
              {className: building.Name, color: building.Color ?? '#0083C6'})
              .addTo(this.map)
              .on('click', onPolygonClick);
        }
      }
    }
    // this.showMarker(49.452858, 11.093235, 'Technische Hochschule Nürnberg Georg Simon Ohm', true);
  }

  /** Remove map when we have multiple map object */
  ngOnDestroy() {
    this.map.remove();
  }

  public async onSearch(searchInput: string) {
    await this.router.navigate(['/tabs/navigation'], {queryParams: {location: searchInput}});
  }

  private async execSearch(searchInput: string) {
    this.model.errorMessage = '';
    this.progressIsVisible = true;

    try {
      this.clearSearchMarkers();

      if (await this.execDirectSearch(searchInput)) {
        this.searchMarker.push(
            this.showMarker(this.model.latestSearchResult.LatLng.lat, this.model.latestSearchResult.LatLng.lng,
                this.model.latestSearchResult.Name, true)
        );
        this.map.flyTo(this.model.latestSearchResult.LatLng, this.DEFAULT_ZOOM);
        await this.showElementDrawer();
        return;
      }

      // Look for all possible buildings via Search-API
      const buildings = await this.dataService.get_buildings_search(searchInput).toPromise();
      if (buildings !== null && buildings.length > 0) {
        // found building
        // await this.model.addRecentSearch(searchInput);
        for (const building of buildings) {
          const coordinates = this.getCenterCoordinateFromBody(building.edges.Body);
          this.searchMarker.push(
            this.showMarker(coordinates.Latitude, coordinates.Longitude, building.Name, false));
        }
      }
      else {
        // look for campus on Search-API if no building is found
        const campusArray = await this.dataService.get_campus_search(searchInput).toPromise();
        if (campusArray !== null && campusArray.length > 0) {
          // found building
          // await this.model.addRecentSearch(searchInput);
          for (const camp of campusArray) {
            this.searchMarker.push(
              this.showMarker(camp.Latitude, camp.Longitude, camp.Name, false));
          }
        }
      }
    } catch (ex) {
      this.handleInputError(ex, searchInput);
    } finally {
      this.progressIsVisible = false;
    }
  }

  private async execDirectSearch(searchInput: string): Promise<boolean> {
    // look for indexed values
    // location
    try {
      const location = await this.dataService.get_location(searchInput).toPromise();
      if (location) {
        const locationBuilding = await this.dataService.get_building(location.Building).toPromise();
        const coordinates = this.getCenterCoordinateFromBody(locationBuilding.edges.Body);
        await this.model.addRecentSearchLocation(location, {lat: coordinates.Latitude, lng: coordinates.Longitude}, locationBuilding);
        return true;
      }
    } catch (e) {
      console.log(e);
    }
    // building
    try {
      const building = await this.dataService.get_building(searchInput, false).toPromise();
      if (building) {
        const coordinates = this.getCenterCoordinateFromBody(building.edges.Body);
        await this.model.addRecentSearchBuilding(building, {lat: coordinates.Latitude, lng: coordinates.Longitude});
        return true;
      }
    } catch (e) {
      console.log(e);
    }
    // campus
    try {
      const campus = await this.dataService.get_campus(searchInput, false).toPromise();
      if (campus) {
        await this.model.addRecentSearchCampus(campus);
        return true;
      }
    } catch (e) {
      console.log(e);
    }

    return false;
  }

  // public async onRoute(event: IRouteLocation[]) {
  //   await this.router.navigate(['/tabs/navigation'], {queryParams: {destination: event}});
  // }

  public async onRoute(event: IRouteLocation[]) {
    console.log(event);
    this.clearRoutes();
    await this.drawerManager.SetState(NavDrawerState.RouteView, false);

    const routes: OsmRoute[] = await this.openStreetMapService.GetRoute(
        event[0].LatLng, event[1].LatLng);

    this.model.SetOsmRouteAsRoute(routes[0], event[0], event[1]);


    const polyline = Leaflet.polyline(this.model.Route.Coordinates, {color: 'red'}).addTo(this.map);
    this.map.flyTo(polyline.getCenter(), this.DEFAULT_ZOOM);
    await this.map.fitBounds(polyline.getBounds());
    this.routes.push(polyline);
  }

  public async showElementDrawer() {
    if (this.isHybridPlatform) {
        await Keyboard.hide();
    }
    await this.drawerManager.SetState(NavDrawerState.LocationView);
  }

  public onCampusClick(c: CampusViewModel) {
    this.map.flyTo(c.LatLng, this.DEFAULT_ZOOM);
  }

  public async onDrawerManagerStateChange(newState: NavDrawerState) {
    console.log(newState);
    switch (newState) {
      case NavDrawerState.SearchView:
        await this.router.navigate(['/tabs/navigation']);
        break;
      case NavDrawerState.LocationView:
        await this.router.navigate(['/tabs/navigation'], {queryParams: {location: this.model.latestSearchResult.Name}});
        // this.clearRoutes();
        // this.map.flyTo(this.model.latestSearchResult.LatLng, this.DEFAULT_ZOOM);
        break;
      case NavDrawerState.RouteView:
        await this.router.navigate(['/tabs/navigation'], {queryParams: {destination: this.model.latestSearchResult.Name}});
        break;
      case NavDrawerState.InNavigationView:
        await this.onLaunchNavigation();
        break;
    }
  }

  public async onDetailsClick() {
    await this.router.navigate(['tabs/navigation/detail'],
        {queryParams: this.routes.length === 0
              ? this.model.latestSearchResult.DetailRouterParams : this.model.latestSearchResult.RouteRouterParams});
  }

  private async routeToLatestSearchResult() {
    const geoPosition = await this.geolocation.getCurrentPosition();
    const position = {lat: geoPosition.coords.latitude, lng: geoPosition.coords.longitude};
    const routes: OsmRoute[] = await this.openStreetMapService.GetRoute(
        position, this.model.latestSearchResult.LatLng);

    this.model.SetOsmRouteAsRoute(routes[0], {Name: MapPageComponent.MyLocation, LatLng: position}, this.model.latestSearchResult);
    // await this.drawerManager.setState(NavDrawerState.RouteView);

    const polyline = Leaflet.polyline(this.model.Route.Coordinates, {color: 'red'}).addTo(this.map);
    this.map.flyTo(polyline.getCenter(), this.DEFAULT_ZOOM);
    await this.map.fitBounds(polyline.getBounds());
    this.routes.push(polyline);
  }

  public async onLaunchNavigation() {

    this.navSlides.instructions = this.model.Route.NavigationInstructions;
    await this.navSlides.show();
    // await this.drawerManager.setState(NavDrawerState.InNavigationView);
    this.map.flyTo(this.model.Route.Coordinates[this.model.Route.NavigationInstructions[0].Interval[0]], this.MAX_ZOOM);
  }

  public onSlideChange(index: number) {
    this.onNavigationInstructionClick(this.model.Route.NavigationInstructions[index]);
  }

  public onNavigationInstructionClick(instruction: INavigationInstruction) {
    this.map.flyTo(this.model.Route.Coordinates[instruction.Interval[0]], this.MAX_ZOOM);
  }

  private handleInputError(ex, searchInput: string) {
    if (ex instanceof HttpErrorResponse) {
      const httpError = ex as HttpErrorResponse;
      if (httpError.status === 400) {
        this.errorMessage = 'Studi-Guide can\'t find ' + searchInput;
      } else {
        this.errorMessage = httpError.message;
      }
    } else {
      this.errorMessage = (ex as Error).message;
    }
  }

  private showMarker(lat: number, long: number, popupText: string  = '', showPopUp: boolean) {
    const marker = Leaflet.marker([lat, long]).
    addTo(this.map);

    return popupText ?
        (showPopUp ? marker.bindPopup(popupText).openPopup() : marker.bindPopup(popupText))
        : marker;
  }

  private getCenterCoordinateFromBody(body: IGpsCoordinate[]) {
    const leafletLatLng = [];
    for (const c of body) {
      leafletLatLng.push([c.Latitude, c.Longitude]);
    }

    // cannot use getCenter if polygon is not added to layer (map)
    const p = new Leaflet.Polygon(leafletLatLng).addTo(this.map);

    const ret: IGpsCoordinate = {
      Longitude: p.getCenter().lng,
      Latitude: p.getCenter().lat
    };

    p.remove();

    return ret;
  }

  private clearSearchMarkers() {
    for (const marker of this.searchMarker) {
      this.map.removeLayer(marker);
    }
  }

  private clearRoutes() {
    for (const p of this.routes) {
      p.remove();
    }
    this.routes = [];
    this.model.ClearRoute();
    this.navSlides.hide();
  }

  private delay(ms: number) {
    return new Promise( resolve => setTimeout(resolve, ms) );
  }

}
