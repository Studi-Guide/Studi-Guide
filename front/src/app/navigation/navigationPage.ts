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
  @ViewChild('discoverySearchbar') discoverySearchbarRef;

  public progressIsVisible = false;
  public routeInputIsVisible = false;
  public searchBtnIsVisible = true;
  public closeRouteBtnIsVisible = false;

  public startInput: string;
  public destinationInput: string;

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

  public async showFloorForSearch() {
    if (this.routeInputIsVisible) {
      this.hideRouteSearchbar();
    } else if (this.startInput !== undefined && this.startInput !== '' && this.startInput != null) {
      await this.fetchFloorByLocation(this.startInput);
      this.availableFloorsBtnIsVisible = true;
    }
  }

  private async fetchFloorByLocation(room: string) {
    this.progressIsVisible = true;
    const res = await this.dataService.get_location_search(room).toPromise();
    this.startPin = res.PathNode;
    this.startPinIsVisible = true;
    await this.fetchFloorByItsNumber(res.Building, res.Floor);
    await this.fetchLocations(res.Building, res.Floor);
    this.displayFloor();
  }

  public async showRoute() {
    if (!this.routeInputIsVisible) {
      this.routeInputIsVisible = true;
      const searchbars = document.querySelector('ion-item');
      searchbars.setAttribute('color', 'primary');
      this.searchBtnIsVisible = false;
      this.closeRouteBtnIsVisible = true;
    } else if (this.startInput !== undefined && this.destinationInput !== undefined
        && this.startInput !== '' && this.destinationInput !== ''
        && this.startInput != null && this.destinationInput != null
    ) {
      this.availableFloorsBtnIsVisible = true;
      await this.fetchRouteToDisplay(this.startInput, this.destinationInput);
    }
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

  public async checkWhatIsRequestedByEnterKey() {
    if (this.startInput !== undefined && this.startInput !== '' && this.startInput !== null &&
        !this.routeInputIsVisible
    ) {
      await this.showFloorForSearch();
    } else if (this.startInput !== undefined && this.startInput !== '' && this.startInput !== null &&
        this.routeInputIsVisible &&
        (this.destinationInput === undefined || this.destinationInput === '' || this.destinationInput === null)
    ) {
      // TODO check setFocus on Android, iOS, etc.
      this.discoverySearchbarRef.setFocus();
      await this.showFloorForSearch();
    } else if (this.startInput !== undefined && this.startInput !== '' && this.startInput != null &&
        this.routeInputIsVisible &&
        this.destinationInput !== undefined && this.destinationInput !== '' && this.destinationInput !== null
    ) {
      await this.showRoute();
    }
  }

  public hideRouteSearchbar() {
    this.routeInputIsVisible = false;
    const searchbars = document.querySelector('ion-item');
    searchbars.setAttribute('color', 'light-tint');
    this.searchBtnIsVisible = true;
    this.closeRouteBtnIsVisible = false;
  }

  private isEmptyOrSpaces(str){
    return str === null || str.match(/^ *$/) !== null;
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
          .then(async (data) => {
            if (data['data']) {
              const building = this.startInput.slice(0, 2);
              await this.fetchFloorByItsNumber(building, data['data']);
              await this.fetchLocations(building, data['data']);
              // display route if needed
              const isRouteAvailable = this.route != null;
              if (isRouteAvailable) {
                this.displayNavigationRoute(building, data['data']);
              }

              this.displayFloor();
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