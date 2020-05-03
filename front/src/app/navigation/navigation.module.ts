import { IonicModule } from '@ionic/angular';
import { RouterModule } from '@angular/router';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { NavigationPage } from './navigationPage';
import {AvailableFloorsPage} from '../available-floors/available-floors.page';

@NgModule({
  imports: [
    IonicModule,
    CommonModule,
    FormsModule,
    RouterModule.forChild([{ path: '', component: NavigationPage }])
  ],
  declarations: [NavigationPage]
})
export class NavigationPageModule {}
