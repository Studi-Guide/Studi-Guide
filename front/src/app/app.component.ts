import { Component } from '@angular/core';

import { Platform } from '@ionic/angular';
import {SettingsModel} from './settings/settings.model';
import {Plugins, StatusBarStyle} from '@capacitor/core';
const { SplashScreen, StatusBar} = Plugins;

@Component({
  selector: 'app-root',
  templateUrl: 'app.component.html',
  styleUrls: ['app.component.scss']
})
export class AppComponent {
  constructor(
    private platform: Platform,
    private settingsModel: SettingsModel /* inject this to initialize global settings on startup */
  ) {
    this.initializeApp();
  }

  initializeApp() {
    if (this.platform.is('hybrid')) {
      this.platform.ready().then(() => {
        StatusBar.setStyle({style: StatusBarStyle.Light})
            .then(() => SplashScreen.hide().then());
      });
    }
  }
}
