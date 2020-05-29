import { Injectable } from '@angular/core';
import {PathNode} from '../building-objects-if';

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

    private mapCanvas: HTMLCanvasElement;
    private map: CanvasRenderingContext2D;

    private readonly routeSections:RouteSection[];
    public distance: number;

    constructor(response:ReceivedRoute) {
        this.mapCanvas = document.getElementById('map') as HTMLCanvasElement;
        this.map = this.mapCanvas.getContext('2d');
        this.distance = response.Distance;
        this.routeSections = response.RouteSections;
        // this.calculateSvgPositionForDistance();
        // this.calculateSvgPathForRoute();
    }

    public render(building: string, floor :string) {
        this.renderRoute(building, floor);
        this.renderDistanceOfRoute(building, floor);
        this.renderPinAtRouteStart(building, floor);
    }

    private renderDistanceOfRoute(building: string, floor :string) {
        const routeSection = this.routeSections.find(section => section.Building === building && section.Floor === floor);
        if (routeSection != null){
            const numberOfPathNodes:number = routeSection.Route.length;
            const value:number = routeSection.Distance;
            const x:number = routeSection.Route[Math.round((numberOfPathNodes-1)/2)].Coordinate.X;
            const y:number = routeSection.Route[Math.round((numberOfPathNodes-1)/2)].Coordinate.Y;
            const font = '14px Arial';
            const width = this.map.measureText(value.toString()).width+14;
            const height = parseInt('14px Arial', 10)+14;
            this.map.fillStyle = '#A00';
            this.map.fillRect(x-width/2, y-height/2, width, 20);
            this.map.font = font;
            this.map.fillStyle = '#FFF';
            this.map.fillText(value.toString(), x, y);
        }
    }

    private renderRoute(building: string, floor:string) {
        const routeSection = this.routeSections.find(section => section.Building === building && section.Floor === floor);
        if (routeSection != null) {
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

    private renderPinAtRouteStart(building: string, floor :string) {
        const routeStart = this.getRouteStart(building, floor);
        const x = routeStart.Coordinate.X;
        const y = routeStart.Coordinate.Y;
        const image = new Image();
        image.onload = () => {
            this.map.drawImage(image, x, y, 30, 30);
        };
        image.src = 'assets/pin-outline.png';

    }

    private getRouteStart(building:string, floor:string) {
        const routeSection = this.routeSections.find(section => section.Building === building && section.Floor === floor);
        if (routeSection === this.routeSections[0]) {
            return routeSection.Route[0];
        }
        return null;
    }

    private getRouteEnd(building:string, floor:string) {
        const routeSection = this.routeSections.find(section => section.Building === building && section.Floor === floor);
        const lastRouteSection = this.routeSections[this.routeSections.length-1];
        if (routeSection === lastRouteSection) {
            return lastRouteSection.Route[lastRouteSection.Route.length-1];
        }
        return null;
    }
}