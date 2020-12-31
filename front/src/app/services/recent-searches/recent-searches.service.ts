import {Injectable} from '@angular/core';
import {Storage} from '@ionic/storage';
import {Params} from '@angular/router';
import {LatLngLiteral} from 'leaflet';

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
export class RecentSearchesService {

  private static recentSearchesKey = 'searches';

  constructor(private storage: Storage) {}

  public async addRecentSearch(location: ISearchResultObject) {
    let currentSearches: ISearchResultObject[] = await this.readRecentSearches();
    if (currentSearches === null || currentSearches.length === 0) {
      currentSearches = [];
    }

    let pos = -1;
    for (let i = 0; i < currentSearches.length; i++) {
        if (currentSearches[i].Name === location.Name) {
            pos = i;
            break;
        }
    }

    if (pos >= 0) {
      currentSearches.splice(pos, 1);
    }

    currentSearches.unshift(location);

    if (currentSearches.length > 3) {
      currentSearches.pop();
    }

    await this.storage.set(RecentSearchesService.recentSearchesKey, JSON.stringify([currentSearches]));
    console.log(currentSearches);
  }

  public async readRecentSearches(): Promise<ISearchResultObject[]> {
    const searches = JSON.parse(await this.storage.get(RecentSearchesService.recentSearchesKey));
    return searches != null ? searches[0] : [];
  }
}
