export class Floor {
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
  Id: number;
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
  Floor: string;
  Sections: Section[];
  Campus: string;
  Building: string;
  PathNodes: PathNode[];
}

export interface BuildingData {
  Id: number;
  Name: string;
  Floors: string[];
}

export interface Location {
  Id: number;
  Name: string;
  Description: string;
  Tags: string[];
  PathNode: PathNode;
  Floor: string;
  Building: string;
}

export interface Room extends MapItem, Location{

}


export class SvgPath {
  d: string;
  fill: string;
}

export class SvgLocationName {
  name: string;
  x: number;
  y: number;
}
