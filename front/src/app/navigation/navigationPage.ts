import {room, svgPath} from '../building-objects-if';
import { testDataRooms } from './building-data';
import {Component} from "@angular/core";
import {RequestBuildingDataService} from "../services/requestBuildingData.service";
import {forEach} from "@angular-devkit/schematics";

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
  public calculatedPaths:svgPath[];
  
  // TODO These values we have to determine: which size will have the scrollable map?
  public svgWidth:number = 500; // this.calcSvgWidth();
  public svgHeight:number = 1000; // this.calcSvgHeight();

  // TODO adapt to the current UML model

  private calculateSvgPaths() {
    for (const room of this.testRooms) {
      let roomShapePath:svgPath = {
        d:'',
        fill:''
      };
      roomShapePath.d = NavigationPage.buildRoomSvgPathFromSections(room.sections);
      roomShapePath.fill = room.fill;
      this.calculatedPaths.push(roomShapePath);
      // path = NavigationPage.buildDoorSvgPath(room.doors);
    }
  }

// TODO buildDoorSvgPathFromDoors is missing yet

  private static buildRoomSvgPathFromSections(roomSections) : string {
    let path_d:string = 'M';
    for (const section of roomSections) {
      path_d += section.start.x+' '+section.start.y+' ';
    }
    path_d += 'Z';
    return path_d;
  }

  constructor() {
    this.calculatedPaths = [];
    this.calculateSvgPaths();
  }

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