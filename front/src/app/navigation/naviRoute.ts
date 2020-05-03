import { Injectable } from '@angular/core';
import {PathNode} from '../building-objects-if';

@Injectable({
    providedIn: 'root'
})

export class ReceivedRoute {
    Distance: number;
    Route: PathNode[];
}

export class DistanceToBeDisplayed {
    Value: number;
    X: number;
    Y: number;
}

export class NaviRoute {

    private pathNodesToGo:PathNode[] = [{'Coordinate': { 'X': 0, 'Y': 0, 'Z': 0 }}];
    public distance: DistanceToBeDisplayed = { 'Value':0, 'X': 0, 'Y': 0 };
    public svgRoute: string;

    constructor(response:ReceivedRoute) {
        this.distance.Value = response.Distance;
        this.pathNodesToGo = response.Route;
        this.calculateSvgPositionForDistance();
        this.calculateSvgPathForRoute();
    }

    private calculateSvgPositionForDistance() {
        const numberOfPathNodes:number = this.pathNodesToGo.length;
        this.distance.X = this.pathNodesToGo[Math.round((numberOfPathNodes-1)/2)].Coordinate.X;
        this.distance.Y = this.pathNodesToGo[Math.round((numberOfPathNodes-1)/2)].Coordinate.Y;
    }

    private calculateSvgPathForRoute() {
        let points = '';
        for (const pathNode of this.pathNodesToGo) {
            points += pathNode.Coordinate.X + ',' + pathNode.Coordinate.Y + ' ';
        }
        this.svgRoute = points;
    }

    public getRouteStart() {
        return this.pathNodesToGo[0];
    }

    public getRouteEnd() {
        return this.pathNodesToGo[this.pathNodesToGo.length-1];
    }
}