import {Injectable} from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Env } from '../../environments/environment';
import {Location, MapItem, PathNode, SvgLocationName, SvgPath} from '../building-objects-if';
import {ReceivedRoute} from "../navigation/naviRoute";

@Injectable({
    providedIn: 'root'
})
export class DataService {

    baseUrl:string;// = SERVER_URL // "https://studi-guide.azurewebsites.net"; // for development: http://localhost:8090

    constructor(private httpClient : HttpClient, private env : Env) {
        this.baseUrl = env.serverUrl;
    }

    get_map_floor(building:string, floor:string){
        return this.httpClient.get<MapItem[]>(this.baseUrl + '/buildings/' + building + '/floors/'+ floor + '/maps');
    }

    get_route(start:string, end:string){
        return this.httpClient.get<ReceivedRoute>(this.baseUrl + '/navigation/dir?start=' + start + '&end=' + end );
    }

    get_location_search(name:string) {
        return this.httpClient.get<Location>(this.baseUrl + '/locations/' + name );
    }

    get_locations(building:string, floor:string) {
        return this.httpClient.get<Location[]>(this.baseUrl + '/buildings/'+ building +'/floors/' + floor + '/locations');
    }

    get_building(name:string) {
        return this.httpClient.get(this.baseUrl + '/buildings/' + name );
    }

    get_map_item(pathnodeid:bigint) {
        return this.httpClient.get<MapItem>(this.baseUrl + '/map?pathnodeid=' + pathnodeid);
    }

}
