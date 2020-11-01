import {Component, OnInit, OnDestroy} from '@angular/core';
import * as Leaflet from 'leaflet';
import {LatLngExpression, LatLngLiteral, LeafletMouseEvent} from 'leaflet';
import {DataService} from '../../services/data.service';
import {IGpsCoordinate} from '../../building-objects-if';
import {Router} from '@angular/router';

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
  map: Leaflet.Map;
  private isInitialized = false;

  constructor(
      private _dataService: DataService,
      private router: Router) {
  }

  ngOnInit() { }
  async ionViewDidEnter() {
    if (!this.isInitialized) {
      await this.initializeMap(this.router);
      this.isInitialized = true;
    }
  }

  private async initializeMap(router: Router) {
    const southWest = Leaflet.latLng(49.4126, 11.0111);
    const northEast = Leaflet.latLng(49.5118, 11.2167);
    const bounds = Leaflet.latLngBounds(southWest, northEast);

    // maxZoom for leaflet map is 18
    this.map = Leaflet.map('mapId', {
      maxBounds:bounds,
      maxZoom: 18,
      minZoom: 14
    }).setView([49.452858, 11.093235], 17);

    Leaflet.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
      attribution: 'edupala.com © Angular LeafLet',
    }).addTo(this.map);


    const buildings = await this._dataService.get_buildings().toPromise();

    function onPolygonClick(event:LeafletMouseEvent) {
      router.navigate(['tabs/navigation/detail'], {queryParams: {building: event.target.options.className}})
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

     Leaflet.marker([49.452858, 11.093235]).
      addTo(this.map).
      bindPopup('Technische Hochschule Nürnberg Georg Simon Ohm').
      openPopup();
  }

  /** Remove map when we have multiple map object */
  ngOnDestroy() {
    this.map.remove();
  }

  convertToLeafLetCoordinates(body: IGpsCoordinate[]) {
    const leafletBody:LatLngLiteral[] = []
    for (const coordinate of body){
      leafletBody.push({lat: coordinate.Latitude, lng: coordinate.Longitude});
    }

    return leafletBody;
  }
}
