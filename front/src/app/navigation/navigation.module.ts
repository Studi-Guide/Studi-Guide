import { IonicModule } from '@ionic/angular';
import { RouterModule } from '@angular/router';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { NavigationPage } from './navigationPage';
import {SearchInputComponent} from './search-input/search-input.component';
import {MapViewComponent} from './map-view/map-view.component';
import {IonicBottomDrawerModule} from '../../ionic-bottom-drawer/ionic-bottom-drawer.module';
import {MapPageComponent} from './map-page/map-page.component';
import {FloorButtonComponent} from './floor-button/floor-button.component';

@NgModule({
  imports: [
    IonicModule,
    CommonModule,
    FormsModule,
    RouterModule.forChild([
        { path: '', component: NavigationPage},
        { path: 'map', component: MapPageComponent },
     //   { path: 'navigation/:start/:destination', component: NavigationPage}
        ]),
    IonicBottomDrawerModule
  ],
    declarations: [NavigationPage, SearchInputComponent, MapViewComponent, MapPageComponent, FloorButtonComponent]
})
export class NavigationPageModule {}
