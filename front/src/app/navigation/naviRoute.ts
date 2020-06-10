import { Injectable } from '@angular/core';
import {MapItem, PathNode} from '../building-objects-if';
import {CanvasResolutionConfigurator} from '../services/CanvasResolutionConfigurator';
import {IconOnMapRenderer} from '../services/IconOnMapRenderer';

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

export class NaviRoute {

    private readonly mapCanvas: HTMLCanvasElement;
    private readonly map: CanvasRenderingContext2D;

    private pin: IconOnMapRenderer;
    private flag: IconOnMapRenderer;

    public readonly route: ReceivedRoute
    public distance: number;

    constructor(response: ReceivedRoute) {
        this.mapCanvas = document.getElementById('map') as HTMLCanvasElement;
        this.map = CanvasResolutionConfigurator.setup(this.mapCanvas);
        this.pin = new IconOnMapRenderer(this.map, 'pin-sharp.png');
        this.flag = new IconOnMapRenderer(this.map, 'flag-sharp.png');
        this.route = response;
    }

    public render(floor: string) {
        this.renderRoute(floor);
        this.renderDistanceOfRoute(floor);
        this.renderPinAtRouteStart(floor);
        this.renderFlagAtRouteEnd(floor);
    }

    private renderDistanceOfRoute(floor: string) {
        const routeSections = this.route.RouteSections.filter(section => section.Floor === floor);
        if (routeSections != null && routeSections.length > 0) {
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

    private renderRoute(floor: string) {
        const routeSections = this.route.RouteSections.filter(section => section.Floor === floor);
        if (routeSections != null) {
            for (const routeSection of routeSections) {
                this.map.strokeStyle = '#A00';
                this.map.lineWidth = 3;
                this.map.beginPath();
                this.map.moveTo(routeSection.Route[0].Coordinate.X, routeSection.Route[0].Coordinate.Y);
                for (let i = 1; i < routeSection.Route.length; i++) {
                    this.map.lineTo(routeSection.Route[i].Coordinate.X, routeSection.Route[i].Coordinate.Y);
                }
                this.map.stroke();
                this.map.closePath();
            }
        }
    }

    private renderPinAtRouteStart(floor: string) {
        if (this.route.Start.Floor !== floor) {
            return;
        }

        const x = this.route.Start.Node.Coordinate.X;
        const y = this.route.Start.Node.Coordinate.Y;
        this.pin.render(x - 15, y - 30, 30, 30);
    }

    private renderFlagAtRouteEnd(floor: string) {
        if (this.route.End.Floor !== floor) {
            return;
        }
        const x = this.route.End.Node.Coordinate.X;
        const y = this.route.End.Node.Coordinate.Y;
        this.flag.render(x - 5, y - 30, 30, 30);
    }
}
