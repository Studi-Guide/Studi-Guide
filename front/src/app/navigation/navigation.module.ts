import { IonicModule } from '@ionic/angular';
import { RouterModule } from '@angular/router';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { NavigationPage } from './navigationPage';
import {SearchInputComponent} from "./search-input/search-input.component";

@NgModule({
  imports: [
    IonicModule,
    CommonModule,
    FormsModule,
    RouterModule.forChild([{ path: '', component: NavigationPage }])
  ],
    declarations: [NavigationPage, SearchInputComponent]
})
export class NavigationPageModule {}
