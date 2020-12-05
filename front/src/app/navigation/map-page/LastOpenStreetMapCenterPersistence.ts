import { Storage } from '@ionic/storage';
import {Injectable} from '@angular/core';
import * as Leaflet from 'leaflet';

@Injectable({
    providedIn: 'root'
})
export class LastOpenStreetMapCenterPersistence {

    private static readonly LAST_MAP_EVENT_KEY = 'lastOpenStreetMapEventTarget';

    public static async persist(storage: Storage, mapEventData: PersistedOpenStreetMapData) {
        await storage.set(this.LAST_MAP_EVENT_KEY, mapEventData);
    }

    public async load(storage: Storage, map: Leaflet.Map, defaultZoom: number) {
        storage.ready().then(async () => {
            await storage.get(LastOpenStreetMapCenterPersistence.LAST_MAP_EVENT_KEY).then(lastMapEventData => {
                map.setView([49.452368, 11.093299], defaultZoom);
                if (lastMapEventData != null) {
                    map.flyTo(lastMapEventData.center, lastMapEventData.zoom);
                }
            });
        });
    }
}

class PersistedOpenStreetMapData {
    center: Leaflet.LatLng;
    zoom: number;
}