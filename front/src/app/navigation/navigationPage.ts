import { room } from '../building-objects-if';
import { testDataRooms } from './building-data';
import {Component} from "@angular/core";
import {RequestBuildingDataService} from "../services/requestBuildingData.service";

@Component({
  selector: 'app-navigation',
  templateUrl: 'navigation.page.html',
  styleUrls: ['navigation.page.scss']
})
export class NavigationPage {
  //  public mapIsVisible:boolean = true;
  public startRoom:room;
  public destinationRoom:room;
  // TODO build strings from the building data to bind only the string on the attr.d
  // e.g. "M100 100 L300 100 L300 0 L360 0 L360 130 L100 130 Z"
  public testRooms:room[] = testDataRooms;
  
  // TODO these values should be sent from backend or be clear because of the building data json response
  public svgWidth:number = 500; // this.calcSvgWidth();
  public svgHeight:number = 300; // this.calcSvgHeight();

  constructor() {}

  // TODO adapt to the current UML model
  // private calcSvgWidth() {
  //   let sum:number = 0;
  //   this.testRooms.forEach(room => {
  //     if ( room.x + room.width > sum ) {
  //       sum = room.x + room.width;
  //     }
  //   });
  //   return sum;
  // }
  //
  // private calcSvgHeight() {
  //   let sum:number = 0;
  //   this.testRooms.forEach(room => {
  //     if ( room.y + room.height > sum ) {
  //       sum = room.y + room.height;
  //     }
  //   });
  //   return sum;
  // }

  public discoverFloor() {
    // let floorToDisplay = this.startInput;
    let handleReceivedFloor = function (Tab1page, data) {
      Tab1page.testRooms = data; // JSON.parse()
    };
    let xhr = new RequestBuildingDataService();
    // TODO exchange GET to POST and uncomment floorToDisplay when API is built
    xhr.fetchDiscoverFloorData("GET", /*floorToDisplay,*/ handleReceivedFloor);
  }
}