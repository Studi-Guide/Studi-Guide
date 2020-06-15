import {BuildingData, Location, MapItem, PathNode} from '../building-objects-if';
// import {testDataRooms, testDataPathNodes} from './test-building-data';
import {Component} from '@angular/core';
import {ModalController} from '@ionic/angular';
import {DataService} from '../services/data.service';
import {FloorMap} from './floorMap';
import {NaviRoute, ReceivedRoute} from './naviRoute';
import {AvailableFloorsPage} from '../available-floors/available-floors.page';
import {CanvasResolutionConfigurator} from '../services/CanvasResolutionConfigurator';

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

  constructor(private dataService: DataService,
              private modalCtrl: ModalController) {
    this.dataService = dataService;
  }

  private async fetchFloorByLocation(res: Location) {
    this.currentBuilding = res.Building;
    const mapItems = await this.fetchFloorByItsNumber(res.Building, res.Floor);
    const locations = await this.dataService.get_locations(res.Building, res.Floor).toPromise<Location[]>();
    const map = this.getCanvasMap(mapItems);

    this.floor = this.createMap(map, mapItems, locations);
    this.displayPin(map, this.floor, res.PathNode);
  }

  private async fetchFloorByItsNumber(building: string, floor: string) {
    this.progressIsVisible = true;
    return await this.dataService.get_map_floor(building, floor).toPromise();
  }

  private async fetchRouteToDisplay(start: string, end: string) {
    this.progressIsVisible = true;
    // Get target location
    const endLocation = await this.dataService.get_location(end).toPromise<Location>();
    this.route = new NaviRoute(await this.dataService.get_route(start, end).toPromise());

    await this.RenderNavigationPage(this.route, endLocation.Building, endLocation.Floor);
    this.progressIsVisible = false;
  }

  private async RenderNavigationPage(route: NaviRoute, building: string, floor: string) {
    // TODO allow passing a regex to backend to filter map items
    let mapItems = await this.dataService.get_map_floor(building, floor).toPromise<MapItem[]>() ?? new Array<MapItem>();
    let locations = await this.dataService.get_locations(building, floor).toPromise<Location[]>() ?? new Array<Location>();
    for (const routeSection of route.route.RouteSections) {
      if (routeSection.Floor === floor && routeSection.Building !== building) {
        const items = await this.dataService.get_map_floor(routeSection.Building, routeSection.Floor).toPromise<MapItem[]>();
        mapItems = mapItems.concat(items);

        const locationItems = await this.dataService.get_locations(routeSection.Building, routeSection.Floor).toPromise<Location[]>();
        locations = locations.concat(locationItems);
      }
    }
    const map = this.getCanvasMap(mapItems)
    this.currentBuilding = building;
    this.createMap(map, mapItems, locations);
    this.displayNavigationRoute(map, floor);
  }

  private displayPin(map: CanvasRenderingContext2D, floor: FloorMap, startPin: PathNode) {
    const x = startPin.Coordinate.X - 15;
    const y = startPin.Coordinate.Y - 30;
    floor.renderStartPin(map, x, y, 30, 30);
  }

  private displayNavigationRoute(map: CanvasRenderingContext2D, floor: string) {
    if (this.route != null) {
      this.route.render(map, floor);
    }
  }

  private createMap(map:CanvasRenderingContext2D, mapItems: MapItem[], locations: Location[]) {
      const page = new FloorMap(mapItems);
      page.locationNames = [];
      for (const l of locations) {
        page.locationNames.push({name: l.Name, x: l.PathNode.Coordinate.X, y: l.PathNode.Coordinate.Y});
      }

      page.renderFloorMap(map);
      return page;
  }

  public async onDiscovery(searchInput: string) {
    const locations = await this.dataService.get_locations_search(searchInput).toPromise();
    // TODO present all discovered locations
    if (locations != null && locations.length === 1) {
      this.progressIsVisible = true;
      await this.fetchFloorByLocation(locations[0]);
      this.progressIsVisible = false;
    }

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
      for (const routeSection of this.route.route.RouteSections) {
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
        const mapItems = await this.fetchFloorByItsNumber(this.currentBuilding, data.data);
        const locations = await this.dataService.get_locations(this.currentBuilding, data.data).toPromise<Location[]>();

        const map = this.getCanvasMap(mapItems)

        this.createMap(map, mapItems, locations);

        // display route if needed
        const isRouteAvailable = this.route != null;
        if (isRouteAvailable) {
          this.displayNavigationRoute(map, data.data);
        }
      } else {
        await this.RenderNavigationPage(this.route, this.currentBuilding, data.data);
      }

      this.progressIsVisible = false;
    }
  }

  private getCanvasMap(mapItems: MapItem[]) {
    const mapCanvas = document.getElementById('map') as HTMLCanvasElement;
    let mapHeightNeeded = 0;
    let mapWidthNeeded = 0;
    for (const mapItem of mapItems) {
      if (mapItem.Sections != null) {
        for (const section of mapItem.Sections) {
          if (section.End.X > mapWidthNeeded) {
            mapWidthNeeded = section.End.X;
          }
          if (section.End.Y > mapHeightNeeded) {
            mapHeightNeeded = section.End.Y;
          }
        }
      }
    }

    return CanvasResolutionConfigurator.setup(mapCanvas, mapWidthNeeded, mapHeightNeeded);
  }
}
