import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {IonicDetailViewComponent} from './ionic-detail-view.component';
import {IonicModule} from '@ionic/angular';



@NgModule({
  declarations: [IonicDetailViewComponent],
  imports: [
    CommonModule, IonicModule
  ],
  exports: [IonicDetailViewComponent]
})
export class IonicDetailViewModule { }
