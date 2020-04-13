import {PathNode, Room, svgPath, RoomName} from '../building-objects-if';
// import {testDataRooms, testDataPathNodes} from './test-building-data';
import {Component} from "@angular/core";
import {DataService} from "../services/data.service";
import {FloorMap} from "./floorMap";
import {NaviRoute} from "./naviRoute";

@Component({
  selector: 'app-navigation',
  templateUrl: 'navigation.page.html',
  styleUrls: ['navigation.page.scss']
})

export class NavigationPage {
  public progressIsVisible: boolean = false;
  public routeInputIsVisible: boolean = false;
  public searchBtnIsVisible: boolean = true;
  public routeBtnIsVisible: boolean = true;
  public mapIsVisible: boolean = false;
  public routeIsVisible: boolean = false;

  public startInput: string;
  public destinationInput: string;

  private routeToDisplay: NaviRoute;
  public calculatedRoute: string;

  private floorToDisplay: FloorMap;
  public calculatedRoomPaths: svgPath[];
  public calculatedDoorLines: svgPath[];
  public mapSvgWidth: number;
  public mapSvgHeight: number;
  public roomNames: RoomName[];

//  public testRooms:Room[] = [];
//  public testRoute:PathNode[];

  // public sourceSvg: string;

  constructor(private dataService: DataService) {
    this.dataService = dataService;

    this.calculatedRoute = '';

    this.calculatedRoomPaths = [];
    this.calculatedDoorLines = [];
    this.mapSvgWidth = 0;
    this.mapSvgHeight = 0;
    this.roomNames = [];

    // this.testRooms = testDataRooms;
    // this.testRoute = testDataPathNodes;
    // this.testRoute = NavigationPage.testRenderPathNodes();
  }

  public showFloorForSearch() {
    if (this.routeInputIsVisible) {
      this.routeInputIsVisible = false;
    } else if (this.startInput != undefined && this.startInput != '' && this.startInput != null) {
      this.mapIsVisible = true;
      this.fetchFloorForSearch(this.startInput);
    }
  }

  private fetchFloorForSearch(room: string) {
    this.progressIsVisible = true;
    this.dataService.get_room_search(room).subscribe((res : Room[])=>{
      this.floorToDisplay = new FloorMap(res);
      this.displayFloor();
    });
  }

  public showRoute() {
    if (!this.routeInputIsVisible) {
      this.routeInputIsVisible = true;
    } else if (this.startInput != undefined && this.destinationInput != undefined
        && this.startInput != '' && this.destinationInput != ''
        && this.startInput != null && this.destinationInput != null
    ) {
      this.mapIsVisible = true;
      // TODO only in KA.304 is the 4. character always the floor
      this.fetchFloorForNavi(this.startInput[3]);
      this.fetchRouteToDisplay(this.startInput, this.destinationInput); // 'KA.308','KA.313'
    }
  }

  private fetchFloorForNavi(floor:string) {
    this.progressIsVisible = true;
    this.dataService.get_floor(floor).subscribe((res : Room[])=>{
      this.floorToDisplay = new FloorMap(res);
      this.displayFloor();
    });
  }

  private fetchRouteToDisplay(start:string, end:string) {
    this.progressIsVisible = true;
    this.dataService.get_route(start, end).subscribe((res : PathNode[])=>{
      this.routeToDisplay = new NaviRoute(res);
      console.log(res);

      this.routeToDisplay.calculateSvgPathForRoute();
      // this.sourceSvg = '<image x="100" y="200" width="20" height="20" xlink:href="../../assets/navigation-svgs/race-flag.svg" />';
      this.calculatedRoute = this.routeToDisplay.svgRoute;

      this.progressIsVisible = false;
      this.routeIsVisible = true;
    });
  }

  private displayFloor() {
    this.floorToDisplay.calculateSvgPathsAndSvgWidthHeight();
    this.mapSvgHeight = this.floorToDisplay.svgHeight;
    this.mapSvgWidth = this.floorToDisplay.svgWidth;
    this.calculatedRoomPaths = this.floorToDisplay.calculatedRoomPaths;
    this.calculatedDoorLines = this.floorToDisplay.calculatedDoorLines;
    this.floorToDisplay.collectAllRoomNames();
    this.roomNames = this.floorToDisplay.allRoomNames;

    this.progressIsVisible = false;
    this.mapIsVisible = true;
  }

/*  private static testRenderPathNodes() : Coordinate[] {
    let pathNodes:Coordinate[] = [];
        for (const room of testDataRooms) {
          for (const pathNode of room.PathNodes) {
            pathNodes.push(pathNode.Coordinate);
          }
          for (const door of room.Doors) {
            pathNodes.push(door.pathNode.Coordinate);
          }
        }
    return pathNodes;
  }*/
}