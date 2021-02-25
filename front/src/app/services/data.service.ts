import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import { Env } from '../../environments/environment';
import {IBuilding, ICampus, ILocation, IMapItem} from '../building-objects-if';
import {CacheService} from './cache.service';
import {IReceivedRoute} from '../route-objects-if';

@Injectable({
    providedIn: 'root'
})
export class DataService {

    baseUrl: string; // = SERVER_URL // "https://studi-guide.azurewebsites.net"; // for development: http://localhost:8090

    constructor(private httpClient: HttpClient, private env: Env, private cache: CacheService) {
        console.log('Using ', env.serverUrl);
        this.baseUrl = env.serverUrl;
    }

    get_map_floor(building: string, floor: string){
        return this.cache.Get<IMapItem[]>(this.httpClient, this.baseUrl + '/maps/buildings/' + building + '/floors/' + floor);
    }

    get_map_items(campus: string, floor: string, building: string) {
        const campusStr = campus ?? '';
        const buildingStr = building ?? '';
        const floorStr = floor ?? '';
        const request = this.baseUrl + '/maps?floor=' + floorStr + '&campus=' + campusStr + '&building=' + buildingStr;
        return this.cache.Get<IMapItem[]>(
            this.httpClient,
            request);
    }

    get_locations_items(campus: string, floor: string, building: string) {
        const campusStr = campus ?? '';
        const buildingStr = building ?? '';
        const floorStr = floor ?? '';
        return this.cache.Get<ILocation[]>(
            this.httpClient,
            this.baseUrl + '/locations?floor=' + floorStr + '&campus=' + campusStr + '&building=' + buildingStr);
    }

    get_route(start: string, end: string){
        return this.cache.Get<IReceivedRoute>(this.httpClient, this.baseUrl + '/navigation/dir?start='
            + encodeURI(start) + '&end=' + encodeURI(end));
    }

    get_locations_all() {
        return this.cache.Get<ILocation[]>(this.httpClient, this.baseUrl + '/locations');
    }

    get_locations_search(name: string) {
        return this.cache.Get<ILocation[]>(this.httpClient, this.baseUrl + '/locations?search=' + encodeURI(name) );
    }

    get_location(name: string) {
        return this.cache.Get<ILocation>(this.httpClient, this.baseUrl + '/locations/' +  encodeURI(name));
    }

    get_locations(building: string, floor: string) {
        return this.cache.Get<ILocation[]>(this.httpClient, this.baseUrl + '/buildings/' + building + '/floors/' + floor + '/locations');
    }

    get_building(name: string, logError: boolean = true) {
        return this.cache.Get<IBuilding>(this.httpClient, this.baseUrl + '/buildings/' + encodeURI(name), logError);
    }

    get_building_floor(name: string, logError: boolean = true) {
        return this.cache.Get<string[]>(this.httpClient, this.baseUrl + '/buildings/' + encodeURI(name) + '/floors', logError);
    }

    get_buildings_search(search: string = '') {
        const searchParam = search ? '?name=' + search : '';
        return this.cache.Get<IBuilding[]>(this.httpClient, this.baseUrl + '/buildings' + encodeURI(searchParam));
    }

    get_map_item(pathNodeId: number) {
        return this.cache.Get<IMapItem[]>(this.httpClient, this.baseUrl + '/maps?pathnodeid=' + pathNodeId);
    }

    get_campus(name: string, logError: boolean = true) {
        return this.cache.Get<ICampus>(this.httpClient, this.baseUrl + '/campus/' + encodeURI(name), logError);
    }

    get_campus_search(search: string = '') {
        const searchParam = search ? '?search=' + search : '';
        return this.cache.Get<ICampus[]>(this.httpClient, this.baseUrl + '/campus' + encodeURI(searchParam));
    }

    get_rssFeed(feedId: string){
        return this.httpClient.get(this.baseUrl + '/rssfeed/' + feedId,  {responseType: 'text'});
    }
}
