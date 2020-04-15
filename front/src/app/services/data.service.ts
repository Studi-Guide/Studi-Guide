import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
    providedIn: 'root'
})
export class DataService {
    //baseUrl:string = "https://studi-guide.azurewebsites.net"; // for development: http://localhost:8090
    baseUrl:string = "http://127.0.0.1:8080"

    constructor(private httpClient : HttpClient) {

    }

    get_map_floor(floor:any){
        return this.httpClient.get(this.baseUrl + '/map/floor/' + floor);
    }

    get_route(start:string, end:string){
        return this.httpClient.get(this.baseUrl + '/navigation/dir?start=' + start + '&end=' + end );
    }

    get_room_search(room:string) {
        return this.httpClient.get(this.baseUrl + '/roomlist/room/' + room );
    }

    get_locations(floor:any) {
        return this.httpClient.get(this.baseUrl + '/location/?floor=' + floor);
    }

}