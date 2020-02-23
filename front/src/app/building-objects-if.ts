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
  
export interface room {
  name: string;
  section: section[];
  alias;
  doors: door[];
  fill: string;
}

export interface door {
  coordinates: coordinate[];
}

export interface section {
  start: coordinate;
  end: coordinate;
}

export interface coordinate {
  x: number;
  y: number;
  // z: int;
}