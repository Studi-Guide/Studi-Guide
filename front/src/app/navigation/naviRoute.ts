import { Injectable } from '@angular/core';
import {PathNode} from '../building-objects-if';
import {Route} from "@angular/router";

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

export class DistanceToBeDisplayed {
    Value: number;
    X: number;
    Y: number;
}

export class NaviRoute {

    private routeSections:RouteSection[];
    public distance: number;

    constructor(response:ReceivedRoute) {
        this.distance = response.Distance;
        this.routeSections = response.RouteSections;
        // this.calculateSvgPositionForDistance();
        // this.calculateSvgPathForRoute();
    }

    public calculateSvgPositionForDistance(building: string,  floor :string) {
        const routeSection = this.routeSections.find(section => section.Building === building && section.Floor === floor);
        const rtnDistance = new DistanceToBeDisplayed();
        if (routeSection != null){
            const numberOfPathNodes:number = routeSection.Route.length;
            rtnDistance.Value = this.distance;
            rtnDistance.X = routeSection.Route[Math.round((numberOfPathNodes-1)/2)].Coordinate.X;
            rtnDistance.Y = routeSection.Route[Math.round((numberOfPathNodes-1)/2)].Coordinate.Y;
        }

        return rtnDistance;
    }

    public calculateSvgPathForRoute(building: string,  floor:string) {
        const routeSection = this.routeSections.find(section => section.Building === building && section.Floor === floor);
        let points = '';
        if (routeSection != null) {
            for (const pathNode of routeSection.Route) {
                points += pathNode.Coordinate.X + ',' + pathNode.Coordinate.Y + ' ';
            }
        }
        return points;
    }

    public getRouteStart() {
        return this.routeSections[0].Route[0];
    }

    public getRouteEnd() {
        const lastroutesection = this.routeSections[this.routeSections.length-1];
        return lastroutesection.Route[lastroutesection.Route.length-1];
    }
}