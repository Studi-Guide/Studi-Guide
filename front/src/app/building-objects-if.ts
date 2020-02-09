export interface floor {
  rooms:room[];
  corridors:corridor[];
};
  
export interface room {
  name: string;
  fill: string;
  width: number;
  height: number;
  x: number;
  y: number;
};
  
export interface corridor {
  name: string;
  fill: string;
  width: number;
  height: number;
  x: number;
  y: number;
};