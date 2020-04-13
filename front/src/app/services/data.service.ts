import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
    providedIn: 'root'
})
export class DataService {
    baseUrl:string = "http://localhost:8090"; // https://studi-guide.azurewebsites.net/roomlist/floor/0

    constructor(private httpClient : HttpClient) {

    }

    get_floor(floor:string){
        return this.httpClient.get(this.baseUrl + '/roomlist/floor' + floor);
    }

    get_route(start:string, end:string){
        return this.httpClient.get(this.baseUrl + '/navigation/dir/?startroom=' + start + '&endroom=' + end );
    }

}