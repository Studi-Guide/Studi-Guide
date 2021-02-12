import { IonicModule } from '@ionic/angular';
import { RouterModule } from '@angular/router';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { LunchPage } from './lunch.page';
import {IonicExpanderComponent} from '../../ionic-expander/ionic-expander.component';
import {IonicExpanderModule} from "../../ionic-expander/ionic-expander.module";

@NgModule({
  imports: [
    IonicModule,
    CommonModule,
    FormsModule,
    RouterModule.forChild([{path: '', component: LunchPage}]),
    IonicExpanderModule
  ],
  declarations: [LunchPage]
})
export class LunchModule {}
