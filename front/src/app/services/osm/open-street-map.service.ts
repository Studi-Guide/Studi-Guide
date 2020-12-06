import {LOCALE_ID, Injectable, Inject} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Env} from '../../../environments/environment';
import {LatLngLiteral} from 'leaflet';

export interface GraphHopperInstruction {
  distance: number;
  heading: number;
  interval: number[];
  sign: number;
  street_name: string;
  text: string;
  time: number;
}

export interface GraphHopperPoints {
  coordinates: [number, number][];
  type: string;
}

export interface GraphHopperPath {
  distance: number;
  time: number;
  points: GraphHopperPoints;
  instructions: GraphHopperInstruction[];
}

export interface GraphHopperRoute {
  paths: GraphHopperPath[];
}

@Injectable({
  providedIn: 'root'
})
export class OpenStreetMapService {

  private static serverPath = '/osm/route';

  constructor(@Inject(LOCALE_ID) private locale: string,
              private httpClient : HttpClient,
              private env : Env) { }

  public async GetRoute(startPoint:LatLngLiteral, endPoint:LatLngLiteral) : Promise<GraphHopperRoute> {
    // curl "https://graphhopper.com/api/1/route?point=51.131,12.414&point=48.224,3.867&vehicle=car&locale=de&calc_points=false&key=api_key"
    const r = await this.httpClient.get(this.env.serverUrl+OpenStreetMapService.serverPath
        +'?start='+startPoint.lat+','+startPoint.lng+'&end='+endPoint.lat+','+endPoint.lng).toPromise() as GraphHopperRoute;
    return r;
  }

}
