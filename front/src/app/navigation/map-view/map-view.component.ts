import {AfterViewInit, Component, EventEmitter, Output} from '@angular/core';
import {DataService} from '../../services/data.service';
import {CanvasResolutionConfigurator} from '../../services/CanvasResolutionConfigurator';
import {ILocation, IMapItem, IPathNode} from '../../building-objects-if';
import {IconOnMapRenderer} from '../../services/IconOnMapRenderer';
import * as pip from 'point-in-polygon';
import {IReceivedRoute} from '../../route-objects-if';
import {MapItemRendererCanvas} from './map-item-renderer.canvas';
import {LocationRendererCanvas} from './location-renderer.canvas';
import {RouteRendererCanvas} from './route-renderer.canvas';
import {RendererProvider} from './renderer-provider';
import {NavigationPage} from '../navigation.page';
import panzoom, {PanZoom} from 'panzoom';
import {CanvasTouchHelper} from '../../services/CanvasTouchHelper';

@Component({
  selector: 'app-map-view',
  templateUrl: './map-view.component.html',
  styleUrls: ['./map-view.component.scss'],
})
export class MapViewComponent implements AfterViewInit {
  public static DISTANCE_SCALAR = 5.0;

  public get CurrentRoute(): IReceivedRoute {
    return this.currentRoute;
  }

  public get CurrentBuilding(): string {
    return this.currentBuilding;
  }

  constructor(private dataService: DataService) {
  }
  public currentBuilding: string;
  private currentRoute: IReceivedRoute;
  public currentFloor: string;
  private clickThreshold = 20;

  private renderingContext: CanvasRenderingContext2D;
  private mapItemRenderer: MapItemRendererCanvas[] = [];
  private locationRenderer: LocationRendererCanvas[] = [];
  private routeRenderer: RouteRendererCanvas[] = [];
  panZoomController: PanZoom;

  @Output() locationClick = new EventEmitter<ILocation>();
  @Output() mapScroll = new EventEmitter<any>();
  @Output() floorChanged =  new EventEmitter<any>();

  private static calculateMovePosition(x: number, y: number) {
    const element = document.getElementById('canvas-wrapper');
    const parentElement = element.parentElement.parentElement;
    const isBigDevice = window.matchMedia('(min-width: 1200px)').matches;

    const height = isBigDevice ? parentElement.clientHeight / 2 : 2 * parentElement.clientHeight / 5;
    const width = element.clientWidth / (isBigDevice ? (3.0 / 2.0) : 2.0);

    const targetY =  isBigDevice ? height - y : height - Math.min(y, parentElement.clientHeight);
    const targetX = width - x;
    return {x: targetX, y: targetY};
  }

  ngAfterViewInit() {
    const element: MapViewComponent = this;
    if (!this.panZoomController) {
      this.panZoomController = panzoom(document.getElementById('map'),
          {
            maxZoom: 2.0,
            minZoom: 0.25,
            initialZoom: 0.7,
            bounds: true,
            boundsPadding: 0.1,
            // Enable touch recognition on child events
            async onTouch(e) {
              if (e.touches.length === 1 ){
                console.log(e.touches[0]);
                await element.onElementClick(e.touches[0].clientX, e.touches[0].clientY, e.target as HTMLElement);
              }

              return false;
            }
          });
    }
  }

  public async showRoute(route: IReceivedRoute, startLocation: ILocation) {
    this.stopAllAnimations();
    await this.renderNavigationPage(route, startLocation.Building, startLocation.Floor);
  }

  public async showDiscoveryLocation(location: string) {
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

    this.clearMapCanvas();
    this.RefreshMap();
    this.displayPin(res.PathNode);
    this.currentFloor = res.Floor;
    this.CenterMap(res.PathNode.Coordinate.X, res.PathNode.Coordinate.Y);
    return res;
  }

