import {ICampus, ILocation} from '../building-objects-if';

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

    public availableCampus: ICampus[];
}
