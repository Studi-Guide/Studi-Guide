import {Coordinate, Room, Section, svgPath} from '../building-objects-if';
import {testDataRooms} from './test-building-data';
import {Component} from "@angular/core";
import {RequestBuildingDataService} from "../services/requestBuildingData.service";

@Component({
  selector: 'app-navigation',
  templateUrl: 'navigation.page.html',
  styleUrls: ['navigation.page.scss']
})

export class NavigationPage {
  public progressIsVisible:boolean = false;
  public routeInputIsVisible:boolean = false;
  public searchBtnIsVisible:boolean = true;
  public routeBtnIsVisible:boolean = true;
  public mapIsVisible:boolean = false;
  public startInput:string;
  public destinationInput:string;
  public startRoom:Room;
  public destinationRoom:Room;
  public testRooms:Room[] = testDataRooms;
  public testPathNodes:Coordinate[];
  public calculatedRoomPaths:svgPath[];
  public calculatedDoorLines:svgPath[];
  
  public svgWidth:number = 0;
  public svgHeight:number = 0;

  private static buildDoorSvgLineFromSection(doorSection:Section) : string {
    let path:string = 'M' + doorSection.Start.X + ' ' + doorSection.Start.Y;
    path += ' L' + doorSection.End.X + ' ' + doorSection.End.Y;
    return path;
  }

  private static buildRoomSvgPathFromSections(roomSections:Section[]) : string {
    let path_d:string = 'M';
    for (const section of roomSections) {
      if (path_d !== 'M') {
        path_d += 'L';
      }
      path_d += section.Start.X + ' ' + section.Start.Y + ' ';
    }
    path_d += 'Z';
    return path_d;
  }

  private calculateSvgPathsAndSvgWidthHeight() {
    for (const room of this.testRooms) {
      let roomShapePath:svgPath = {
        d : NavigationPage.buildRoomSvgPathFromSections(room.sections),
        fill : room.Color
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
      for (const section of room.sections) {
        if (section.End.X > this.svgWidth) {
          this.svgWidth = section.End.X;
        }
        if (section.End.Y > this.svgHeight) {
          this.svgHeight = section.End.Y;
        }
      }
      // bottom navigation bar overlays svg
      this.svgHeight += 1;
      this.svgWidth += 0.15;
    }
  }

  private static testRenderPathNodes() : Coordinate[] {
    let pathNodes:Coordinate[] = [];
    for (const room of testDataRooms) {
      for (const pathNode of room.pathNodes) {
        pathNodes.push(pathNode);
      }
      for (const door of room.doors) {
        pathNodes.push(door.pathNode);
      }
    }
    return pathNodes;
  }

  constructor() {
    this.calculatedRoomPaths = [];
    this.calculatedDoorLines = [];
    this.calculateSvgPathsAndSvgWidthHeight();
    this.testPathNodes = NavigationPage.testRenderPathNodes();
  }

  public showFloor() {
    if (this.routeInputIsVisible) {
      this.routeInputIsVisible = false;
    } else if (this.startInput != undefined) {
      this.mapIsVisible = true;
    }
  }

  public showRoute() {
    if (!this.routeInputIsVisible) {
      this.routeInputIsVisible = true;
    } else if (this.startInput != undefined && this.destinationInput != undefined) {
      this.mapIsVisible = true;
    }
  }

  public discoverFloor() {
    // let floorToDisplay = this.startInput;
    let handleReceivedFloor = function (data) {
      console.log(data); // JSON.parse()
    };
    let xhr = new RequestBuildingDataService();
    // TODO add input data fetching from UI
    xhr.fetchDiscoverFloorData('GET', 'http://localhost:8090/api', 'KA.3', handleReceivedFloor);
  }
}