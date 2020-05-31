import {Location, MapItem, PathNode} from '../building-objects-if';
// import {testDataRooms, testDataPathNodes} from './test-building-data';
import {Component, ViewChild} from '@angular/core';
import {ModalController} from '@ionic/angular';
import {DataService} from '../services/data.service';
import {FloorMap} from './floorMap';
import {NaviRoute, ReceivedRoute} from './naviRoute';
import {AvailableFloorsPage} from '../available-floors/available-floors.page';

@Component({
  selector: 'app-navigation',
  templateUrl: 'navigation.page.html',
  styleUrls: ['navigation.page.scss']
})

export class NavigationPage {

  public progressIsVisible = false;

  public currentBuilding: string;

  public startPinIsVisible = false;

  private floor: FloorMap;
  private route: NaviRoute;
  public startPin: PathNode;

  public availableFloorsBtnIsVisible = false;

//  public testRooms:Room[] = [];
//  public testRoute:PathNode[];

  constructor(private dataService: DataService,
              private modalCtrl: ModalController) {
    this.dataService = dataService;

    // this.testRooms = testDataRooms;
    // this.testRoute = testDataPathNodes;
    // this.testRoute = NavigationPage.testRenderPathNodes();
  }


  private async fetchFloorByLocation(room: string) {
    this.progressIsVisible = true;
    const res = await this.dataService.get_location_search(room).toPromise();
    this.startPin = res.PathNode;
    this.startPinIsVisible = true;
    this.currentBuilding = res.Building;
    await this.fetchFloorByItsNumber(res.Building, res.Floor);
    await this.fetchLocations(res.Building, res.Floor);
    this.displayFloor();
  }

  private async fetchFloorByItsNumber(building:string, floor:string) {
    this.progressIsVisible = true;
    const res = await this.dataService.get_map_floor(building, floor).toPromise();
    this.floor = new FloorMap(res);
  }

  private async fetchRouteToDisplay(start:string, end:string) {
    this.progressIsVisible = true;
    const res1 = await this.dataService.get_location_search(start).toPromise<Location>();
    const res2 = await this.dataService.get_map_floor(res1.Building, res1.Floor).toPromise<MapItem[]>();
    this.currentBuilding = res1.Building;
    this.floor = new FloorMap(res2);
    await this.fetchLocations(res2[0].Building, res2[0].Floor);
    this.dataService.get_route(start, end).subscribe((res3 : ReceivedRoute)=>{
      this.route = new NaviRoute(res3);
      this.displayFloor();
      this.displayNavigationRoute(res2[0].Building, res2[0].Floor);
      this.progressIsVisible = false;
      this.startPinIsVisible = true;
    });
  }

  private displayNavigationRoute(building: string, floor: string){
    if (this.route !=null) {
      this.route.render(building, floor);
    }
  }

  private async fetchLocations(building:string, floor:string) {
    const res = await this.dataService.get_locations(building, floor).toPromise<Location[]>();
    this.floor.locationNames = [];
    for(const l of res) {
      this.floor.locationNames.push({name: l.Name, x: l.PathNode.Coordinate.X, y: l.PathNode.Coordinate.Y});
    }
  }

  private displayFloor() {
    this.floor.renderFloorMap();
    this.progressIsVisible = false;
    this.availableFloorsBtnIsVisible = true;
  }

  public async onSearch(searchInput:string) {
    await this.fetchFloorByLocation(searchInput);
    this.availableFloorsBtnIsVisible = true;
  }

  public async onRoute(routeInput:string[]) {
    await this.fetchRouteToDisplay(routeInput[0], routeInput[1]);
  }

  private isEmptyOrSpaces(str){
    return str === null || str.match(/^ *$/) !== null;
  }

  async presentAvailableFloorModal() {
    this.startPinIsVisible = false;
    this.dataService.get_building(this.currentBuilding).subscribe(async (res: JSON) => {
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
          .then(async (data) => {
            if (data['data']) {
              await this.fetchFloorByItsNumber(this.currentBuilding, data['data']);
              await this.fetchLocations(this.currentBuilding, data['data']);

              this.displayFloor();
              // display route if needed
              const isRouteAvailable = this.route != null;
              if (isRouteAvailable) {
                this.displayNavigationRoute(this.currentBuilding, data['data']);
              }

              this.progressIsVisible = false;
              this.startPinIsVisible = isRouteAvailable;
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