import { IonicModule } from '@ionic/angular';
import { RouterModule } from '@angular/router';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { NavigationPage } from './navigation.page';
import {SearchInputComponent} from './search-input/search-input.component';
import {MapViewComponent} from './map-view/map-view.component';
import {IonicBottomDrawerModule} from '../../ionic-bottom-drawer/ionic-bottom-drawer.module';
import {MapPageComponent} from './map-page/map-page.component';
import {FloorButtonComponent} from './floor-button/floor-button.component';
import {NavigationInstructionSlidesComponent} from './navigation-instruction-slides/navigation-instruction-slides.component';
import {NavDrawerManagerComponent} from './nav-drawer-manager/nav-drawer-manager.component';

@NgModule({
  imports: [
    IonicModule,
    CommonModule,
    FormsModule,
    RouterModule.forChild([
        { path: 'detail', component: NavigationPage},
        { path: '', component: MapPageComponent },
        ]),
    IonicBottomDrawerModule
  ],
    declarations: [
        NavigationPage,
        SearchInputComponent,
        MapViewComponent,
        MapPageComponent,
        FloorButtonComponent,
        NavigationInstructionSlidesComponent,
        NavDrawerManagerComponent
    ]
})
export class NavigationPageModule {}
