import { Geolocation } from '@ionic-native/geolocation/ngx';
import {Injectable} from '@angular/core';

@Injectable()
export class GeolocationMock extends Geolocation {
    latitude = 49.44667;
    longitude = 11.08164;

    // @ts-ignore
    getCurrentPosition() {
        return new Promise((resolve, reject) => {
            resolve({
                coords: {
                    latitude: this.latitude,
                    longitude: this.longitude
                }
            });
        });
    }

    setLatitude(latitude) {
        this.latitude = latitude;
        console.log('Latitude changed to: ', latitude);
    }

    setLongitude(longitude) {
        this.longitude = longitude;
        console.log('Longitude changed to: ', longitude);
    }
}
