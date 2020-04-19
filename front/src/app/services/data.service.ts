import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
    providedIn: 'root'
})
export class DataService {
    baseUrl:string = "https://studi-guide.azurewebsites.net"; // for development: http://localhost:8090

    constructor(private httpClient : HttpClient) {

    }

    get_floor(floor:any) {
        return this.httpClient.get(this.baseUrl + '/roomlist/floor/' + floor);
    }

    get_route(start:string, end:string)  {
        return this.httpClient.get(this.baseUrl + '/navigation/dir?startroom=' + start + '&endroom=' + end ); // '/navigation/dir/startroom/KA.012/endroom/KA.013
    }

    get_room_search(room:string) {
        return this.httpClient.get(this.baseUrl + '/roomlist/room/' + room );
    }
}