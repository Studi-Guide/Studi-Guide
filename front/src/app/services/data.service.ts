import {Injectable} from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Env } from '../../environments/environment';
import {BuildingData, Location, MapItem} from '../building-objects-if';
import {ReceivedRoute} from '../navigation/map-view/naviRouteRenderer';
import {Observable, of} from "rxjs";

@Injectable({
    providedIn: 'root'
})
export class DataService {

    baseUrl:string;// = SERVER_URL // "https://studi-guide.azurewebsites.net"; // for development: http://localhost:8090

    cache = {};

    constructor(private httpClient : HttpClient, private env : Env) {
        this.baseUrl = env.serverUrl;
    }

    get_map_floor(building:string, floor:string){
        return this.runGetRequest<MapItem[]>(this.baseUrl + '/buildings/' + building + '/floors/'+ floor + '/maps');
    }

    get_route(start:string, end:string){
        return this.runGetRequest<ReceivedRoute>(this.baseUrl + '/navigation/dir?start=' + start + '&end=' + end );
    }

    get_locations_search(name:string) {
        return this.runGetRequest<Location[]>(this.baseUrl + '/locations?search=' + name );
    }

    get_location(name:string) {
        return this.runGetRequest<Location>(this.baseUrl + '/locations/' + name );
    }

    get_locations(building:string, floor:string) {
        return this.runGetRequest<Location[]>(this.baseUrl + '/buildings/'+ building +'/floors/' + floor + '/locations');
    }

    get_building(name:string) {
        return this.runGetRequest<BuildingData>(this.baseUrl + '/buildings/' + name);
    }

    get_map_item(pathnodeid:number) {
        return this.runGetRequest<MapItem[]>(this.baseUrl + '/maps?pathnodeid=' + pathnodeid);
    }

    runGetRequest<T>(request: string): Observable<T> {
        if (this.cache[request]) {
            console.log('Returning cached value!')
            return this.cache[request]
        }

        const observable = this.httpClient.get<T>(request);
        this.cache[request] =observable;
        return observable;
    }
}
