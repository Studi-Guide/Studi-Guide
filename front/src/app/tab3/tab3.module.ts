import { IonicModule } from '@ionic/angular';
import { RouterModule } from '@angular/router';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { Tab3Page } from './tab3.page';
import {IonicExpanderComponent} from '../../ionic-expander/ionic-expander.component';
import {IonicExpanderModule} from "../../ionic-expander/ionic-expander.module";

@NgModule({
  imports: [
    IonicModule,
    CommonModule,
    FormsModule,
    RouterModule.forChild([{path: '', component: Tab3Page}]),
    IonicExpanderModule
  ],
  declarations: [Tab3Page]
})
export class Tab3PageModule {}
