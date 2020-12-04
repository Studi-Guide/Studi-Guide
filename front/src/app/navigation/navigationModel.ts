import {DrawerObject, IBuilding, ICampus, ILocation} from '../building-objects-if';
import {Injectable} from '@angular/core';
import {Params} from "@angular/router";
import {LatLngLiteral} from "leaflet";



export interface ISearchResultObject {
    Name: string;
    Description: string;
    Information: [string, any][];
    DetailRouterParams: Params;
    RouteRouterParams: Params;
    LatLng: LatLngLiteral;
}

@Injectable({
    providedIn: 'root'
})
export class NavigationModel {
    public recentSearches : string[] = [];
    public errorMessage: string;
    public latestSearchResult: ISearchResultObject = {
        Name: '',
        Description: '',
        Information: [],
        DetailRouterParams: {},
        RouteRouterParams: {},
        LatLng: {lat: 0, lng: 0}
    };
    public availableCampus: ICampus[] = [];

    public SetCampusAsSearchResultObject(c:ICampus) {
        this.latestSearchResult = {
            Name: c.Name,
            Description: c.ShortName,
            Information:
                [
                    ['ShortName: ', c.ShortName],
                    ['Longitude: ', c.Longitude],
                    ['Latitude: ', c.Latitude],
                ],
            DetailRouterParams: {},
            RouteRouterParams: {},
            LatLng: {lat: c.Latitude, lng: c.Longitude}
        };
    }

    public SetBuildingAsSearchResultObject(b:IBuilding, latLng: LatLngLiteral) {
        this.latestSearchResult = {
            Name: b.Name,
            Description: 'Campus:' + b.Campus,
            Information: [],
            DetailRouterParams: {building: b.Name},
            RouteRouterParams: {},
            LatLng: latLng
        }
    }
    public SetLocationAsSearchResultObject(l:ILocation, latLng: LatLngLiteral) {
        const details: [string, any][] = [['Building: ',  l.Name]];
        if (l.Tags) {
            details.push(['Tags: ', l.Tags.join(', ')]);
        }
        this.latestSearchResult = {
            Name: l.Name,
            Description: l.Description,
            Information: details,
            DetailRouterParams: {location: l.Name},
            RouteRouterParams: {start: l.Building+'.Entrance', destination: l.Name},
            LatLng: latLng
        };
    }
}
