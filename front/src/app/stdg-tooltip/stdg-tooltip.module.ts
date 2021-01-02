import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {IonicModule} from '@ionic/angular';
import {StdgTooltipComponent} from './stdg-tooltip.component';



@NgModule({
  declarations: [StdgTooltipComponent],
  imports: [
    CommonModule, IonicModule
  ],
  exports: [StdgTooltipComponent]
})
export class StdgTooltipModule { }
