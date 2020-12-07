import {Component, OnDestroy, OnInit, ViewChild} from '@angular/core';
import {Storage} from '@ionic/storage';
import * as Leaflet from 'leaflet';
import {LatLngLiteral, LeafletMouseEvent} from 'leaflet';
import {DataService} from '../../services/data.service';
import {IGpsCoordinate} from '../../building-objects-if';
import {Router} from '@angular/router';
import {NavigationModel} from '../navigationModel';
import {CampusViewModel} from '../campusViewModel';
import {DrawerState} from '../../../ionic-bottom-drawer/drawer-state';
import {SearchResultProvider} from '../../services/searchResultProvider';
import {IonContent, Platform} from '@ionic/angular';
import {IonicBottomDrawerComponent} from '../../../ionic-bottom-drawer/ionic-bottom-drawer.component';
import {Geolocation} from '@ionic-native/geolocation/ngx';
import {HttpErrorResponse} from '@angular/common/http';
import {GraphHopperRoute, OpenStreetMapBounds, OpenStreetMapService} from '../../services/osm/open-street-map.service';
import {SearchInputComponent} from '../search-input/search-input.component';
import {NavigationInstructionSlidesComponent} from '../navigation-instruction-slides/navigation-instruction-slides.component';
import {INavigationInstruction} from '../navigation-instruction-slides/navigation-instruction-if';
import {LastOpenStreetMapCenterPersistenceService} from '../../services/LastOpenStreetMapCenterPersistence.service';
import {Plugins} from '@capacitor/core';

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
      private _dataService: DataService,
      private router: Router,
      public model: NavigationModel,
      private storage: Storage,
      private dataService: DataService,
      private geolocation: Geolocation,
      private openStreetMapService: OpenStreetMapService,
      private lastOpenStreetMapCenterPersistence: LastOpenStreetMapCenterPersistenceService,
      private platform: Platform
  ) {
     this.isHybridPlatform = this.platform.is('hybrid');
  }

  private readonly isHybridPlatform: boolean;
  map: Leaflet.Map;
  private searchMarker: Leaflet.Marker[] = [];
  private routes: Leaflet.Polyline[] = [];
  public availableCampus: CampusViewModel[] = [];
  public progressIsVisible = false;
  @ViewChild('drawerContent') drawerContent : IonContent;
  @ViewChild('searchDrawer') searchDrawer : IonicBottomDrawerComponent;
  @ViewChild('locationDrawer') locationDrawer : IonicBottomDrawerComponent;
  @ViewChild('routeDrawer') routeDrawer : IonicBottomDrawerComponent;
  @ViewChild('inNavigationDrawer') inNavigationDrawer : IonicBottomDrawerComponent;

  @ViewChild('searchInput') searchInput : SearchInputComponent
  @ViewChild('navSlides') navSlides : NavigationInstructionSlidesComponent;
  errorMessage: string;
  private currentPositionMarker: Leaflet.Marker = null;
  private isInitialized = false;

  private readonly DEFAULT_ZOOM = 17;
  private readonly MAX_ZOOM = 18;
  private readonly MIN_ZOOM = 14;

  private static convertToLeafLetCoordinates(body: IGpsCoordinate[]) {
    const leafletBody:LatLngLiteral[] = []
    for (const coordinate of body){
      leafletBody.push({lat: coordinate.Latitude, lng: coordinate.Longitude});
    }

    return leafletBody;
  }

  async ngOnInit() {

    if (!this.model.recentSearches || this.model.recentSearches.length === 0) {
      const searches = await SearchResultProvider.readRecentSearch(this.storage);
      if (searches !== null) {
        this.model.recentSearches = searches;
        console.log(this.model.recentSearches);
      }
    }

    if (!this.model.availableCampus || this.model.availableCampus.length === 0) {
      this.model.availableCampus = await this.dataService.get_campus_search().toPromise()
    }

    if (this.model.availableCampus) {
      for (const campus of this.model.availableCampus) {
        this.availableCampus.push(new CampusViewModel(campus))
      }
    }

    await Promise.all([
        this.searchDrawer.SetState(IonicBottomDrawerComponent.GetRecommendedDrawerStateForDevice()),
        this.locationDrawer.SetState(DrawerState.Hidden),
        this.routeDrawer.SetState(DrawerState.Hidden),
        this.inNavigationDrawer.SetState(DrawerState.Hidden)]);
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
  }

  private async initializeMap(router: Router) {
    const osmBounds = await this.openStreetMapService.GetBounds()
    const southWest = Leaflet.latLng(osmBounds.SouthWest.Lat, osmBounds.SouthWest.Lng);
    const northEast = Leaflet.latLng(osmBounds.NorthEast.Lat, osmBounds.NorthEast.Lng);

    const bounds = Leaflet.latLngBounds(southWest, northEast);
    console.log(osmBounds, bounds);

    // maxZoom for leaflet map is 18
    this.map = Leaflet.map('leafletMap', {
      maxBounds:bounds,
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

    const buildings = await this._dataService.get_buildings_search().toPromise();

    function onPolygonClick(event:LeafletMouseEvent) {
      router.navigate(['tabs/navigation/detail'],
          {queryParams: {building: event.target.options.className}})
          .then(() => console.log(event.latlng, event.target.options.className));
    }

    if (buildings !== null) {
      for (const building of buildings) {
        if (building.Body !== null && building.Body.length > 0) {
          Leaflet.polygon(MapPageComponent.convertToLeafLetCoordinates(building.Body),
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

  async onSearch(searchInput: string) {
    this.model.errorMessage = '';
    this.progressIsVisible = true;

    try {
      this.clearSearchMarkers();

      // look for indexed values
      // location
      try {
        const location = await this.dataService.get_location(searchInput).toPromise();
          if (location) {
            const locationBuilding = await this.dataService.get_building(location.Building).toPromise();
            const coordinates = this.getCenterCoordinateFromBody(locationBuilding.Body);
            this.model.SetLocationAsSearchResultObject(location, {lat: coordinates.Latitude, lng: coordinates.Longitude});
            this.searchMarker.push(
                this.showMarker(coordinates.Latitude, coordinates.Longitude, 'Room ' + location.Name, true));

            this.map.flyTo(this.model.latestSearchResult.LatLng, this.DEFAULT_ZOOM);

            await this.showElementDrawer();
            SearchResultProvider.addRecentSearch(searchInput, this.model, this.storage);
            return;
          }
        } catch (e) {
          console.log(e);
      }
      // building
      try {
        const building = await this.dataService.get_building(searchInput, false).toPromise();
        if (building) {
          const coordinates = this.getCenterCoordinateFromBody(building.Body);
          this.model.SetBuildingAsSearchResultObject(building, {lat: coordinates.Latitude, lng: coordinates.Longitude});
          this.map.flyTo(this.model.latestSearchResult.LatLng, this.DEFAULT_ZOOM);
          this.searchMarker.push(
              this.showMarker(coordinates.Latitude, coordinates.Longitude, 'Building ' + building.Name, true));

          await this.showElementDrawer();
          SearchResultProvider.addRecentSearch(searchInput, this.model, this.storage);
          return;
        }
      } catch (e) {
        console.log(e);
      }
      // campus
      try {
        const campus = await this.dataService.get_campus(searchInput, false).toPromise();
        if (campus) {
          this.model.SetCampusAsSearchResultObject(campus);
          this.searchMarker.push(
              this.showMarker(campus.Latitude, campus.Longitude, 'Campus ' + campus.Name, true));

          this.map.flyTo([campus.Latitude, campus.Longitude], this.DEFAULT_ZOOM);
          await this.showElementDrawer();
          SearchResultProvider.addRecentSearch(searchInput, this.model, this.storage);
          return;
        }
      } catch (e) {
        console.log(e);
      }

      // Look for all possible buildings via Search-API
      const buildings = await this.dataService.get_buildings_search(searchInput).toPromise();
      if (buildings !== null && buildings.length > 0) {
        // found building
        SearchResultProvider.addRecentSearch(searchInput, this.model, this.storage);
        for (const buld of buildings) {
          const coordinates = this.getCenterCoordinateFromBody(buld.Body);
          this.searchMarker.push(
            this.showMarker(coordinates.Latitude, coordinates.Longitude, buld.Name, false));
        }
      }
      else {
        // look for campus on Search-API if no building is found
        const campusArray = await this.dataService.get_campus_search(searchInput).toPromise();
        if (campusArray !== null && campusArray.length > 0) {
          // found building
          SearchResultProvider.addRecentSearch(searchInput, this.model, this.storage);
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

  onRoute(event: string[]) {
  }

  recentSearchClick(s: string) {

  }

  public onDrawerStateChange(state:DrawerState) {
    // in case the view is not initialized
    if (this.drawerContent === undefined) {
      return;
    }

    this.drawerContent.scrollY = state === DrawerState.Top;
  }

  public async onCloseLocationDrawer(event:any) {
    this.searchInput.clearDestinationInput();
    await this.locationDrawer.SetState(DrawerState.Hidden);
    await this.searchDrawer.SetState(IonicBottomDrawerComponent.GetRecommendedDrawerStateForDevice());
    this.clearSearchMarkers();
  }

  public async onCloseRouteDrawer(event:any) {
    this.clearRoutes();
    await this.routeDrawer.SetState(DrawerState.Hidden);
    await this.locationDrawer.SetState(IonicBottomDrawerComponent.GetRecommendedDrawerStateForDevice());
    this.map.flyTo(this.model.latestSearchResult.LatLng, this.DEFAULT_ZOOM);
  }

  public async showElementDrawer() {
    if (this.isHybridPlatform) {
        await Keyboard.hide();
    }

    await this.searchDrawer.SetState(DrawerState.Hidden);
    await this.locationDrawer.SetState(IonicBottomDrawerComponent.GetRecommendedDrawerStateForDevice());
  }

  public onCampusClick(c:CampusViewModel) {
    this.map.flyTo(c.LatLng, this.DEFAULT_ZOOM);
  }

  public async onRouteBtnClick() {
    const position = await this.geolocation.getCurrentPosition();
    console.log(position);
    const route:GraphHopperRoute = await this.openStreetMapService.GetRoute(
        {lat: position.coords.latitude, lng: position.coords.longitude},
        this.model.latestSearchResult.LatLng);

    console.log(route);
    this.model.SetGraphHopperRouteAsRoute(route);

    await this.locationDrawer.SetState(DrawerState.Hidden);
    await this.routeDrawer.SetState(IonicBottomDrawerComponent.GetRecommendedDrawerStateForDevice());

    const polyline = Leaflet.polyline(this.model.Route.Coordinates, {color: 'red'}).addTo(this.map);
    this.map.flyTo(polyline.getCenter(), this.DEFAULT_ZOOM);
    await this.map.fitBounds(polyline.getBounds());
    this.routes.push(polyline);
  }

  public async onLaunchNavigation() {
    await this.routeDrawer.SetState(DrawerState.Hidden);

    this.navSlides.instructions = this.model.Route.NavigationInstructions;
    this.navSlides.show();

    await this.inNavigationDrawer.SetState(IonicBottomDrawerComponent.GetRecommendedDrawerStateForDevice());

    this.map.flyTo(this.model.Route.Coordinates[this.model.Route.NavigationInstructions[0].interval[0]], this.MAX_ZOOM);
  }

  public onSlideChange(index:number) {
    this.onNavigationInstructionClick(this.model.Route.NavigationInstructions[index]);
  }

  public async onEndRouteClick() {
    await this.inNavigationDrawer.SetState(DrawerState.Hidden);
    await this.locationDrawer.SetState(IonicBottomDrawerComponent.GetRecommendedDrawerStateForDevice());
    this.clearRoutes();
    this.map.flyTo(this.model.latestSearchResult.LatLng, this.DEFAULT_ZOOM);
  }

  async detailsBtnClick() {
      await this.router.navigate(['tabs/navigation/detail'],
          {queryParams: this.routes.length === 0
                ? this.model.latestSearchResult.DetailRouterParams : this.model.latestSearchResult.RouteRouterParams});
  }

  public onNavigationInstructionClick(instruction:INavigationInstruction) {
    this.map.flyTo(this.model.Route.Coordinates[instruction.interval[0]], this.MAX_ZOOM);
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
    addTo(this.map)

    return popupText ?
        (showPopUp ? marker.bindPopup(popupText).openPopup() : marker.bindPopup(popupText))
        : marker;
  }

  private getCenterCoordinateFromBody(body: IGpsCoordinate[]) {
    const leafletLatLng = []
    for (const c of body) {
      leafletLatLng.push([c.Latitude, c.Longitude]);
    }

    // cannot use getCenter if polygon is not added to layer (map)
    const p = new Leaflet.Polygon(leafletLatLng).addTo(this.map);

    const ret:IGpsCoordinate = {
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
    this.navSlides.hide();
  }

  public UseDrawerForNavigation() :boolean {
    return !(IonicBottomDrawerComponent.GetRecommendedDrawerStateForDevice() === DrawerState.Bottom);
  }

  private delay(ms: number) {
    return new Promise( resolve => setTimeout(resolve, ms) );
  }

  async onSearchFocus($event: string) {
    if (this.platform.is('hybrid')) {
      await this.searchDrawer.SetState(DrawerState.Top)
    }
  }
}
