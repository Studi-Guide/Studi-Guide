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
//  private http:HttpClient;
  public startRoom:room;
  public destinationRoom:room;
  public testRooms:room[] = testDataRooms;
  
  // TODO these values should be sent from backend
  public svgWidth:number = this.calcSvgWidth();
  public svgHeight:number = this.calcSvgHeight();

  constructor(private http:HttpClient) {
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

  // public requestBuildingData(floorOne, floorTwo=null){
  //   // call ajax here
  //   if (floorTwo != null) {
  //     // a route is requested
  //   } else {
  //     // only the view of one floor is needed
  //     // sendPostRequest(floorOne);
  //   }
  // };

  // public sendPostRequest(/*dataToSent*/) {
  //   // let res = this.http.get('https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fsagarhani.files.wordpress.com%2F2015%2F07%2Fduck_duck_go.png&f=1&nofb=1'); // http://127.0.0.1:8080/roomlist/
  //
  //   let res = null;
  //   $http({
  //     method : "GET",
  //     url : "https://www.w3schools.com/js/demo_get.asp?t=0.5814227109783404"
  //   }).then(function(response) {
  //     res = response.data;
  //   });
  //   console.log(res);
  // }

  public getData() {
    var req = new XMLHttpRequest();
    req.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
        console.log(req.responseText)
      }
    };
    let method:string = "GET";
    // TODO solve blocked Cross-Origin request caused by different ports of the dev servers (8100 & 8080)
    let url:string = "http://127.0.0.1:8080/roomlist/";
    let async:boolean = true;
    req.open(method, url, async);
    req.send();
  }
}