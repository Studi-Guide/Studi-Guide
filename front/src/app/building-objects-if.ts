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

export interface Coordinate {
  X: number;
  Y: number;
  Z: number;
}

export class PathNode {
  X: number;
  Y: number;
  Z: number;
}

export class Door implements Section {
  Start: Coordinate;
  End: Coordinate;
  pathNode: PathNode;
}

export interface Section {
  Start: Coordinate;
  End: Coordinate;
}

export class Room {
  name: string;
  sections: Section[];
  alias: string[];
  pathNodes: PathNode[];
  doors: Door[];
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