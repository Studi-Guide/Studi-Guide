import {Room, svgPath, RoomName, PathNode} from '../building-objects-if';
// import {testDataRooms, testDataPathNodes} from './test-building-data';
import {Component} from "@angular/core";
import {DataService} from "../services/data.service";
import {FloorMap} from "./floorMap";
import {DistanceToBeDisplayed, NaviRoute, ReceivedRoute} from "./naviRoute";

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

  public startInput: string;
  public destinationInput: string;

  private route: NaviRoute;
  public distanceToDisplay: DistanceToBeDisplayed;
  public calculatedRoute: string;
  public routeIsVisible: boolean = false;
  public routeStart: PathNode; // = {"Coordinate":{"X":0, "Y":0, "Z":0}}
  public routeEnd: PathNode; // = {"Coordinate":{"X":0, "Y":0, "Z":0}};

  public startIsVisible: boolean = false;
  public distanceIsVisible: boolean = false;

  private floor: FloorMap;
  public calculatedRoomPaths: svgPath[];
  public calculatedDoorLines: svgPath[];
  public mapSvgWidth: number;
  public mapSvgHeight: number;
  public roomNames: RoomName[];
  public mapIsVisible: boolean = false;

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
      this.fetchFloorByRoom(this.startInput);
    }
  }

  private fetchFloorByRoom(room: string) {
    this.progressIsVisible = true;
    this.dataService.get_room_search(room).subscribe((res : Room)=>{
      this.fetchFloorByItsNumber(res.Floor);
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
//      this.fetchFloorByRoom(this.startInput);
      this.fetchRouteToDisplay(this.startInput, this.destinationInput);
    }
  }

  private fetchFloorByItsNumber(floor:any) {
    this.progressIsVisible = true;
    this.dataService.get_floor(floor).subscribe((res : Room[])=>{
      this.floor = new FloorMap(res);
      this.routeIsVisible = false;
      this.displayFloor();
    });
  }

  private fetchRouteToDisplay(start:string, end:string) {
    this.progressIsVisible = true;
    this.dataService.get_room_search(start).subscribe((res1 : Room)=>{
      this.dataService.get_floor(res1.Floor).subscribe((res2 : Room[])=>{
        this.floor = new FloorMap(res2);
        this.dataService.get_route(start, end).subscribe((res3 : ReceivedRoute)=>{
          this.route = new NaviRoute(res3);

          // this.sourceSvg = '<image x="100" y="200" width="20" height="20" xlink:href="../../assets/navigation-svgs/race-flag.svg" />';
          this.distanceToDisplay = this.route.distance;
          this.calculatedRoute = this.route.svgRoute;

          this.progressIsVisible = false;
          this.routeIsVisible = true;
          this.routeStart = this.route.getRouteStart();

          this.displayFloor();
          this.startIsVisible = true;
          this.distanceIsVisible = true;
        });
      });
    });
  }

  private displayFloor() {
    this.floor.calculateSvgPathsAndSvgWidthHeight();
    this.mapSvgHeight = this.floor.svgHeight;
    this.mapSvgWidth = this.floor.svgWidth;
    this.calculatedRoomPaths = this.floor.calculatedRoomPaths;
    this.calculatedDoorLines = this.floor.calculatedDoorLines;
    this.floor.collectAllRoomNames();
    this.roomNames = this.floor.allRoomNames;

    this.progressIsVisible = false;
    this.mapIsVisible = true;
    this.distanceIsVisible = false;
    this.startIsVisible = false;
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