// This file can be replaced during build by using the `fileReplacements` array.
// `ng build --prod` replaces `environment.ts` with `environment.prod.ts`.
// The list of file replacements can be found in `angular.json`.

import {Injectable} from '@angular/core';
import {Platform} from '@ionic/angular';

export const environment = {
  production: false
}

@Injectable()
export class Env {

  serverUrl = 'http://localhost:8080';
  production = false;

  constructor(public plt: Platform) {
    console.log(plt.platforms());
    if (plt.is('hybrid')){
      console.log('Android or iOS app recognized');
      // this.serverUrl ='https://studi-guide.azurewebsites.net';
    }
  }
}

/*
 * For easier debugging in development mode, you can import the following file
 * to ignore zone related error stack frames such as `zone.run`, `zoneDelegate.invokeTask`.
 *
 * This import should be commented out in production mode because it will have a negative impact
 * on performance if an error is thrown.
 */
// import 'zone.js/dist/zone-error';  // Included with Angular CLI.
