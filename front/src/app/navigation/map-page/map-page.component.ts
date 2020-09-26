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
  ionViewDidEnter() { this.leafletMap(); }

  private leafletMap() {
    this.map = Leaflet.map('mapId').setView([49.452858, 11.093235], 17);
    Leaflet.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
      attribution: 'edupala.com © Angular LeafLet',
    }).addTo(this.map);

    Leaflet.polygon(this.KV_cornerCoordinates(), {color: '#C81111'}).addTo(this.map);
    Leaflet.polygon(this.KA_cornerCoordinates(), {color: '#0083C6'}).addTo(this.map);

    Leaflet.marker([49.452858, 11.093235]).addTo(this.map).bindPopup('Technische Hochschule Nürnberg Georg Simon Ohm').openPopup();
  }

  /** Remove map when we have multiple map object */
  ngOnDestroy() {
    this.map.remove();
  }

  private KV_cornerCoordinates() : Leaflet.LatLngExpression[]
  {
    return [
      [49.452645,11.0924275],
      [49.45261,11.09275],
      [49.4520935,11.092615],
      [49.4521275,11.092297]
    ];
  }

  private KA_cornerCoordinates() : Leaflet.LatLngExpression[]
  {
    return [
      [49.4530325,11.092505],
      [49.45301,11.092705],
      [49.452992,11.09269825],
      [49.452915,11.0933975],
      [49.452935,11.093398],
      [49.4529148,11.0935875],
      [49.45289,11.09358],
      [49.452835,11.0940575],
      [49.452777575,11.09403538],
      [49.45276,11.094245],
      [49.452858,11.0942725],
      [49.45283,11.09451],
      [49.45255525,11.0944425],
      [49.4525825,11.0942125],
      [49.45255,11.09420],
      [49.452571,11.0939825],
      [49.45266175,11.0940075],
      [49.452715,11.0935275],
      [49.452695,11.09352285],
      [49.452715,11.0933425],
      [49.452735,11.0933475],
      [49.45281175,11.09265],
      [49.452705,11.092625],
      [49.4527275,11.092422]
    ];
  }
}
