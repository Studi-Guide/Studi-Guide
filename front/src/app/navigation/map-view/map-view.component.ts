import {AfterViewInit, Component, ViewChild} from '@angular/core';
import {DataService} from '../../services/data.service';
import {CanvasResolutionConfigurator} from '../../services/CanvasResolutionConfigurator';
import {Floor, Location, MapItem, PathNode} from '../../building-objects-if';
import {FloorMapRenderer} from './floorMapRenderer';
import {NaviRouteRenderer, ReceivedRoute} from './naviRouteRenderer';
import {IconOnMapRenderer} from '../../services/IconOnMapRenderer';

@Component({
  selector: 'app-map-view',
  templateUrl: './map-view.component.html',
  styleUrls: ['./map-view.component.scss'],
})
export class MapViewComponent implements AfterViewInit {
  private currentBuilding: string;
  private currentRoute:ReceivedRoute;

  private routeRenderer:NaviRouteRenderer;
  private floorMapRenderer:FloorMapRenderer;

  public get CurrentRoute():ReceivedRoute {
    return this.currentRoute;
  }

  public get CurrentBuilding():string {
    return this.currentBuilding;
  }

  constructor(private dataService: DataService) {
  }

  ngAfterViewInit() {
    this.routeRenderer = new NaviRouteRenderer(this.dataService);
    // discovery init
  }

  public async showRoute(start:string, end:string) {
    this.routeRenderer.stopAnimation();

    // Get target location
    const startLocation = await this.dataService.get_location(start).toPromise<Location>();
    this.currentRoute = await this.dataService.get_route(start, end).toPromise();

    await this.renderNavigationPage(startLocation.Building, startLocation.Floor);
  }

  public async showDiscoveryLocation(location:string) {
    this.routeRenderer.stopAnimation();
    const foundLocations = await this.dataService.get_locations_search(location).toPromise();
    if (foundLocations === null || foundLocations.length === 0) {
       throw new Error('Studi-Guide can\'t find ' + location);
    }

    // TODO show multiple locations in future
    const res = foundLocations[0];
    this.currentBuilding = res.Building;
    const items = await this.dataService.get_map_floor(this.currentBuilding, res.Floor).toPromise();
    const locations = await this.dataService.get_locations(res.Building, res.Floor).toPromise<Location[]>();
    const map = this.getCanvasMap(items);
    this.floorMapRenderer = new FloorMapRenderer(items, locations);
    this.floorMapRenderer.renderFloorMap(map);
    this.displayPin(map, res.PathNode);
  }

  public async showFloor(building:string, floor:string) {
    this.routeRenderer.stopAnimation();
    if (this.currentRoute != null) {
      await this.renderNavigationPage(this.currentBuilding, floor);
    }
    else {
      const res = await this.dataService.get_map_floor(building, floor).toPromise();
      const map = this.getCanvasMap(res);
      const locations = await this.dataService.get_locations(this.currentBuilding, floor).toPromise<Location[]>();
      this.floorMapRenderer = new FloorMapRenderer(res, locations);
      this.floorMapRenderer.renderFloorMap(map);
    }
  }

  private async renderNavigationPage(building: string, floor: string) {
    // TODO allow passing a regex to backend to filter map items
    let mapItems = await this.dataService.get_map_floor(building, floor).toPromise<MapItem[]>() ?? new Array<MapItem>();
    let locations = await this.dataService.get_locations(building, floor).toPromise<Location[]>() ?? new Array<Location>();
    for (const routeSection of this.currentRoute.RouteSections) {
      if (routeSection.Floor === floor && routeSection.Building !== building) {
        const items = await this.dataService.get_map_floor(routeSection.Building, routeSection.Floor).toPromise<MapItem[]>();
        mapItems = mapItems.concat(items);

        const locationItems = await this.dataService.get_locations(routeSection.Building, routeSection.Floor).toPromise<Location[]>();
        locations = locations.concat(locationItems);
      }
    }

    const map = this.getCanvasMap(mapItems);
    this.currentBuilding = building;
    this.floorMapRenderer = new FloorMapRenderer(mapItems, locations);
    this.floorMapRenderer.renderFloorMap(map);
    await this.routeRenderer.render(map, this.currentRoute, floor);
    this.routeRenderer.startAnimation();
  }

  private displayPin(map: CanvasRenderingContext2D, pathNode:PathNode) {

    const x = pathNode.Coordinate.X - 15;
    const y = pathNode.Coordinate.Y - 30;
    const iconOnMapRenderer = new IconOnMapRenderer( 'pin-sharp.png');
    iconOnMapRenderer.render(map, x, y, 30, 30);
  }

  public clearMapCanvas() {
    const map = (document.getElementById('map') as HTMLCanvasElement).getContext('2d');
    if (map != null) {
      map.clearRect(0, 0, map.canvas.width, map.canvas.height);
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
