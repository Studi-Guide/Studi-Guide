import { Injectable } from '@angular/core';
import {PathNode} from '../building-objects-if';
import {CanvasResolutionConfigurator} from '../services/CanvasResolutionConfigurator';
import {IconOnMapRenderer} from '../services/IconOnMapRenderer';

@Injectable({
    providedIn: 'root'
})

export class ReceivedRoute {
    Distance: number;
    RouteSections: RouteSection[];
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

    public readonly routeSections:RouteSection[];
    public distance: number;

    constructor(response:ReceivedRoute) {
        this.mapCanvas = document.getElementById('map') as HTMLCanvasElement;
        this.map = CanvasResolutionConfigurator.setup(this.mapCanvas);
        this.pin = new IconOnMapRenderer(this.map,'pin-sharp.png');
        this.flag = new IconOnMapRenderer(this.map,'flag-sharp.png');
        this.distance = response.Distance;
        this.routeSections = response.RouteSections;
    }

    public render(floor :string) {
        this.renderRoute(floor);
        this.renderDistanceOfRoute(floor);
        this.renderPinAtRouteStart(floor);
        this.renderFlagAtRouteEnd(floor);
        this.renderFlashingStairWell(floor);
    }

    private renderDistanceOfRoute(floor :string) {
        const routeSections = this.routeSections.filter(section => section.Floor === floor);
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

    private renderRoute(floor:string) {
        const routeSections = this.routeSections.filter(section => section.Floor === floor);
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

    private renderPinAtRouteStart(floor :string) {
        const routeStart = this.getRouteStart(floor);
        if (routeStart == null) {
            return;
        }
        const x = routeStart.Coordinate.X;
        const y = routeStart.Coordinate.Y;
        this.pin.render(x-15, y-30, 30, 30);
    }

    private renderFlashingStairWell(building, floor): void {
        ;
    }

    private getRouteStart(floor:string) {
        const routeSection = this.routeSections.find(section =>section.Floor === floor);
        if (routeSection === this.routeSections[0]) {
            return routeSection.Route[0];
        }
        return null;
    }

    private renderFlagAtRouteEnd(floor :string) {
        const routeEnd = this.getRouteEnd(floor);
        if (routeEnd == null) {
            return;
        }
        const x = routeEnd.Coordinate.X;
        const y = routeEnd.Coordinate.Y;
        this.flag.render(x-5, y-30, 30, 30);
    }

    private getRouteEnd(floor:string) {
        const filtered = this.routeSections.filter(section => section.Floor === floor);
        if (filtered != null && filtered.length > 0) {
            const lastRouteSection = filtered[filtered.length-1];
             if (lastRouteSection === this.routeSections[this.routeSections.length - 1]) {
                return lastRouteSection.Route[lastRouteSection.Route.length - 1];
            }
        }
        return null;
    }
}
