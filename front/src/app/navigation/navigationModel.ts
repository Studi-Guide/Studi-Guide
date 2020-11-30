import {DrawerObject, ICampus, ILocation} from '../building-objects-if';
import {Injectable} from '@angular/core';

@Injectable({
    providedIn: 'root'
})
export class NavigationModel {
    public recentSearches : string[] = [];
    public errorMessage: string;
    public selectedObject:DrawerObject = new DrawerObject();
    public availableCampus: ICampus[] = [];
}
