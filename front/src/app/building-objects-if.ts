export interface ICoordinate {
  X: number;
  Y: number;
  Z: number;
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
  render(renderingContext:any);
}

export interface IBuilding {
  Id: number;
  Name: string;
  Floors: string[];
  Campus: string;
}

export interface ILocation {
  Id: number;
  Name: string;
  Description: string;
  Tags: string[];
  Floor: string;
  Building: string;
  PathNode: IPathNode;
}

export interface IRoom extends IMapItem, ILocation{
  Id: number;
}
