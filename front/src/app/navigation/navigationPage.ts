import {BuildingData, Location, MapItem, PathNode} from '../building-objects-if';
// import {testDataRooms, testDataPathNodes} from './test-building-data';
import {Component} from '@angular/core';
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
  public availableFloorsBtnIsVisible = false;

  public currentBuilding: string;

  private floor: FloorMap;
  private route: NaviRoute;

  public startPin: PathNode;

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
    this.currentBuilding = res.Building;
    await this.fetchFloorByItsNumber(res.Building, res.Floor);

    const locations = await this.dataService.get_locations(res.Building, res.Floor).toPromise<Location[]>();
    this.displayLocations(locations);
    this.displayFloor();
    this.displayPin();
  }

  private async fetchFloorByItsNumber(building: string, floor: string) {
    this.progressIsVisible = true;
    const res = await this.dataService.get_map_floor(building, floor).toPromise();
    this.floor = new FloorMap(res);
  }

  private async fetchRouteToDisplay(start: string, end: string) {
    this.progressIsVisible = true;
    // Get target location
    const endLocation = await this.dataService.get_location_search(end).toPromise<Location>();
    this.route = new NaviRoute(await this.dataService.get_route(start, end).toPromise());

    await this.RenderNavigationPage(this.route, endLocation.Building, endLocation.Floor);
    this.progressIsVisible = false;
  }

  private async RenderNavigationPage(route: NaviRoute, building: string, floor: string) {
    // TODO allow passing a regex to backend to filter map items
    let mapItems = await this.dataService.get_map_floor(building, floor).toPromise<MapItem[]>() ?? new Array<MapItem>();
    let locations = await this.dataService.get_locations(building, floor).toPromise<Location[]>() ?? new Array<Location>();
    for (const routeSection of route.routeSections) {
      if (routeSection.Floor === floor && routeSection.Building !== building) {
        const items = await this.dataService.get_map_floor(routeSection.Building, routeSection.Floor).toPromise<MapItem[]>();
        mapItems = mapItems.concat(items);

        const locationItems = await this.dataService.get_locations(routeSection.Building, routeSection.Floor).toPromise<Location[]>();
        locations = locations.concat(locationItems);
      }
    }

    this.currentBuilding = building;
    this.floor = new FloorMap(mapItems);
    this.displayLocations(locations);
    this.displayFloor();
    this.displayNavigationRoute(floor);
  }

  private displayPin() {
    const x = this.startPin.Coordinate.X - 15;
    const y = this.startPin.Coordinate.Y - 30;
    this.floor.pin.render(x, y, 30, 30);
  }

  private displayNavigationRoute(floor: string) {
    if (this.route != null) {
      this.route.render(floor);
    }
  }

  private displayLocations(locations: Location[]) {
    this.floor.locationNames = [];
    for (const l of locations) {
      this.floor.locationNames.push({name: l.Name, x: l.PathNode.Coordinate.X, y: l.PathNode.Coordinate.Y});
    }
  }

  private displayFloor() {
    this.floor.renderFloorMap();
    this.progressIsVisible = false;
    this.availableFloorsBtnIsVisible = true;
  }

  public async onDiscovery(searchInput: string) {
    await this.fetchFloorByLocation(searchInput);
    this.availableFloorsBtnIsVisible = true;
  }

  public async onRoute(routeInput: string[]) {
    await this.fetchRouteToDisplay(routeInput[0], routeInput[1]);
  }

  private isEmptyOrSpaces(str) {
    return str === null || str.match(/^ *$/) !== null;
  }

  async presentAvailableFloorModal() {
    let floors = new Array<string>();
    if (this.route == null) {
      const building = await this.dataService.get_building(this.currentBuilding).toPromise<BuildingData>();
      floors = floors.concat(building.Floors);
    } else {
      // get all floors from all buildings on the route
      for (const routeSection of this.route.routeSections) {
        const building = await this.dataService.get_building(routeSection.Building).toPromise<BuildingData>();
        floors = floors.concat(building.Floors);
      }

      // distinct array
      floors = floors.filter((n, i) => floors.indexOf(n) === i);
    }

    const availableFloorModal = await this.modalCtrl.create({
      component: AvailableFloorsPage,
      cssClass: 'floor-modal',
      componentProps: {
        floors
      }
    })
    await availableFloorModal.present();

    const data = await availableFloorModal.onDidDismiss()
    if (data.data) {
      if (this.route == null) {
        await this.fetchFloorByItsNumber(this.currentBuilding, data.data);
        const locations = await this.dataService.get_locations(this.currentBuilding, data.data).toPromise<Location[]>();
        await this.displayLocations(locations);

        this.displayFloor();
        // display route if needed
        const isRouteAvailable = this.route != null;
        if (isRouteAvailable) {
          this.displayNavigationRoute(data.data);
        }
      } else {
        await this.RenderNavigationPage(this.route, this.currentBuilding, data.data);
      }

      this.progressIsVisible = false;
    }
  }
}
