import {Component, OnInit, OnDestroy, ViewChild, AfterViewInit} from '@angular/core';
import {Storage} from '@ionic/storage';
import * as Leaflet from 'leaflet';
import {LatLng, latLng, LatLngLiteral, LeafletMouseEvent} from 'leaflet';
import {DataService} from '../../services/data.service';
import {IGpsCoordinate, ILocation} from '../../building-objects-if';
import {Router} from '@angular/router';
import {NavigationModel} from '../navigationModel';
import {CampusViewModel} from '../campusViewModel';
import {DrawerState} from '../../../ionic-bottom-drawer/drawer-state';
import {SearchResultProvider} from '../../services/searchResultProvider';
import {IonContent} from '@ionic/angular';
import {IonicBottomDrawerComponent} from '../../../ionic-bottom-drawer/ionic-bottom-drawer.component';
import { Geolocation } from '@ionic-native/geolocation/ngx';
import {HttpErrorResponse} from '@angular/common/http';
import {GraphHopperService} from '../../services/graph-hopper/graph-hopper.service';

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
export class MapPageComponent implements OnInit, OnDestroy, AfterViewInit {
  map: Leaflet.Map;
  private searchMarker: Leaflet.Marker[] = [];
  public availableCampus: CampusViewModel[] = [];
  public progressIsVisible = false;
  public selectedItem: {Description:string, Name:string, Type: string} = { Description: '', Name: '', Type:''};
  @ViewChild('drawerContent') drawerContent : IonContent;
  @ViewChild('searchDrawer') searchDrawer : IonicBottomDrawerComponent;
  @ViewChild('locationDrawer') locationDrawer : IonicBottomDrawerComponent;
  errorMessage: string;
  private currentPositionMarker: Leaflet.Marker = null;
  private isInitialized = false;

  constructor(
      private _dataService: DataService,
      private router: Router,
      public model: NavigationModel,
      private storage: Storage,
      private dataService: DataService,
      private geolocation: Geolocation,
      private ghService: GraphHopperService) {
  }

  async ngAfterViewInit() {
    await this.locationDrawer.SetState(DrawerState.Hidden);
  }

  async ngOnInit() {
    if (this.model.recentSearches === null) {
      const searches = await SearchResultProvider.readRecentSearch(this.storage);
      if (searches !== null) {
        this.model.recentSearches = searches;
        console.log(this.model.recentSearches);
      }
    }

    if (this.model.availableCampus.length === 0)
    {
      this.model.availableCampus = await this.dataService.get_campus_search().toPromise()
    }

    for (const campus of this.model.availableCampus) {
      this.availableCampus.push(new CampusViewModel(campus))
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
  }

  private async initializeMap(router: Router) {
    const southWest = Leaflet.latLng(49.4126, 11.0111);
    const northEast = Leaflet.latLng(49.5118, 11.2167);
    const bounds = Leaflet.latLngBounds(southWest, northEast);

    // maxZoom for leaflet map is 18
    this.map = Leaflet.map('leafletMap', {
      maxBounds:bounds,
      maxZoom: 18,
      minZoom: 14
    })
        .setView([49.452368, 11.093299], 17);

    Leaflet.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
      attribution: 'edupala.com © Angular LeafLet',
    }).addTo(this.map);


    const buildings = await this._dataService.get_buildings_search().toPromise();

    function onPolygonClick(event:LeafletMouseEvent) {
      router.navigate(['tabs/navigation/detail'],
          {queryParams: {building: event.target.options.className}})
          .then(r => console.log(event.latlng, event.target.options.className));
    }

    if (buildings !== null) {
      for (const building of buildings) {
        if (building.Body !== null && building.Body.length > 0) {
          Leaflet.polygon(this.convertToLeafLetCoordinates(building.Body),
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
      try {
        const building = await this.dataService.get_building(searchInput, false).toPromise();
        if (building) {
          const coordinates = this.getCoordinateFromBody(building.Body);
          this.searchMarker.push(
              this.showMarker(coordinates.Latitude, coordinates.Longitude, 'Building ' + building.Name, true));

          this.map.setView([coordinates.Latitude, coordinates.Longitude], 17)

          await this.showElementDrawer({Name: building.Name, Description: 'Campus:' + building.Campus, Type: 'Building'});
          return;
        }
      } catch (e) {
        console.log(e);
      }

      try {
        const campus = await this.dataService.get_campus(searchInput, false).toPromise();
        if (campus) {
          this.searchMarker.push(
              this.showMarker(campus.Latitude, campus.Longitude, 'Campus ' + campus.Name, true));

          this.map.setView([campus.Latitude, campus.Longitude], 17)
          await this.showElementDrawer({Name: campus.Name, Description: campus.ShortName, Type: 'Campus'});
          return;
        }
      } catch (e) {
        console.log(e);
      }

      // Look for all possible buildings via Search-API
      const buildings = await this.dataService.get_buildings_search(searchInput).toPromise();
      if (buildings !== null && buildings.length > 0) {

          // found building
          for (const buld of buildings) {
            const coordinates = this.getCoordinateFromBody(buld.Body);
            this.searchMarker.push(
              this.showMarker(coordinates.Latitude, coordinates.Longitude, buld.Name, false));
          }
      }
      else {
        // look for campus on Search-API if no building is found
        const campusArray = await this.dataService.get_campus_search(searchInput).toPromise();
        if (campusArray !== null && campusArray.length > 0) {

          // found building
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

  onRoute($event: string[]) {
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
    await this.locationDrawer.SetState(DrawerState.Hidden);
    await this.searchDrawer.SetState(DrawerState.Docked);
  }

  public async showElementDrawer(location: { Description:string, Name:string, Type:string }) {
    await this.locationDrawer.SetState(DrawerState.Hidden);
    this.selectedItem = location;
    await this.searchDrawer.SetState(DrawerState.Hidden);
    await this.locationDrawer.SetState(DrawerState.Docked);
  }

  navigationBtnClick() {
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

  private getCoordinateFromBody(body: IGpsCoordinate[]) {
    let totalLat = 0, totalLong = 0;
    for (const coordinate of body)
    {
      totalLat += coordinate.Latitude;
      totalLong += coordinate.Longitude;
    }
    const centerLat = totalLat / body.length;
    const centerLong = totalLong / body.length;

    return {
      Longitude: centerLong,
      Latitude: centerLat
    }
  }

  private clearSearchMarkers() {
    for (const marker of this.searchMarker) {
      this.map.removeLayer(marker);
    }
  }

  private convertToLeafLetCoordinates(body: IGpsCoordinate[]) {
    const leafletBody:LatLngLiteral[] = []
    for (const coordinate of body){
      leafletBody.push({lat: coordinate.Latitude, lng: coordinate.Longitude});
    }

    return leafletBody;
  }

  // TODO remove this method
  public async TestGetRouteEndPoint() {
    console.log('Hello World');
    const position = await this.geolocation.getCurrentPosition();
    console.log(position);
    const route = await this.ghService.GetRouteEndpoint(
        {lat: position.coords.latitude, lng: position.coords.longitude},
        {lat: 49.45281, lng: 11.09347});
    console.log(route);
    //console.log(route.paths[0].points.coordinates);

    const leafletLatLng = [];
    for(const coordinate of route.paths[0].points.coordinates) {
      console.log(coordinate[0], coordinate[1]);
      leafletLatLng.push([coordinate[1], coordinate[0]]);
    }
    console.log(leafletLatLng);
    const polyline = Leaflet.polyline(leafletLatLng, {color: 'red'}).addTo(this.map);

  }
}
