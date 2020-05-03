import {Location, MapItem, PathNode, SvgLocationName, SvgPath} from '../building-objects-if';
// import {testDataRooms, testDataPathNodes} from './test-building-data';
import {Component, ViewChild} from '@angular/core';
import {ModalController} from '@ionic/angular';
import {DataService} from '../services/data.service';
import {FloorMap} from './floorMap';
import {DistanceToBeDisplayed, NaviRoute, ReceivedRoute} from './naviRoute';
import {AvailableFloorsPage} from '../available-floors/available-floors.page';

@Component({
  selector: 'app-navigation',
  templateUrl: 'navigation.page.html',
  styleUrls: ['navigation.page.scss']
})

export class NavigationPage {
  @ViewChild('discoverySearchbar') discoverySearchbarRef;

  public progressIsVisible = false;
  public routeInputIsVisible = false;
  public searchBtnIsVisible = true;
  public closeRouteBtnIsVisible = false;

  public startInput: string;
  public destinationInput: string;

  private route: NaviRoute;
  public distanceToDisplay: DistanceToBeDisplayed;
  public calculatedRoute: string;
  public routeIsVisible = false;
  public startPin: PathNode;
  public routeEnd: PathNode;

  public startPinIsVisible = false;
  public distanceIsVisible = false;

  private floor: FloorMap;
  public calculatedRoomPaths: SvgPath[];
  public calculatedDoorLines: SvgPath[];
  public mapSvgWidth: number;
  public mapSvgHeight: number;
  public locations: SvgLocationName[];
  public mapIsVisible = false;

//  public testRooms:Room[] = [];
//  public testRoute:PathNode[];

  constructor(private dataService: DataService,
              private modalCtrl: ModalController) {
    this.dataService = dataService;

    this.calculatedRoute = '';

    this.calculatedRoomPaths = [];
    this.calculatedDoorLines = [];
    this.mapSvgWidth = 0;
    this.mapSvgHeight = 0;
    this.locations = [];

    // this.testRooms = testDataRooms;
    // this.testRoute = testDataPathNodes;
    // this.testRoute = NavigationPage.testRenderPathNodes();
  }

  public showFloorForSearch() {
    if (this.routeInputIsVisible) {
      this.hideRouteSearchbar();
    } else if (this.startInput !== undefined && this.startInput !== '' && this.startInput != null) {
      this.fetchFloorByLocation(this.startInput);
      this.routeIsVisible = false;
      this.startPinIsVisible = false;
      this.mapIsVisible = true;
    }
  }

  private fetchFloorByLocation(room: string) {
    this.progressIsVisible = true;
    this.dataService.get_location_search(room).subscribe((res : Location)=>{
      this.startPin = res.PathNode;
      this.startPinIsVisible = true;
      this.fetchFloorByItsNumber(res.Building, res.Floor);
      this.fetchLocations(res.Building, res.Floor);
    });
  }

  public showRoute() {
    if (!this.routeInputIsVisible) {
      this.routeInputIsVisible = true;
      // TODO set #discoverySearchbar color blue
      const searchbars = document.querySelector('ion-item');
      searchbars.setAttribute('color', 'primary');
      this.searchBtnIsVisible = false;
      this.closeRouteBtnIsVisible = true;
    } else if (this.startInput !== undefined && this.destinationInput !== undefined
        && this.startInput !== '' && this.destinationInput !== ''
        && this.startInput != null && this.destinationInput != null
    ) {
      this.mapIsVisible = true;
      this.fetchRouteToDisplay(this.startInput, this.destinationInput);
    }
  }

  private fetchFloorByItsNumber(building:string, floor:string) {
    this.progressIsVisible = true;
    this.dataService.get_map_floor(building, floor).subscribe((res : MapItem[])=>{
      this.floor = new FloorMap(res);
      this.displayFloor();
    });
  }

  private fetchRouteToDisplay(start:string, end:string) {
    this.progressIsVisible = true;
    this.dataService.get_location_search(start).subscribe((res1 : Location)=>{
      this.dataService.get_map_floor(res1.Building, res1.Floor).subscribe((res2 : MapItem[])=>{
        this.fetchLocations(res2[0].Building, res2[0].Floor);
        this.floor = new FloorMap(res2);
        this.dataService.get_route(start, end).subscribe((res3 : ReceivedRoute)=>{
          this.route = new NaviRoute(res3);

          this.distanceToDisplay = this.route.distance;
          this.calculatedRoute = this.route.svgRoute;

          this.startPin = this.route.getRouteStart();
          this.routeEnd = this.route.getRouteEnd();
          this.displayFloor();

          this.progressIsVisible = false;
          this.routeIsVisible = true;
          this.startPinIsVisible = true;
          this.distanceIsVisible = true;
        });
      });
    });
  }

  private fetchLocations(building:string, floor:string) {
    this.dataService.get_locations(building, floor).subscribe((res : Location[])=>{
      this.locations = [];
      for(const l of res) {
        this.locations.push({name: l.Name, x: l.PathNode.Coordinate.X, y: l.PathNode.Coordinate.Y})
      }
    })
  }

  private displayFloor() {
    this.floor.calculateSvgPathsAndSvgWidthHeight();
    this.mapSvgHeight = this.floor.svgHeight;
    this.mapSvgWidth = this.floor.svgWidth;
    this.calculatedRoomPaths = this.floor.calculatedRoomPaths;
    this.calculatedDoorLines = this.floor.calculatedDoorLines;

    this.progressIsVisible = false;
    this.mapIsVisible = true;
    this.routeIsVisible = false;
    this.distanceIsVisible = false;
  }

  public checkWhatIsRequestedByEnterKey() {
    if (this.startInput !== undefined && this.startInput !== '' && this.startInput !== null &&
        !this.routeInputIsVisible
    ) {
      this.showFloorForSearch();
    } else if (this.startInput !== undefined && this.startInput !== '' && this.startInput !== null &&
        this.routeInputIsVisible &&
        (this.destinationInput === undefined || this.destinationInput === '' || this.destinationInput === null)
    ) {
      // TODO check setFocus on Android, iOS, etc.
      this.discoverySearchbarRef.setFocus();
      this.showFloorForSearch();
    } else if (this.startInput !== undefined && this.startInput !== '' && this.startInput != null &&
        this.routeInputIsVisible &&
        this.destinationInput !== undefined && this.destinationInput !== '' && this.destinationInput !== null
    ) {
      this.showRoute();
    }
  }

  public hideRouteSearchbar() {
    this.routeInputIsVisible = false;
    const searchbars = document.querySelector('ion-item');
    searchbars.setAttribute('color', 'light-tint');
    this.searchBtnIsVisible = true;
    this.closeRouteBtnIsVisible = false;
  }

  async presentAvailableFloorModal() {
    this.startPinIsVisible = false;
    this.dataService.get_building(this.startInput.slice(0, 2)).subscribe(async (res: JSON) => {
      // @ts-ignore
      const {Floors} = res;
      const availableFloorModal = await this.modalCtrl.create({
        component: AvailableFloorsPage,
        cssClass: 'floor-modal',
        componentProps: {
          floors: Floors
        }
      });
      await availableFloorModal.present();

      availableFloorModal.onDidDismiss()
          .then((data) => {
            if (data.data) {
              const building = this.startInput.slice(0, 2);
              this.fetchFloorByItsNumber(building, data.data)
              this.fetchLocations(building, data.data)
            }
          })


    });
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