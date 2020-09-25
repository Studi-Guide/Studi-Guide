import {Component, OnInit, OnDestroy} from '@angular/core';
import * as Leaflet from 'leaflet';
import { antPath } from 'leaflet-ant-path';

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

  constructor() { }

  ngOnInit() { }
  ionViewDidEnter() { this.initializeMap(); }

  private initializeMap() {
    const southWest = Leaflet.latLng(49.4126, 11.0111);
    const northEast = Leaflet.latLng(49.5118, 11.2167);
    const bounds = Leaflet.latLngBounds(southWest, northEast);

    this.map = Leaflet.map('mapId', {
      maxBounds:bounds,
      maxZoom: 19,
      minZoom: 14
    }).setView([49.452858, 11.093235], 17);

    Leaflet.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
      attribution: 'edupala.com © Angular LeafLet',
    }).addTo(this.map);

    Leaflet.marker([49.452858, 11.093235]).
      addTo(this.map).
      bindPopup('Technische Hochschule Nürnberg Georg Simon Ohm').
      openPopup();
  }

    /** Remove map when we have multiple map object */
    ngOnDestroy() {
      this.map.remove();
    }
}
