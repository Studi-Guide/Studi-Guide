import {Inject, Injectable} from '@angular/core';
import {DOCUMENT} from '@angular/common';
import {Platform} from '@ionic/angular';

export const environment = {
  production: true
};

@Injectable()
export class Env {

  serverUrl: string;
  production = true;

  constructor(public plt: Platform,
              @Inject(DOCUMENT) private document: Document) {
    this.serverUrl = document.location.origin;

    console.log(plt.platforms());
    if (plt.is('hybrid')){
      console.log('Android or iOS app recognized');
      this.serverUrl = 'https://studi-guide-ii.azurewebsites.net';
    }
  }

}
