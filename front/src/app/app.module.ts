import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { RouteReuseStrategy } from '@angular/router';

import { IonicModule, IonicRouteStrategy } from '@ionic/angular';
import { SplashScreen } from '@ionic-native/splash-screen/ngx';
import { StatusBar } from '@ionic-native/status-bar/ngx';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

import { HttpClientModule } from '@angular/common/http';
import {Env} from '../environments/environment';
import {IonicStorageModule} from '@ionic/storage';
import {WINDOW_PROVIDERS} from './services/windowProvider';
import { LeafletModule } from '@asymmetrik/ngx-leaflet';
import { Geolocation } from '@ionic-native/geolocation/ngx';
import {GeolocationMock} from './services/GeolocationMock';

@NgModule({
  declarations: [AppComponent],
  entryComponents: [],
  imports: [
    BrowserModule,
    HttpClientModule,
    LeafletModule,
    IonicModule.forRoot(),
    AppRoutingModule,
    IonicStorageModule.forRoot()
  ],
  providers: [
    StatusBar,
    SplashScreen,
    Env,
    WINDOW_PROVIDERS,
    { provide: RouteReuseStrategy, useClass: IonicRouteStrategy },
      // TODO remove on release
    { provide: Geolocation, useClass: GeolocationMock}
    // Geolocation
  ],
  bootstrap: [AppComponent]
})
export class AppModule {}
