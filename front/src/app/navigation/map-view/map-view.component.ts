import {AfterViewInit, Component, EventEmitter, Output, ViewChild} from '@angular/core';
import {DataService} from '../../services/data.service';
import {CanvasResolutionConfigurator, TranslationPosition} from '../../services/CanvasResolutionConfigurator';
import {Floor, Location, MapItem, PathNode} from '../../building-objects-if';
import {FloorMapRenderer} from './floorMapRenderer';
import {NaviRouteRenderer, ReceivedRoute} from './naviRouteRenderer';
import {IconOnMapRenderer} from '../../services/IconOnMapRenderer';
import * as pip from 'point-in-polygon';
import {CanvasTouchHelper} from '../../services/CanvasTouchHelper';

@Component({
  selector: 'app-map-view',
  templateUrl: './map-view.component.html',
  styleUrls: ['./map-view.component.scss'],
})
export class MapViewComponent implements AfterViewInit {
  private currentBuilding: string;
  private currentRoute:ReceivedRoute;
  private currentFloor:string;
  private clickThreshold = 20;
  private routeRenderer:NaviRouteRenderer;
  public floorMapRenderer:FloorMapRenderer;

  @Output() locationClick = new EventEmitter<Location>();

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
  }

  public async showRoute(route:ReceivedRoute, startLocation:Location) {
    this.routeRenderer.stopAnimation();
    await this.renderNavigationPage(route, startLocation.Building, startLocation.Floor);
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

    // TODO shift map got get res.PathNode into focus
    const map = this.getCanvasMap(items, 0, 0);
    this.floorMapRenderer = new FloorMapRenderer(items, locations);
    this.floorMapRenderer.renderFloorMap(map);
    this.displayPin(map, res.PathNode);
    this.currentFloor = res.Floor;
    return res;
  }

  public async showFloor(floor:string, building:string) {
    this.routeRenderer.stopAnimation();
    if (this.currentRoute != null) {
      await this.renderNavigationPage(this.currentRoute, this.currentBuilding, floor);
    }
    else {
      const res = await this.dataService.get_map_items('', floor, building).toPromise()
      const map = this.getCanvasMap(res, 0, 0);
      const locations = await this.dataService.get_locations_items('', floor, building).toPromise();
      this.floorMapRenderer = new FloorMapRenderer(res, locations);
      this.floorMapRenderer.renderFloorMap(map);
    }
    this.currentFloor = floor;
  }

  public async showDiscoveryMap(campus:string, floor: string) {
      const items = await this.dataService.get_map_items(campus, floor, '').toPromise();
      const locations = await this.dataService.get_locations_items(campus, floor, '').toPromise();
      const map = this.getCanvasMap(items, 0,0);
      this.floorMapRenderer = new FloorMapRenderer(items, locations);
      this.floorMapRenderer.renderFloorMap(map);
      this.currentFloor = floor;
  }

  private async renderNavigationPage(route:ReceivedRoute, building: string, floor: string) {
    // TODO allow passing a regex to backend to filter map items
    let mapItems = await this.dataService.get_map_floor(building, floor).toPromise<MapItem[]>() ?? new Array<MapItem>();
    let locations = await this.dataService.get_locations(building, floor).toPromise<Location[]>() ?? new Array<Location>();
    for (const routeSection of route.RouteSections) {
      if (routeSection.Floor === floor && routeSection.Building !== building) {
        const items = await this.dataService.get_map_floor(routeSection.Building, routeSection.Floor).toPromise<MapItem[]>();
        mapItems = mapItems.concat(items);

        const locationItems = await this.dataService.get_locations(routeSection.Building, routeSection.Floor).toPromise<Location[]>();
        locations = locations.concat(locationItems);
      }
    }

    const map = this.getCanvasMap(mapItems, 0, 0);
    this.currentBuilding = building;
    this.floorMapRenderer = new FloorMapRenderer(mapItems, locations);
    this.floorMapRenderer.renderFloorMap(map);
    await this.routeRenderer.render(map, route, floor);
    this.routeRenderer.startAnimation();
    this.currentRoute = route;
    this.currentFloor = floor;
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

  private getCanvasMap(mapItems: MapItem[], positionX: number, positionY: number, scale: number = 1) {
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

    // increase map size
    mapWidthNeeded += Math.abs(positionX);
    mapHeightNeeded += Math.abs(positionY);
    const position = new TranslationPosition();
    position.X = positionX;
    position.Y = positionY;
    return CanvasResolutionConfigurator.setup(mapCanvas, mapWidthNeeded, mapHeightNeeded,scale, position);
  }

  public async onClickTouch(event:MouseEvent) {

    const point = CanvasTouchHelper.CalculateXY({x: event.clientX, y:event.clientY});

    if(this.currentRoute != null) {
      const items: MapItem[] = await this.routeRenderer.getInteractiveStairWellMapItems(this.currentRoute, this.currentFloor);

      for (const mapItem of items) {
        const polygon = [];
        for (const section of mapItem.Sections) {
          polygon.push([section.Start.X, section.Start.Y]);
        }
        if (pip(point, polygon)) {
          await this.showNextFloor(mapItem);
          return;
        }
      }
    }
    // Track clicks/touches on locations
    const locations:Location[] = this.floorMapRenderer.locationNames
    if (locations != null) {
      for (const location of locations) {
        if (Math.abs(location.PathNode.Coordinate.X - point[0]) < this.clickThreshold
            && Math.abs(location.PathNode.Coordinate.Y - point[1]) < this.clickThreshold) {
          this.locationClick.emit(location);
          return;
        }
      }
    }
  }

  private async showNextFloor(item: MapItem) {
    for (let i = 0; i < this.currentRoute.RouteSections.length-1; i++) {
      if (this.currentRoute.RouteSections[i].Floor === item.Floor && this.currentRoute.RouteSections[i].Building === item.Building) {
        this.currentFloor = this.currentRoute.RouteSections[i+1].Floor;
        this.currentBuilding = this.currentRoute.RouteSections[i+1].Building;
        await this.showFloor(this.currentFloor, this.currentBuilding);
        return;
      }
    }
  }
}
