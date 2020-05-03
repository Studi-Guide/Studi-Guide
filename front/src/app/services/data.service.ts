import {Injectable} from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Env } from '../../environments/environment';

@Injectable({
    providedIn: 'root'
})
export class DataService {

    baseUrl:string;// = SERVER_URL // "https://studi-guide.azurewebsites.net"; // for development: http://localhost:8090

    constructor(private httpClient : HttpClient, private env : Env) {
        this.baseUrl = env.serverUrl;
    }

    get_map_floor(building:string, floor:string){
        return this.httpClient.get(this.baseUrl + '/buildings/' + building + '/floors/'+ floor + '/maps');
    }

    get_route(start:string, end:string){
        return this.httpClient.get(this.baseUrl + '/navigation/dir?start=' + start + '&end=' + end );
    }

    get_location_search(name:string) {
        return this.httpClient.get(this.baseUrl + '/locations/' + name );
    }

    get_locations(building:string, floor:string) {
        return this.httpClient.get(this.baseUrl + '/buildings/'+ building +'/floors/' + floor + '/locations');
    }

    get_building(name:string) {
        return this.httpClient.get(this.baseUrl + '/buildings/' + name );
    }

}