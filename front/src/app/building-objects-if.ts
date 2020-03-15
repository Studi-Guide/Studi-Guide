export class floor {
  rooms:room[];
  corridors:corridor[];
}

export class corridor {
  name: string;
  fill: string;
  width: number;
  height: number;
  x: number;
  y: number;
}

export class coordinate {
  x: number;
  y: number;
  z: number;
}

export class door implements section {
  start: coordinate;
  end: coordinate;
}

export interface section {
  start: coordinate;
  end: coordinate;
}

export class room {
  name: string;
  sections: section[];
  alias: string;
  doors: section[];
  fill: string;
}

export class svgPath {
  d: string;
  fill: string;
}