  public async showFloor(floor: string, building: string) {
    this.stopAllAnimations();
    if (this.currentRoute != null) {
      await this.renderNavigationPage(this.currentRoute, this.currentBuilding, floor);
    } else {
      const res = await this.dataService.get_map_items('', floor, building).toPromise();
      this.mapItemRenderer = RendererProvider.GetMapItemRendererCanvas(...res);
      const locations = await this.dataService.get_locations_items('', floor, building).toPromise();
      this.locationRenderer = RendererProvider.GetLocationRendererCanvas(...locations);
      this.RefreshMap();
    }
    this.currentFloor = floor;
    this.currentBuilding = building;
  }

  public async showDiscoveryMap(campus: string, floor: string) {
      const items = await this.dataService.get_map_items(campus, floor, '').toPromise();
      this.mapItemRenderer = RendererProvider.GetMapItemRendererCanvas(...items);

      const locations = await this.dataService.get_locations_items(campus, floor, '').toPromise();
      this.locationRenderer = RendererProvider.GetLocationRendererCanvas(...locations);

      this.RefreshMap();
      this.currentFloor = floor;
      this.FitMap();
  }

  public clearMapCanvas() {
    const map = (document.getElementById('map') as HTMLCanvasElement).getContext('2d');
    if (map != null) {
      map.clearRect(0, 0, map.canvas.width, map.canvas.height);
    }
  }

  public async onClick(event: MouseEvent) {
    await this.onElementClick(event.clientX, event.clientY, event.target as HTMLElement);
  }

