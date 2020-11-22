import {LOCALE_ID, Injectable, Inject} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Env} from '../../../environments/environment';
import {LatLngLiteral} from 'leaflet';
import {catchError} from "rxjs/operators";

@Injectable({
  providedIn: 'root'
})
export class GraphHopperService {

  private ghRootUrl = 'https://graphhopper.com/api/1/';
  private ghRouteUrl = 'route';
  private ghVehicle = 'foot';
  private ghApiKey = 'eb519f1e-396e-4ffd-b535-5abdbf3fc49f';
  // other APIs: private ghMatrixUrl = ...

  constructor(@Inject(LOCALE_ID) private locale: string,
              private httpClient : HttpClient,
              private env : Env) { }

  public async GetRouteEndpoint(startPoint:LatLngLiteral, endPoint:LatLngLiteral) {
    // curl "https://graphhopper.com/api/1/route?point=51.131,12.414&point=48.224,3.867&vehicle=car&locale=de&calc_points=false&key=api_key"
    const r = await this.httpClient.get(this.ghRootUrl+this.ghRouteUrl
        +'?point='+startPoint.lat+','+startPoint.lng+'&point='+endPoint.lat+','+endPoint.lng
        +'&vehicle='+this.ghVehicle+'&key='+this.ghApiKey+'&locale='+this.locale+'&points_encoded=false').toPromise();
    return r;
  }

}
