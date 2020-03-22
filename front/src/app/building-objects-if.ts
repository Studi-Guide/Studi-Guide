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

export class Door implements Section {
  Start: Coordinate;
  End: Coordinate;
  pathNode: Coordinate;
}

export interface Section {
  Start: Coordinate;
  End: Coordinate;
}

export class Room {
  name: string;
  sections: Section[];
  alias: string[];
  pathNodes: Coordinate[];
  doors: Door[];
  Color: string;
}

export class svgPath {
  d: string;
  fill: string;
}