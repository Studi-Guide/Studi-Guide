import {Path} from "@angular-devkit/core";

export class floor {
  rooms: Room[];
  corridors: Corridor[];
}

export class Corridor {
  name: string;
  fill: string;
  width: number;
  height: number;
  X: number;
  Y: number;
}

export class Coordinate {
  X: number;
  Y: number;
  Z: number;
}

export interface PathNode {
  Coordinate: Coordinate;
}

export class Door {
  Id: number;
  Section: Section;
  pathNode: PathNode;
}

export interface Section {
  Start: Coordinate;
  End: Coordinate;
}

export interface MapItem {
  Doors: Door[];
  Color: string;
  Floor: number;
  Sections: Section[];
  Campus: string;
  Building: string;
  PathNodes: PathNode[];
}

export interface Location {
  Id: number;
  Name: string;
  Description: string;
  Tags: string[];
  PathNode: PathNode;
  Floor: number;
}

export interface Room extends MapItem, Location{

}

export class svgPath {
  d: string;
  fill: string;
}

export class SvgLocationName {
  name: string;
  x: number;
  y: number;
}