  public async onElementClick(clientX: number, clientY: number, target: HTMLElement) {
    const transform = this.panZoomController.getTransform();
    const coordinate = CanvasTouchHelper.transformInOriginCoordinate({
      x: clientX, y: clientY
    }, transform.scale, target as HTMLElement);
    const point = [coordinate.x, coordinate.y];
    if (this.currentRoute != null) {
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
          await this.showNextFloor(mapItem, false);
          return;
        }
      }
    }
    // Track clicks/touches on locations
    const clickedLocations: ILocation[] = [];
    for (const l of this.locationRenderer) {
      const location = l.Location;

      if (Math.abs(location.PathNode.Coordinate.X - point[0]) < this.clickThreshold
          && Math.abs(location.PathNode.Coordinate.Y - point[1]) < this.clickThreshold) {
        clickedLocations.push(location);
      }
    }

    if (clickedLocations.length > 0) {
      // nearest Location will be emitted
      const sortedLocations = clickedLocations.length === 1 ?
          clickedLocations :
          clickedLocations.sort((n1, n2) => {
        if (Math.abs(n1.PathNode.Coordinate.X - point[0]) + Math.abs(n1.PathNode.Coordinate.Y - point[1]) >
            Math.abs(n2.PathNode.Coordinate.X - point[0]) + Math.abs(n2.PathNode.Coordinate.Y - point[1])) {
          return 1;
        } else {
          return -1;
        }
      });

      if (sortedLocations[0]?.PathNode){
        this.RefreshMap();
        this.displayPin(sortedLocations[0].PathNode);
      }

      this.locationClick.emit(sortedLocations[0]);
    }
  }

  private async renderNavigationPage(route: IReceivedRoute, building: string, floor: string) {
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

    this.currentBuilding = building;
    this.RefreshMap();
    this.renderRoutes({floor});
    // TODO animate route call here

    for (const m of this.getInteractiveStairWellMapItemRenderer(floor)) {
      m.startAnimation(this.renderingContext, {renderer: this.routeRenderer, floor});
    }

    this.currentRoute = route;
    this.currentFloor = floor;
  }

  private createNewCanvasMap() {
    const mapCanvas = document.getElementById('map') as HTMLCanvasElement;
    let mapHeightNeeded = 0;
    let mapWidthNeeded = 0;
    for (const m of this.mapItemRenderer) {
      const mapItem = m.MapItem;
      if (mapItem.Sections) {
        for (const section of mapItem.Sections) {
          if (section.End.X > mapWidthNeeded) {
            mapWidthNeeded = section.End.X;
          }
          if (section.End.Y > mapHeightNeeded) {
            mapHeightNeeded = section.End.Y;
          }
        }
      }

      if (mapItem.PathNodes) {
        for (const node of mapItem.PathNodes) {
          if (node.Coordinate.X > mapWidthNeeded) {
            mapWidthNeeded = node.Coordinate.X;
          }
          if (node.Coordinate.Y > mapHeightNeeded) {
            mapHeightNeeded = node.Coordinate.Y;
          }
        }
      }
    }

    // increase map size
    this.renderingContext = CanvasResolutionConfigurator.setup(mapCanvas, mapWidthNeeded, mapHeightNeeded);
  }

  private displayPin(pathNode: IPathNode) {
    const x = pathNode.Coordinate.X - 30;
    const y = pathNode.Coordinate.Y - 40;
    const iconOnMapRenderer = new IconOnMapRenderer( 'assets/pin-red.png');
    iconOnMapRenderer.render(this.renderingContext, x, y, 60, 60);
  }

  private async showNextFloor(item: IMapItem, centerMap: boolean) {
    for (let i = 0; i < this.currentRoute.RouteSections.length - 1; i++) {
      if (this.currentRoute.RouteSections[i].Floor === item.Floor && this.currentRoute.RouteSections[i].Building === item.Building) {
        this.currentFloor = this.currentRoute.RouteSections[i + 1].Floor;
        this.currentBuilding = this.currentRoute.RouteSections[i + 1].Building;
        await this.showFloor(this.currentFloor, this.currentBuilding);
        return;
      }
    }
  }

  private getInteractiveStairWellMapItemRenderer(floor: string) {
    const pNodes: IPathNode[] = [];

    for (const r of this.routeRenderer) {
      for (let i = 0; i < r.Route.RouteSections.length - 1; i++) {
        if (r.Route.RouteSections[i].Building !== r.Route.RouteSections[i + 1].Building) {
          continue;
        }
        pNodes.push(r.Route.RouteSections[i].Route[r.Route.RouteSections[i].Route.length - 1]);
      }
    }

    const tmpMItems: MapItemRendererCanvas[] = [];
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
    for (const c of this.routeRenderer) {
      c.stopAnimation(this.renderingContext);
    }

    for (const m of this.mapItemRenderer) {
      m.stopAnimation(this.renderingContext);
    }

    for (const l of this.locationRenderer) {
      l.stopAnimation(this.renderingContext);
    }
  }

  private renderMapItems() {
    for (const c of this.mapItemRenderer) {
      c.render(this.renderingContext);
    }
  }

  private renderLocations() {
    for (const l of this.locationRenderer) {
      l.render(this.renderingContext);
    }
  }

  private renderRoutes(args: any) {
    for (const r of this.routeRenderer) {
      r.render(this.renderingContext, args);
    }
  }

  public async onFloorChangeByFloorButton(floorAndBuildingInput: object) {
    // @ts-ignore
    await this.showAnotherFloorOfCurrentBuilding(floorAndBuildingInput.floor, floorAndBuildingInput.building);
  }

  public async showAnotherFloorOfCurrentBuilding(floor: string, building: string) {
    NavigationPage.progressIsVisible = true;
    await this.showFloor(floor, building);
    this.FitMap();
    NavigationPage.progressIsVisible = false;
    this.floorChanged.emit();
  }

  public CenterMap(x: number, y: number) {
    // Modify the X-position to make use of the available space beside the drawer
    const positionToMove = MapViewComponent.calculateMovePosition( x, y);
    this.panZoomController.smoothMoveTo(positionToMove.x, positionToMove.y);
  }

  public MapSize(): DOMRect {
    const element = document.getElementById('map');
    return element.getBoundingClientRect();
  }

  public FitMap() {
    // Using this before centering the map results in a map which is moved to the top not the center
    // const element = document.getElementById('map');
    // this.panZoomController.zoomTo(element.clientWidth, element.clientHeight, 0.7);
    const size = this.MapSize();
    this.CenterMap( size.width / 2, size.height / 2);
  }

  public ClearRoute() {
    this.stopAllAnimations();
    this.currentRoute = null;
  }

  public RefreshMap() {
    this.createNewCanvasMap();
    this.renderMapItems();
    this.renderLocations();
  }
}
