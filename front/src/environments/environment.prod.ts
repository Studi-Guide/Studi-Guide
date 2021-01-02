import {Inject, Injectable} from '@angular/core';
import {DOCUMENT} from '@angular/common';
import {Platform} from '@ionic/angular';

export const environment = {
  production: true
};

@Injectable()
export class Env {
  serverUrl = 'https://studi-guide-ii.azurewebsites.net';
  production = true;
}
