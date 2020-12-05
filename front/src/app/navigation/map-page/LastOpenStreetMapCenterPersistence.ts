import { Storage } from '@ionic/storage';
import {Injectable} from '@angular/core';
import * as Leaflet from 'leaflet';

@Injectable({
    providedIn: 'root'
})
export class LastOpenStreetMapCenterPersistence {

    constructor(private storage: Storage) {}

    private static readonly LAST_CENTER_KEY = 'lastOpenStreetMapCenter';

    public static async persist(storage: Storage, newCenter: Leaflet.LatLng) {
        await storage.set(this.LAST_CENTER_KEY, newCenter);
    }

    public async load(map: Leaflet.Map, zoom: number) {
        this.storage.ready().then(async () => {
            await this.storage.get(LastOpenStreetMapCenterPersistence.LAST_CENTER_KEY).then(lastCenter => {
                const center = lastCenter == null ? [49.452368, 11.093299] : lastCenter;
                map.setView(center, zoom);
            });
        });
    }

    public test(map: Leaflet.Map) {
        console.log(map.getCenter().toString());
    }
}