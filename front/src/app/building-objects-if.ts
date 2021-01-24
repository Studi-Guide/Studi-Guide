export interface ICoordinate {
  X: number;
  Y: number;
  Z: number;
}

export interface IFile {
  Name: string;
  Path: string;
}

export interface IPathNode {
  Id: number;
  Coordinate: ICoordinate;
}

export interface ISection {
  Id: number;
  Start: ICoordinate;
  End: ICoordinate;
}

export interface IDoor {
  Id: number;
  Section: ISection;
  pathNode: IPathNode;
}

export interface IMapItem {
  Doors: IDoor[];
  Color: string;
  Sections: ISection[];
  Campus: string;
  Building: string;
  PathNodes: IPathNode[];
  Floor: string;
}

export interface IGpsCoordinate {
  Longitude: number
  Latitude: number
}

export interface IBuilding {
  Id: number;
  Name: string;
  Color: string;
  Floors: string[];
  Campus: string;
  Body: IGpsCoordinate[];
}

export interface ILocation {
  Id: number;
  Name: string;
  Description: string;
  Tags: string[];
  Floor: string;
  Building: string;
  PathNode: IPathNode;
  Images: IFile[];
  Icon: string;
}

export interface IAddress {
  City: string;
  Country: string;
  Number: string;
  PLZ: number;
  Street: string;
  id: number;
}

export interface ICampusEdges {
  Address: IAddress[];
}

export interface ICampus {
  Latitude: number;
  Longitude: number;
  Name: string;
  ShortName: string;
  edges: ICampusEdges;
  id: number;
}

export interface IRoom extends IMapItem, ILocation{
  Id: number;
}

export interface IRenderer {
  render(renderingContext:any, args?:any)
  startAnimation(renderingContext:any, args?:any)
  stopAnimation(renderingContext:any, args?:any)
}
