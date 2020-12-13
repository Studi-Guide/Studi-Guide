import {LOCALE_ID, Injectable, Inject} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Env} from '../../../environments/environment';
import {LatLngLiteral} from 'leaflet';

export interface OpenStreetMapLiteral {
  Lat: number
  Lng: number
}

export interface OpenStreetMapBounds {
  SouthWest: OpenStreetMapLiteral
  NorthEast: OpenStreetMapLiteral
}

export interface OsmRouteInstruction {
  Distance: number;
  Heading: number;
  Interval: number[];
  Sign: number;
  StreetName: string;
  Text: string;
  Time: number;
}

export interface OsmRoutePoints {
  Coordinates: OpenStreetMapLiteral[];
  Type: string;
}

export interface OsmRoute {
  Distance: number;
  Time: number;
  Points: OsmRoutePoints;
  Instructions: OsmRouteInstruction[];
}

@Injectable({
  providedIn: 'root'
})
export class OpenStreetMapService {

  private static serverPathPrefix = '/osm';
  private static routePath = '/route'
  private static boundsPath = '/bounds';

  constructor(@Inject(LOCALE_ID) private locale: string,
              private httpClient : HttpClient,
              private env : Env) { }

  public async GetRoute(startPoint:LatLngLiteral, endPoint:LatLngLiteral) : Promise<OsmRoute[]> {
    // curl "https://graphhopper.com/api/1/route?point=51.131,12.414&point=48.224,3.867&vehicle=car&locale=de&calc_points=false&key=api_key"
    const r = await this.httpClient.get(this.env.serverUrl+OpenStreetMapService.serverPathPrefix+OpenStreetMapService.routePath
        +'?start='+startPoint.lat+','+startPoint.lng+'&end='+endPoint.lat+','+endPoint.lng+'&locale='+this.locale)
        .toPromise() as OsmRoute[];
    return r;
  }

  public async GetBounds() : Promise<OpenStreetMapBounds> {
    return await this.httpClient.get(this.env.serverUrl
        +OpenStreetMapService.serverPathPrefix+OpenStreetMapService.boundsPath).toPromise() as OpenStreetMapBounds
  }

}
