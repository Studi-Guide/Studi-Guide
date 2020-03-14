import {room, section, svgPath} from '../building-objects-if';
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
  public testRooms:room[] = testDataRooms;
  public calculatedRoomPaths:svgPath[];
  public calculatedDoorLines:svgPath[];
  
  // TODO These values we have to determine: which size will have the scrollable map?
  public svgWidth:number = 500; // this.calcSvgWidth();
  public svgHeight:number = 1200; // this.calcSvgHeight();

  // TODO adapt to the current UML model

  private calculateSvgPaths() {
    for (const room of this.testRooms) {
      let roomShapePath:svgPath = {
        d : NavigationPage.buildRoomSvgPathFromSections(room.sections),
        fill : room.fill
      };
      this.calculatedRoomPaths.push(roomShapePath);
      if (room.doors.length >= 1) {
        for (const door of room.doors) {
          let doorLine:svgPath = {
            d : NavigationPage.buildDoorSvgLineFromSection(door),
            fill : roomShapePath.fill
          };
          this.calculatedDoorLines.push(doorLine);
        }
      }
    }
  }

  private static buildDoorSvgLineFromSection(doorSection:section) : string {
    let path:string = 'M' + doorSection.start.x + ' ' + doorSection.start.y;
    path += ' L' + doorSection.end.x + ' ' + doorSection.end.y;
    return path;
  }

  private static buildRoomSvgPathFromSections(roomSections:section[]) : string {
    let path_d:string = 'M';
    for (const section of roomSections) {
      if (path_d !== 'M') {
        path_d += 'L';
      }
      path_d += section.start.x + ' ' + section.start.y + ' ';
    }
    path_d += 'Z';
    return path_d;
  }

  constructor() {
    this.calculatedRoomPaths = [];
    this.calculatedDoorLines = [];
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