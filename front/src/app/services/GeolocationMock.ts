import { Geolocation } from '@ionic-native/geolocation/ngx';

export class GeolocationMock extends Geolocation {
    latitude = 49.4531;
    longitude = 11.0919;

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
