import { IonicModule } from '@ionic/angular';
import { RouterModule } from '@angular/router';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { SettingsPage } from './settings.page';
import {DetailViewSelectLanguageComponent} from '../detail-view-select-language/detail-view-select-language.component';
import {DetailViewSelectLanguageModule} from '../detail-view-select-language/detail-view-select-language.module';

@NgModule({
    imports: [
        IonicModule,
        CommonModule,
        FormsModule,
        RouterModule.forChild([
            {path: '', component: SettingsPage},
            {path: 'language', component: DetailViewSelectLanguageComponent, }
        ]),
        DetailViewSelectLanguageModule
    ],
  declarations: [SettingsPage]
})
export class SettingsPageModule {}
