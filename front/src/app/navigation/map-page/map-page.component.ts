import {Component, OnInit, OnDestroy} from '@angular/core';
import * as Leaflet from 'leaflet';
import {LatLngExpression, LeafletMouseEvent} from 'leaflet';

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

  constructor() { }

  map: Leaflet.Map;

  private readonly KV_CORNER_COORDINATES:LatLngExpression[] = [
    [49.452645,11.0924275],
    [49.45261,11.09275],
    [49.4520935,11.092615],
    [49.4521275,11.092297]
  ];

  private readonly KA_CORNER_COORDINATES:LatLngExpression[] = [
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

  private readonly KB_CORNER_COORDINATES:LatLngExpression[] = [
    [49.452445,11.09271],
    [49.4524275,11.0928575],
    [49.4525275,11.092885],
    [49.4525245,11.092965],
    [49.452545,11.0929725],
    [49.452539,11.0930575],
    [49.4525125,11.0930525],
    [49.452505,11.09312675],
    [49.4524625,11.093115],
    [49.452455,11.093169],
    [49.4522275,11.09311],
    [49.4522325,11.0930535],
    [49.45209875,11.09302075],
    [49.4521255,11.0927825],
    [49.4523425,11.0928375],
    [49.45236,11.092685]
  ];

  ngOnInit() { }
  ionViewDidEnter() { this.leafletMap(); }

  private leafletMap() {
    this.map = Leaflet.map('mapId').setView([49.452858, 11.093235], 17);
    Leaflet.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
      attribution: 'edupala.com © Angular LeafLet',
    }).addTo(this.map);

    [
      {coordinates: this.KV_CORNER_COORDINATES, options: {className: 'KV', color: '#C81111'}},
      {coordinates: this.KA_CORNER_COORDINATES, options: {className: 'KA', color: '#0083C6'}},
      {coordinates: this.KB_CORNER_COORDINATES, options: {className: 'KB', color: '#11C811'}}
    ].forEach(
        // tslint:disable-next-line:no-shadowed-variable
        (polygon) => {
          Leaflet.polygon(polygon.coordinates, polygon.options)
              .addTo(this.map)
              .on('click', this.onPolygonClick);
        }
    )

    Leaflet.marker([49.452858, 11.093235])
        .addTo(this.map)
        .bindPopup('Technische Hochschule Nürnberg Georg Simon Ohm')
        .openPopup();
  }

  private onPolygonClick(event:LeafletMouseEvent) : void
  {
    console.log(event.latlng, event.target.options.className);
    document.getElementById('close-open-map').click();
  }

  /** Remove map when we have multiple map object */
  ngOnDestroy() {
    this.map.remove();
  }
}
