import { Injectable } from '@angular/core';
import {Storage} from '@ionic/storage';

@Injectable({
  providedIn: 'root'
})
export class RecentSearchesService {

  private static recentSearchesKey = 'searches';

  constructor(private storage: Storage) {}

  public async addRecentSearch(location:string) {
    let currentSearches:string[] = await this.readRecentSearches();
    if(currentSearches === null || currentSearches.length === 0) {
      currentSearches = [];
    }

    if (currentSearches.includes(location)) {
      currentSearches.splice(currentSearches.indexOf(location), 1);
    }

    currentSearches.unshift(location);

    if (currentSearches.length > 3) {
      currentSearches.pop();
    }

    await this.storage.set(RecentSearchesService.recentSearchesKey, JSON.stringify([currentSearches]));
    console.log(currentSearches);
  }

  public async readRecentSearches() : Promise<string[]> {
    const searches = JSON.parse(await this.storage.get(RecentSearchesService.recentSearchesKey));
    return searches != null ? searches[0] : [];
  }
}
