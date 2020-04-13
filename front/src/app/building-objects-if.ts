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

export interface Room {
  Name: string;
  Floor: number;
  Sections: Section[];
  PathNodes: PathNode[];
  Doors: Door[];
  Color: string;
}

export class svgPath {
  d: string;
  fill: string;
}

export class RoomName {
  name: string;
  x: number;
  y: number;
}