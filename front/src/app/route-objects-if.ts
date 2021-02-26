import {IPathNode} from './building-objects-if';

export interface IReceivedRoute {
    Distance: number;
    Start: IRoutePoint;
    End: IRoutePoint;
    RouteSections: IRouteSection[];
}

export interface IRoutePoint {
    Node: IPathNode;
    Name: string;
    Floor: string;
}

export interface IRouteSection {
    Route: IPathNode[];
    Description: string;
    Distance: number;
    Building: string;
    Floor: string;
}
