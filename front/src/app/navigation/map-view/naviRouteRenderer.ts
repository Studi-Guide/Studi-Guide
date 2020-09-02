import { Injectable } from '@angular/core';
import {IMapItem, IPathNode} from '../../building-objects-if';
import {CanvasResolutionConfigurator} from '../../services/CanvasResolutionConfigurator';
import {IconOnMapRenderer} from '../../services/IconOnMapRenderer';
import {DataService} from '../../services/data.service';
import {resolveFileWithPostfixes} from '@angular/compiler-cli/ngcc/src/utils';

@Injectable({
    providedIn: 'root'
})

export class ReceivedRoute {
    Distance:       number;
    Start:          RoutePoint;
    End:            RoutePoint;
    RouteSections:  RouteSection[];
}

export class RoutePoint {
    Node:           IPathNode;
    Floor:          string;
}

export class RouteSection {
    Route: 		IPathNode[];
    Description: string;
    Distance: 	 number;
    Building: 	string;
    Floor: 		string;
}

export class NaviRouteRenderer {

    private pin: IconOnMapRenderer;
    private flag: IconOnMapRenderer;
    private doAnim = false;
    private animationCallback: () => void;


    constructor(private dataService:DataService) {
    }

    public async render(map: CanvasRenderingContext2D, route:ReceivedRoute, floor :string) {
        this.renderRoute(map, route, floor);
        this.renderDistanceOfRoute(map, route, floor);
        this.renderPinAtRouteStart(map, route, floor);
        this.renderFlagAtRouteEnd(map, route, floor);
        await this.renderFlashingStairWell(map, route, floor);
    }

    public startAnimation() {
        this.doAnim = true;
        if (this.animationCallback != null) {
            window.requestAnimationFrame(this.animationCallback);
        }
    }

    public stopAnimation() {
        this.doAnim = false;
        this.animationCallback = null;
    }

    private renderDistanceOfRoute(map: CanvasRenderingContext2D, route:ReceivedRoute, floor :string) {
        const routeSections = route.RouteSections.filter(section => section.Floor === floor);
        if (routeSections != null && routeSections.length > 0){
            for (const routeSection of routeSections) {
                const numberOfPathNodes: number = routeSection.Route.length;
                const value: number = routeSection.Distance;
                const x: number = routeSection.Route[Math.round((numberOfPathNodes - 1) / 2)].Coordinate.X;
                const y: number = routeSection.Route[Math.round((numberOfPathNodes - 1) / 2)].Coordinate.Y;
                const font = '14px Arial';
                const width = map.measureText(value.toString()).width + 14;
                const height = parseInt('14px Arial', 10) + 14;
                map.fillStyle = '#A00';
                map.fillRect(x - width / 2, y - height / 2, width, 20);
                map.font = font;
                map.fillStyle = '#FFF';
                map.fillText(value.toString(), x, y);
            }
        }
    }

    private renderRoute(map: CanvasRenderingContext2D, route:ReceivedRoute, floor :string) {
        const routeSections = route.RouteSections.filter(section => section.Floor === floor);
        if (routeSections != null) {
            for (const routeSection of routeSections) {
                map.strokeStyle = '#A00';
                map.lineWidth = 3;
                map.beginPath();
                map.moveTo(routeSection.Route[0].Coordinate.X,routeSection.Route[0].Coordinate.Y);
                for (let i = 1; i < routeSection.Route.length; i++) {
                    map.lineTo(routeSection.Route[i].Coordinate.X,routeSection.Route[i].Coordinate.Y);
                }
                map.stroke();
                map.closePath();
            }
        }
    }

    private renderPinAtRouteStart(map: CanvasRenderingContext2D, route:ReceivedRoute, floor :string) {
        this.pin = new IconOnMapRenderer('pin-sharp.png');
        if (route.Start.Floor !== floor) {
            return;
        }

        const x = route.Start.Node.Coordinate.X;
        const y = route.Start.Node.Coordinate.Y;
        this.pin.render(map,x-15, y-30, 30, 30);
    }

    public async getInteractiveStairWellMapItems(route:ReceivedRoute, floor:string) {
        const pNodes:IPathNode[] = [];
        for (let i = 0; i < route.RouteSections.length-1; i++) {
            if (route.RouteSections[i].Building !== route.RouteSections[i+1].Building)
                continue;
            pNodes.push(route.RouteSections[i].Route[route.RouteSections[i].Route.length-1]);
        }

        const tmpMItemss:IMapItem[][] = [];
        for (const pNode of pNodes) {
            tmpMItemss.push(await this.dataService.get_map_item(pNode.Id).toPromise<IMapItem[]>());
        }

        const mItems:IMapItem[] = [];
        for (const mapItems of tmpMItemss) {
            for (const mapItem of mapItems) {
                if (mapItem.Floor !== floor) {
                    continue;
                }
                mItems.push(mapItem);
            }
        }

        return mItems;
    }

    private async renderFlashingStairWell(map: CanvasRenderingContext2D, route:ReceivedRoute, floor :string) {

        const mItems:IMapItem[] = await this.getInteractiveStairWellMapItems(route, floor);

        if (mItems.length === 0) {
            this.stopAnimation();
            return;
        }

        this.doAnim = true;

        this.animationCallback = () => {
            if (this.doAnim !== true) {
                return;
            }
            map.save();
            const date = new Date();
            for (const mapItem of mItems) {
                map.beginPath();
                map.moveTo(mapItem.Sections[0].Start.X, mapItem.Sections[0].Start.Y);
                for (let i = 1; i < mapItem.Sections.length; i++) {
                    map.lineTo(mapItem.Sections[i].Start.X, mapItem.Sections[i].Start.Y);
                }
                map.lineTo(mapItem.Sections[0].Start.X, mapItem.Sections[0].Start.Y);
                map.lineWidth = 1;
                map.strokeStyle = (date.getSeconds() % 2 === 0) ? '#FFF' : '#000';
                map.stroke();
                map.closePath();
            }
            map.restore();

            window.requestAnimationFrame(this.animationCallback);

        };

    }

    private renderFlagAtRouteEnd(map: CanvasRenderingContext2D, route:ReceivedRoute, floor :string) {
        this.flag = new IconOnMapRenderer('flag-sharp.png');
        if (route.End.Floor !== floor) {
            return;
        }
        const x = route.End.Node.Coordinate.X;
        const y = route.End.Node.Coordinate.Y;
        this.flag.render(map, x-5, y-30, 30, 30);
    }
}
