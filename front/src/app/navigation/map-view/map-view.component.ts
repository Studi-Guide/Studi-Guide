import {AfterViewInit, Component, EventEmitter, Output, ViewChild} from '@angular/core';
import {DataService} from '../../services/data.service';
import {CanvasResolutionConfigurator, TranslationPosition} from '../../services/CanvasResolutionConfigurator';
import {ILocation, IMapItem, IPathNode} from '../../building-objects-if';
import {FloorMapRenderer} from './floorMapRenderer';
import {NaviRouteRenderer} from './naviRouteRenderer';
import {IconOnMapRenderer} from '../../services/IconOnMapRenderer';
import * as pip from 'point-in-polygon';
import {CanvasTouchHelper} from '../../services/CanvasTouchHelper';
import {IReceivedRoute} from "../../route-objects-if";

@Component({
  selector: 'app-map-view',
  templateUrl: './map-view.component.html',
  styleUrls: ['./map-view.component.scss'],
})
export class MapViewComponent implements AfterViewInit {
  private currentBuilding: string;
  private currentRoute:IReceivedRoute;
  private currentFloor:string;
  private clickThreshold = 20;
  private routeRenderer:NaviRouteRenderer;
  public floorMapRenderer:FloorMapRenderer;

  @Output() locationClick = new EventEmitter<ILocation>();

  public get CurrentRoute():IReceivedRoute {
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

  public async showRoute(route:IReceivedRoute, startLocation:ILocation) {
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
    const locations = await this.dataService.get_locations(res.Building, res.Floor).toPromise<ILocation[]>();

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

  private async renderNavigationPage(route:IReceivedRoute, building: string, floor: string) {
    // TODO allow passing a regex to backend to filter map items
    let mapItems = await this.dataService.get_map_floor(building, floor).toPromise<IMapItem[]>() ?? new Array<IMapItem>();
    let locations = await this.dataService.get_locations(building, floor).toPromise<ILocation[]>() ?? new Array<ILocation>();
    for (const routeSection of route.RouteSections) {
      if (routeSection.Floor === floor && routeSection.Building !== building) {
        const items = await this.dataService.get_map_floor(routeSection.Building, routeSection.Floor).toPromise<IMapItem[]>();
        mapItems = mapItems.concat(items);

        const locationItems = await this.dataService.get_locations(routeSection.Building, routeSection.Floor).toPromise<ILocation[]>();
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

  private displayPin(map: CanvasRenderingContext2D, pathNode:IPathNode) {

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

  private getCanvasMap(mapItems: IMapItem[], positionX: number, positionY: number) {
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
    return CanvasResolutionConfigurator.setup(mapCanvas, mapWidthNeeded, mapHeightNeeded,1, position);
  }

  public async onClickTouch(event:MouseEvent) {

    const point = CanvasTouchHelper.CalculateXY(event, event.currentTarget as HTMLElement);

    if(this.currentRoute != null) {
      const items: IMapItem[] = await this.routeRenderer.getInteractiveStairWellMapItems(this.currentRoute, this.currentFloor);

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
    const locations:ILocation[] = this.floorMapRenderer.locationNames
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

  private async showNextFloor(item: IMapItem) {
    for (let i = 0; i < this.currentRoute.RouteSections.length-1; i++) {
      if (this.currentRoute.RouteSections[i].Floor === item.Floor && this.currentRoute.RouteSections[i].Building === item.Building) {
        this.currentFloor = this.currentRoute.RouteSections[i+1].Floor;
        this.currentBuilding = this.currentRoute.RouteSections[i+1].Building;
        await this.showFloor(this.currentFloor, this.currentBuilding);
        return;
      }
    }
  }


  private draw(scale:number, translatePosition) {

  }
}
