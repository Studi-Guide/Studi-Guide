import {AfterViewInit, Component, EventEmitter, Output} from '@angular/core';
import {DataService} from '../../services/data.service';
import {CanvasResolutionConfigurator, TranslationPosition} from '../../services/CanvasResolutionConfigurator';
import {ILocation, IMapItem, IPathNode} from '../../building-objects-if';
import {IconOnMapRenderer} from '../../services/IconOnMapRenderer';
import * as pip from 'point-in-polygon';
import {CanvasTouchHelper} from '../../services/CanvasTouchHelper';
import {IReceivedRoute} from '../../route-objects-if';
import {MapItemRendererCanvas} from './map-item-renderer.canvas';
import {LocationRendererCanvas} from './location-renderer.canvas';
import {RouteRendererCanvas} from './route-renderer.canvas';
import {RendererProvider} from './renderer-provider';

@Component({
  selector: 'app-map-view',
  templateUrl: './map-view.component.html',
  styleUrls: ['./map-view.component.scss'],
})
export class MapViewComponent implements AfterViewInit {
  public currentBuilding: string;
  private currentRoute:IReceivedRoute;
  public currentFloor:string;
  private clickThreshold = 20;

  private renderingContext:CanvasRenderingContext2D;
  private mapItemRenderer:MapItemRendererCanvas[] = [];
  private locationRenderer:LocationRendererCanvas[] = [];
  private routeRenderer:RouteRendererCanvas[] = [];

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

  }

  public async showRoute(route:IReceivedRoute, startLocation:ILocation) {

    this.stopAllAnimations();

    await this.renderNavigationPage(route, startLocation.Building, startLocation.Floor);
  }

  public async showDiscoveryLocation(location:string) {
    this.stopAllAnimations();
    const foundLocations = await this.dataService.get_locations_search(location).toPromise();
    if (foundLocations === null || foundLocations.length === 0) {
       throw new Error('Studi-Guide can\'t find ' + location);
    }

    // TODO show multiple locations in future
    const res = foundLocations[0];
    this.currentBuilding = res.Building;
    const items = await this.dataService.get_map_floor(this.currentBuilding, res.Floor).toPromise();
    this.mapItemRenderer = RendererProvider.GetMapItemRendererCanvas(...items);

    const locations = await this.dataService.get_locations(res.Building, res.Floor).toPromise<ILocation[]>();
    this.locationRenderer = RendererProvider.GetLocationRendererCanvas(...locations);

    // TODO shift map got get res.PathNode into focus
    this.clearMapCanvas();
    this.createNewCanvasMap(0, 0);
    this.renderMapItems();
    this.renderLocations();
    this.displayPin(res.PathNode);
    this.currentFloor = res.Floor;
    return res;
  }

  public async showFloor(floor:string, building:string) {
    this.stopAllAnimations();
    if (this.currentRoute != null) {
      await this.renderNavigationPage(this.currentRoute, this.currentBuilding, floor);
    }
    else {
      const res = await this.dataService.get_map_items('', floor, building).toPromise()
      this.mapItemRenderer = RendererProvider.GetMapItemRendererCanvas(...res);
      this.createNewCanvasMap(0,0);

      const locations = await this.dataService.get_locations_items('', floor, building).toPromise();
      this.locationRenderer = RendererProvider.GetLocationRendererCanvas(...locations);

      this.renderMapItems();
      this.renderLocations();
    }
    this.currentFloor = floor;
  }

  public async showDiscoveryMap(campus:string, floor: string) {
      const items = await this.dataService.get_map_items(campus, floor, '').toPromise();
      this.mapItemRenderer = RendererProvider.GetMapItemRendererCanvas(...items);

      const locations = await this.dataService.get_locations_items(campus, floor, '').toPromise();
      this.locationRenderer = RendererProvider.GetLocationRendererCanvas(...locations);

      this.createNewCanvasMap(0,0);

      this.renderMapItems();
      this.renderLocations();

      this.currentFloor = floor;
  }

  public clearMapCanvas() {
    const map = (document.getElementById('map') as HTMLCanvasElement).getContext('2d');
    if (map != null) {
      map.clearRect(0, 0, map.canvas.width, map.canvas.height);
    }
  }

  public async onClickTouch(event:MouseEvent) {

    const coordinate = CanvasTouchHelper.transformInOriginCoordinate({
      x: event.clientX, y:event.clientY}, event.target as HTMLCanvasElement);
    const point = [coordinate.x, coordinate.y];
    if(this.currentRoute != null) {
      const items: IMapItem[] = [];
      for (const m of this.getInteractiveStairWellMapItemRenderer(this.currentFloor)) {
        items.push(m.MapItem);
      }

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
    for (const l of this.locationRenderer) {
      const location = l.Location;
      if (Math.abs(location.PathNode.Coordinate.X - point[0]) < this.clickThreshold
          && Math.abs(location.PathNode.Coordinate.Y - point[1]) < this.clickThreshold) {
        this.locationClick.emit(location);
        return;
      }
    }
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

    this.mapItemRenderer = RendererProvider.GetMapItemRendererCanvas(...mapItems);
    this.locationRenderer = RendererProvider.GetLocationRendererCanvas(...locations);
    this.routeRenderer = RendererProvider.GetRouteRendererCanvas(route);

    this.createNewCanvasMap(0, 0);
    this.currentBuilding = building;

    this.renderMapItems();
    this.renderLocations();

    this.renderRoutes({floor});
    // TODO animate route call here

    for(const m of this.getInteractiveStairWellMapItemRenderer(floor)) {
      m.startAnimation(this.renderingContext, {renderer: this.routeRenderer, floor});
    }

    this.currentRoute = route;
    this.currentFloor = floor;
  }

  private createNewCanvasMap(positionX: number, positionY: number, scale: number = 1) {
    const mapCanvas = document.getElementById('map') as HTMLCanvasElement;
    let mapHeightNeeded = 0;
    let mapWidthNeeded = 0;
    for (const m of this.mapItemRenderer) {
      const mapItem = m.MapItem;
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
    this.renderingContext = CanvasResolutionConfigurator.setup(mapCanvas, mapWidthNeeded, mapHeightNeeded,scale, position);
  }

  private displayPin(pathNode:IPathNode) {
    const x = pathNode.Coordinate.X - 15;
    const y = pathNode.Coordinate.Y - 30;
    const iconOnMapRenderer = new IconOnMapRenderer( 'pin-sharp.png');
    iconOnMapRenderer.render(this.renderingContext, x, y, 30, 30);
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

  private getInteractiveStairWellMapItemRenderer(floor:string) {
    const pNodes:IPathNode[] = [];

    for (const r of this.routeRenderer) {
      for (let i = 0; i < r.Route.RouteSections.length-1; i++) {
        if (r.Route.RouteSections[i].Building !== r.Route.RouteSections[i+1].Building)
          continue;
        pNodes.push(r.Route.RouteSections[i].Route[r.Route.RouteSections[i].Route.length-1]);
      }
    }

    const tmpMItems:MapItemRendererCanvas[] = [];
    for (const pNode of pNodes) {
      for (const m of this.mapItemRenderer) {
        if (m.MapItem.Floor === floor && m.MapItem.PathNodes.filter(p => p.Id === pNode.Id).length > 0) {
          tmpMItems.push(m);
        }
      }
    }

    return tmpMItems;
  }

  private stopAllAnimations() {
    for (const c of this.routeRenderer)
      c.stopAnimation(this.renderingContext);

    for (const m of this.mapItemRenderer)
      m.stopAnimation(this.renderingContext);

    for (const l of this.locationRenderer)
      l.stopAnimation(this.renderingContext);
  }

  private renderMapItems() {
    for (const c of this.mapItemRenderer)
      c.render(this.renderingContext);
  }

  private renderLocations() {
    for (const l of this.locationRenderer)
      l.render(this.renderingContext);
  }

  private renderRoutes(args:any) {
    for (const r of this.routeRenderer)
      r.render(this.renderingContext, args);
  }
}
