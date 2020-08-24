import { IonicModule } from '@ionic/angular';
import { RouterModule } from '@angular/router';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { Tab4Page } from './tab4.page';
import {IonicDetailViewModule} from "../../ionic-detail-view/ionic-detail-view.module";
import {IonicDetailViewComponent} from "../../ionic-detail-view/ionic-detail-view.component";

@NgModule({
    imports: [
        IonicModule,
        CommonModule,
        FormsModule,
        RouterModule.forChild([
            {path: '', component: Tab4Page},
            {path: 'language', component: IonicDetailViewComponent, }
        ]),
        IonicDetailViewModule
    ],
  declarations: [Tab4Page]
})
export class Tab4PageModule {}
