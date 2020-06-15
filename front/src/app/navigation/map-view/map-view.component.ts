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

  private canvasHtmlElement: HTMLCanvasElement;
  private map: CanvasRenderingContext2D;
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
    this.canvasHtmlElement = document.getElementById('map') as HTMLCanvasElement;
    this.map = CanvasResolutionConfigurator.setup(this.canvasHtmlElement);
    this.routeRenderer = new NaviRouteRenderer(this.dataService, this.map);
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
    const res = await this.dataService.get_location(location).toPromise<Location>();
    this.currentBuilding = res.Building;
    const items = await this.dataService.get_map_floor(this.currentBuilding, res.Floor).toPromise();
    const locations = await this.dataService.get_locations(res.Building, res.Floor).toPromise<Location[]>();

    this.floorMapRenderer = new FloorMapRenderer(items, locations, this.map, this.canvasHtmlElement);
    this.floorMapRenderer.renderFloorMap();
    this.displayPin(res.PathNode);

  }

  public async showFloor(building:string, floor:string) {
    this.routeRenderer.stopAnimation();
    const res = await this.dataService.get_map_floor(building, floor).toPromise();
    const locations = await this.dataService.get_locations(this.currentBuilding, floor).toPromise<Location[]>();
    this.floorMapRenderer = new FloorMapRenderer(res, locations, this.map, this.canvasHtmlElement);
    this.floorMapRenderer.renderFloorMap();

    if (this.currentRoute != null) {
      await this.renderNavigationPage(this.currentBuilding, floor);
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

    this.currentBuilding = building;
    this.floorMapRenderer = new FloorMapRenderer(mapItems, locations, this.map, this.canvasHtmlElement);
    this.floorMapRenderer.renderFloorMap();
    await this.routeRenderer.render(this.currentRoute, floor);
    this.routeRenderer.startAnimation();
  }

  private displayPin(pathNode:PathNode) {

    const x = pathNode.Coordinate.X - 15;
    const y = pathNode.Coordinate.Y - 30;
    const iconOnMapRenderer = new IconOnMapRenderer(this.map, 'pin-sharp.png');
    iconOnMapRenderer.render(x, y, 30, 30);
  }

  private clearMapCanvas() {
    if (this.map != null) {
      this.map.clearRect(0, 0, this.canvasHtmlElement.width, this.canvasHtmlElement.height);
    }
  }

}
