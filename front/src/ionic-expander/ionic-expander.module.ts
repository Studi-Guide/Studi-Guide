import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {IonicExpanderComponent} from './ionic-expander.component';
import {IonicModule} from '@ionic/angular';

@NgModule({
  declarations: [IonicExpanderComponent],
  imports: [
    CommonModule,
    IonicModule
  ],
  exports: [IonicExpanderComponent]
})
export class IonicExpanderModule { }
