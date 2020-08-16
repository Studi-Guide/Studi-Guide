import {Injectable} from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Env } from '../../environments/environment';
import {BuildingData, Location, MapItem} from '../building-objects-if';
import {ReceivedRoute} from '../navigation/map-view/naviRouteRenderer';
import {CacheService} from './cache.service';

@Injectable({
    providedIn: 'root'
})
export class DataService {

    baseUrl:string;// = SERVER_URL // "https://studi-guide.azurewebsites.net"; // for development: http://localhost:8090

    constructor(private httpClient : HttpClient, private env : Env, private cache: CacheService) {
        console.log('Using' + env.serverUrl);
        this.baseUrl = env.serverUrl;
    }

    get_map_floor(building:string, floor:string){
        return this.cache.Get<MapItem[]>(this.httpClient, this.baseUrl + '/buildings/' + building + '/floors/'+ floor + '/maps');
    }

    get_map_items(campus:string, floor: string, buildingstr:string) {
        return this.cache.Get<MapItem[]>(
            this.httpClient,
            this.baseUrl + '/maps?floor=' + floor ?? '' + '&campus=' + campus ?? '' + '&building=' + buildingstr ?? '');
    }

    get_locations_items(campus:string, floor: string, buildingstr:string) {
        return this.cache.Get<Location[]>(
            this.httpClient,
            this.baseUrl + '/locations?floor=' + floor ?? '' + '&campus=' + campus ?? '' + '&building=' + buildingstr ?? '');
    }

    get_route(start:string, end:string){
        return this.cache.Get<ReceivedRoute>(this.httpClient, this.baseUrl + '/navigation/dir?start=' + start + '&end=' + end );
    }

    get_locations_search(name:string) {
        return this.cache.Get<Location[]>(this.httpClient,this.baseUrl + '/locations?search=' + name );
    }

    get_location(name:string) {
        return this.cache.Get<Location>(this.httpClient,this.baseUrl + '/locations/' + name );
    }

    get_locations(building:string, floor:string) {
        return this.cache.Get<Location[]>(this.httpClient,this.baseUrl + '/buildings/'+ building +'/floors/' + floor + '/locations');
    }

    get_building(name:string) {
        return this.cache.Get<BuildingData>(this.httpClient,this.baseUrl + '/buildings/' + name);
    }

    get_map_item(pathnodeid:number) {
        return this.cache.Get<MapItem[]>(this.httpClient,this.baseUrl + '/maps?pathnodeid=' + pathnodeid);
    }
}
