import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment, SERVER_URL} from '../../environments/environment';

@Injectable({
    providedIn: 'root'
})
export class DataService {

    baseUrl:string = SERVER_URL // "https://studi-guide.azurewebsites.net"; // for development: http://localhost:8090

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