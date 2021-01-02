import { IonicModule } from '@ionic/angular';
import { RouterModule } from '@angular/router';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { NewsPage } from './news.page';
import {IonicExpanderModule} from "../../ionic-expander/ionic-expander.module";

@NgModule({
  imports: [
    IonicModule,
    CommonModule,
    FormsModule,
    RouterModule.forChild([
      {path: '', component: NewsPage}]),
    IonicExpanderModule
  ],
  declarations: [NewsPage]
})
export class NewsPageModule {}
