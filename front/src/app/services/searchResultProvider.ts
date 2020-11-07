import {Injectable} from '@angular/core';
import {NavigationModel} from '../navigation/navigationModel';

@Injectable({
    providedIn: 'root'
})
export class SearchResultProvider{
    private recentSearchesKey = 'searches';

    constructor(private storage: Storage) {
    }


    public addRecentSearch(location:string, model: NavigationModel) {
        if (model.recentSearches.includes(location)) {
            model.recentSearches.splice(model.recentSearches.indexOf(location), 1);
        }

        model.recentSearches.unshift(location);

        if (model.recentSearches.length > 3) {
            model.recentSearches.pop();
        }

        this.storage.set(this.recentSearchesKey, JSON.stringify([model.recentSearches]));
        console.log(model.recentSearches);
    }

    public async readRecentSearch() {
        const searches = JSON.parse(await this.storage.get(this.recentSearchesKey));
        return searches != null ? searches[0] : null;
    }
}
