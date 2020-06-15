import { Injectable } from '@angular/core';
import {MapItem, PathNode} from '../../building-objects-if';
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
    Node:           PathNode;
    Floor:          string;
}

export class RouteSection {
    Route: 		PathNode[];
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


    constructor(private dataService:DataService, private map:CanvasRenderingContext2D) {
    }

    public async render(route:ReceivedRoute, floor :string) {
        this.renderRoute(route, floor);
        this.renderDistanceOfRoute(route, floor);
        this.renderPinAtRouteStart(route, floor);
        this.renderFlagAtRouteEnd(route, floor);
        await this.renderFlashingStairWell(route, floor);
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

    private renderDistanceOfRoute(route:ReceivedRoute, floor :string) {
        const routeSections = route.RouteSections.filter(section => section.Floor === floor);
        if (routeSections != null && routeSections.length > 0){
            for (const routeSection of routeSections) {
                const numberOfPathNodes: number = routeSection.Route.length;
                const value: number = routeSection.Distance;
                const x: number = routeSection.Route[Math.round((numberOfPathNodes - 1) / 2)].Coordinate.X;
                const y: number = routeSection.Route[Math.round((numberOfPathNodes - 1) / 2)].Coordinate.Y;
                const font = '14px Arial';
                const width = this.map.measureText(value.toString()).width + 14;
                const height = parseInt('14px Arial', 10) + 14;
                this.map.fillStyle = '#A00';
                this.map.fillRect(x - width / 2, y - height / 2, width, 20);
                this.map.font = font;
                this.map.fillStyle = '#FFF';
                this.map.fillText(value.toString(), x, y);
            }
        }
    }

    private renderRoute(route:ReceivedRoute, floor :string) {
        const routeSections = route.RouteSections.filter(section => section.Floor === floor);
        if (routeSections != null) {
            for (const routeSection of routeSections) {
                this.map.strokeStyle = '#A00';
                this.map.lineWidth = 3;
                this.map.beginPath();
                this.map.moveTo(routeSection.Route[0].Coordinate.X,routeSection.Route[0].Coordinate.Y);
                for (let i = 1; i < routeSection.Route.length; i++) {
                    this.map.lineTo(routeSection.Route[i].Coordinate.X,routeSection.Route[i].Coordinate.Y);
                }
                this.map.stroke();
                this.map.closePath();
            }
        }
    }

    private renderPinAtRouteStart(route:ReceivedRoute, floor :string) {
        this.pin = new IconOnMapRenderer(this.map,'pin-sharp.png');
        if (route.Start.Floor !== floor) {
            return;
        }

        const x = route.Start.Node.Coordinate.X;
        const y = route.Start.Node.Coordinate.Y;
        this.pin.render(x-15, y-30, 30, 30);
    }

    private async renderFlashingStairWell(route:ReceivedRoute, floor :string) {
        const pNodes:PathNode[] = [];
        for (let i = 0; i < route.RouteSections.length-1; i++) {
            if (route.RouteSections[i].Building !== route.RouteSections[i+1].Building)
                continue;
            pNodes.push(route.RouteSections[i].Route[route.RouteSections[i].Route.length-1]);
        }

        const tmpMItemss:MapItem[][] = [];
        for (const pNode of pNodes) {
            tmpMItemss.push(await this.dataService.get_map_item(pNode.Id).toPromise<MapItem[]>());
        }

        const mItems:MapItem[] = [];
        for (const mapItems of tmpMItemss) {
            for (const mapItem of mapItems) {
                if (mapItem.Floor !== floor) {
                    continue;
                }
                mItems.push(mapItem);
            }
        }

        if (mItems.length === 0) {
            this.stopAnimation();
            return;
        }

        this.doAnim = true;

        this.animationCallback = () => {
            if (this.doAnim !== true) {
                return;
            }
            this.map.save();
            const date = new Date();
            for (const mapItem of mItems) {
                this.map.beginPath();
                this.map.moveTo(mapItem.Sections[0].Start.X, mapItem.Sections[0].Start.Y);
                for (let i = 1; i < mapItem.Sections.length; i++) {
                    this.map.lineTo(mapItem.Sections[i].Start.X, mapItem.Sections[i].Start.Y);
                }
                this.map.lineTo(mapItem.Sections[0].Start.X, mapItem.Sections[0].Start.Y);
                this.map.lineWidth = 1;
                this.map.strokeStyle = (date.getSeconds() % 2 === 0) ? '#FFF' : '#000';
                this.map.stroke();
                this.map.closePath();
            }
            this.map.restore();

            window.requestAnimationFrame(this.animationCallback);

        };

    }

    private renderFlagAtRouteEnd(route:ReceivedRoute, floor :string) {
        this.flag = new IconOnMapRenderer(this.map,'flag-sharp.png');
        if (route.End.Floor !== floor) {
            return;
        }
        const x = route.End.Node.Coordinate.X;
        const y = route.End.Node.Coordinate.Y;
        this.flag.render(x-5, y-30, 30, 30);
    }
}
