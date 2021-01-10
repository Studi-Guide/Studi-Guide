import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {IonicHorizontalScrollableComponent} from './ionic-horizontal-scrollable/ionic-horizontal-scrollable.component';



@NgModule({
  declarations: [IonicHorizontalScrollableComponent],
  imports: [
    CommonModule,
  ],
  exports: [IonicHorizontalScrollableComponent]
})
export class IonicVirtualScrollContentModule { }
