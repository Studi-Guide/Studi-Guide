import {HttpClient, HttpHandler} from '@angular/common/http';
import { floor } from '../building-objects-if';
import { room } from '../building-objects-if';
import { corridor } from '../building-objects-if';
import { testDataRooms } from './building-data';
import {Component} from "@angular/core";
import {parseJson} from "@angular-devkit/core";
import {stringify} from "querystring";

@Component({
  selector: 'app-tab1',
  templateUrl: 'tab1.page.html',
  styleUrls: ['tab1.page.scss']
})
export class Tab1Page {
//  public mapIsVisible:boolean = true;
//  public cssSvgMap:string;
  private http:HttpClient;
  public startRoom:room;
  public destinationRoom:room;
  public testRooms:room[] = testDataRooms;
  
  // TODO these values should be sent from backend
  public svgWidth:number = this.calcSvgWidth();
  public svgHeight:number = this.calcSvgHeight();

  constructor() {
    // this.http = new HttpClient('sdfsdf');
  }

  private calcSvgWidth() {
    let sum:number = 0;
    this.testRooms.forEach(room => {
      if ( room.x + room.width > sum ) {
        sum = room.x + room.width;
      }
    });
    return sum;
  }

  private calcSvgHeight() {
    let sum:number = 0;
    this.testRooms.forEach(room => {
      if ( room.y + room.height > sum ) {
        sum = room.y + room.height;
      }
    });
    return sum;
  }

  public requestBuildingData(floorOne, floorTwo=null){
    // call ajax here
    if (floorTwo != null) {
      // a route is requested
    } else {
      // only the view of one floor is needed

    }
  };

  sendPostRequest() {
    let res = this.http.get('http://127.0.0.1:8080/roomlist/');
    console.log(stringify(res));
  }
}