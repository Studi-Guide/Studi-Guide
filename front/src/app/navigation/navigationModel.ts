import {IBuilding, ICampus, ILocation} from '../building-objects-if';
import {Injectable} from '@angular/core';

@Injectable({
    providedIn: 'root'
})
export class NavigationModel {
    public recentSearches : string[] = [];
    public errorMessage: string;
    public selectedLocation:ILocation = {
        Building: '',
        Description: '',
        Floor: '',
        Id: 0,
        Name: '',
        PathNode: {
            Coordinate: {X: 0, Y: 0, Z: 0},
            Id: 0
        },
        Tags: []
    };
    public selectedBuilding: IBuilding = {
        Body: [],
        Campus: '',
        Color: '',
        Floors: [],
        Id: 0,
        Name: ''
    };
    public availableCampus: ICampus[] = [];
}
