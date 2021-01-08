import { IonicModule } from '@ionic/angular';
import { RouterModule } from '@angular/router';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { SettingsPage } from './settings.page';
import {DetailViewSelectLanguageComponent} from '../detail-view-select-language/detail-view-select-language.component';
import {DetailViewSelectLanguageModule} from '../detail-view-select-language/detail-view-select-language.module';
import {DetailViewSelectAppearanceComponent} from './detail-view-select-appearance/detail-view-select-appearance.component';

@NgModule({
    imports: [
        IonicModule,
        CommonModule,
        FormsModule,
        RouterModule.forChild([
            {path: '', component: SettingsPage},
            {path: 'language', component: DetailViewSelectLanguageComponent, },
            {path: 'appearance', component: DetailViewSelectAppearanceComponent, }
        ]),
        DetailViewSelectLanguageModule
    ],
  declarations: [SettingsPage, DetailViewSelectAppearanceComponent]
})
export class SettingsPageModule {}
