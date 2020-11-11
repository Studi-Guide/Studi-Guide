import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {IonicBottomDrawerComponent} from './ionic-bottom-drawer.component';
import {IonicModule} from '@ionic/angular';



@NgModule({
  declarations: [IonicBottomDrawerComponent],
  imports: [
    CommonModule, IonicModule
  ],
  exports: [IonicBottomDrawerComponent]
})
export class IonicBottomDrawerModule { }
