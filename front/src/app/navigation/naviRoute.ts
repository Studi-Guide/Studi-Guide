import { Injectable } from '@angular/core';
import {PathNode} from "../building-objects-if";

export class NaviRoute {

    private pathNodesToGo:PathNode[];
    public svgRoute: string;

    constructor(pathNodesToConnectAndDisplay: PathNode[]) {
        this.pathNodesToGo = pathNodesToConnectAndDisplay;
    }

    public calculateSvgPathForRoute() {
        let points:string = '';
        for (const pathNode of this.pathNodesToGo) {
            points += pathNode.Coordinate.X + ',' + pathNode.Coordinate.Y + ' ';
        }
        this.svgRoute = points;
    }
}