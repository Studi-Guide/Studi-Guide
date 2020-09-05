import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {IonicModule} from '@ionic/angular';
import {DetailViewSelectLanguageComponent} from './detail-view-select-language.component';



@NgModule({
  declarations: [DetailViewSelectLanguageComponent],
  imports: [
    CommonModule, IonicModule
  ],
  exports: [DetailViewSelectLanguageComponent]
})
export class DetailViewSelectLanguageModule { }
