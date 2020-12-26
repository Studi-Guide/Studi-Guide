import {NavigationModel} from '../navigation/navigationModel';
import {Storage} from '@ionic/storage';

export class SearchResultProvider{
    private static recentSearchesKey = 'searches';
    public static addRecentSearch(location:string, model: NavigationModel, storage: Storage) {
        if (model.recentSearches.includes(location)) {
            model.recentSearches.splice(model.recentSearches.indexOf(location), 1);
        }

        model.recentSearches.unshift(location);

        if (model.recentSearches.length > 3) {
            model.recentSearches.pop();
        }

        storage.set(this.recentSearchesKey, JSON.stringify([model.recentSearches]));
        console.log(model.recentSearches);
    }

    public static async readRecentSearches(storage: Storage) {
        const searches = JSON.parse(await storage.get(this.recentSearchesKey));
        return searches != null ? searches[0] : null;
    }
}
