export interface floor {
  rooms:room[];
  corridors:corridor[];
}

export interface corridor {
  name: string;
  fill: string;
  width: number;
  height: number;
  x: number;
  y: number;
}

export interface coordinate {
  x: number;
  y: number;
  z: number;
}

export interface door {
  start: coordinate;
  end: coordinate;
}

export interface section {
  start: coordinate;
  end: coordinate;
}

export interface room {
  name: string;
  sections: section[];
  alias: string;
  doors: section[];
  fill: string;
}

export interface svgPath {
  d: string;
  fill: string;
}