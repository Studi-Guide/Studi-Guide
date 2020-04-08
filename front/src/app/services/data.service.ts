import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
    providedIn: 'root'
})
export class DataService {
    baseUrl:string = "http://localhost:8090/api";

    constructor(private httpClient : HttpClient) {

    }

    get_floor(floor:string){
        return this.httpClient.get(this.baseUrl + '/' + floor);
    }

    get_route(startRoom:string, destinationRoom:string){
        return this.httpClient.get(this.baseUrl + '/' + startRoom + '-' + destinationRoom );
    }

}