import {IBuilding, ICampus, ILocation} from '../building-objects-if';
import {Injectable} from '@angular/core';
import {LatLngLiteral} from 'leaflet';
import {INavigationInstruction} from './navigation-instruction-slides/navigation-instruction-if';
import {OsmRoute} from '../services/osm/open-street-map.service';
import {CampusViewModel} from './campusViewModel';
import {ISearchResultObject, RecentSearchesService} from '../services/recent-searches/recent-searches.service';

export interface IRouteLocation {
    Name: string;
    LatLng: LatLngLiteral;
}

export interface INavigationRoute {
    Start: IRouteLocation;
    Destination: IRouteLocation;
    Coordinates: [number, number][];
    Distance: number;
    NavigationInstructions: INavigationInstruction[];
    Time: number;
}

@Injectable({
    providedIn: 'root'
})
export class NavigationModel {

    constructor(private recentSearchesService: RecentSearchesService) {
        this.recentSearchesService.readRecentSearches().then(r => {
            this.recentSearchesVar = r;
        });
    }

    private recentSearchesVar: ISearchResultObject[] = [];
    public errorMessage: string;
    public get latestSearchResult(): ISearchResultObject {
        return this.recentSearches[0] !== undefined ? this.recentSearches[0] : {
            Name: '',
            Description: '',
            Information: [],
            DetailRouterParams: {},
            RouteRouterParams: {},
            LatLng: {lat: 0, lng: 0},
            Images: []
        };
    }
    public availableCampus: CampusViewModel[] = [];
    public Route: INavigationRoute = {
        Start: {
            Name: '',
            LatLng: {lat: 0, lng: 0}
        },
        Destination: {
            Name: '',
            LatLng: {lat: 0, lng: 0}
        },
        Coordinates: [],
        Distance: 0,
        NavigationInstructions: [],
        Time: 0
    };

    public get recentSearches(): ISearchResultObject[] {
        return this.recentSearchesVar;
    }

    public async addRecentSearch(location: ISearchResultObject) {
        await this.recentSearchesService.addRecentSearch(location);
        this.recentSearchesVar = await this.recentSearchesService.readRecentSearches();
    }

    public async addRecentSearchCampus(c: ICampus) {
        const details: [string, any][] = [
            ['ShortName: ', c.ShortName],
            ['Longitude: ', c.Longitude],
            ['Latitude: ', c.Latitude],
        ];
        const searchResult = {
            Name: c.Name,
            Description: c.ShortName,
            Information: details,
            DetailRouterParams: {},
            RouteRouterParams: {},
            LatLng: {lat: c.Latitude, lng: c.Longitude},
            Images: []
        };
        await this.addRecentSearch(searchResult);
    }
    public async addRecentSearchBuilding(b: IBuilding, latLng: LatLngLiteral) {
        const address = b.edges.Address.Street + ' ' + b.edges.Address.Number + ', ' + b.edges.Address.PLZ;
        const details: [string, any][] = [['Campus: ',  b.edges.Campus.Name]];
        const searchResult = {
            Name: b.Name,
            Description: address,
            Information: details,
            DetailRouterParams: {building: b.Name},
            RouteRouterParams: {},
            LatLng: latLng,
            Images: []
        };
        await this.addRecentSearch(searchResult);
    }
    public async addRecentSearchLocation(l: ILocation, latLng: LatLngLiteral, building: IBuilding) {
        const details: [string, any][] = [['Building: ',  l.Name]];
        if (building) {
            details.push(['Address: ', building.edges.Address.Street + ' ' + building.edges.Address.Number]);
        }

        if (l.Tags) {
            details.push(['Tags: ', l.Tags.join(', ')]);
        }

        const searchResult = {
            Name: l.Name,
            Description: l.Description,
            Information: details,
            DetailRouterParams: {location: l.Name},
            RouteRouterParams: {start: l.Building + '.Entrance', destination: l.Name},
            LatLng: latLng,
            Images: l.Images
        };
        await this.addRecentSearch(searchResult);
    }

    public SetOsmRouteAsRoute(route: OsmRoute, start: IRouteLocation, destination: IRouteLocation) {
        const leafletLatLng: [number, number][] = [];
        for (const coordinate of route.Points.Coordinates) {
            leafletLatLng.push([coordinate.Lat, coordinate.Lng]);
        }

        this.Route = {
            Start: start,
            Destination: destination,
            Coordinates: leafletLatLng,
            Distance: Math.round(route.Distance),
            NavigationInstructions: route.Instructions,
            Time: Math.round(route.Time / 1000 / 60)
        };
    }

    public ClearRoute() {
        this.Route = {
            Start: {
                Name: '',
                LatLng: {lat: 0, lng: 0}
            },
            Destination: {
                Name: '',
                LatLng: {lat: 0, lng: 0}
            },
            Distance: 0,
            Coordinates: [],
            NavigationInstructions: [],
            Time: 0
        };
    }
}
