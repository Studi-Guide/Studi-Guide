import { IonicModule } from '@ionic/angular';
import { RouterModule } from '@angular/router';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { SettingsPage } from './settings.page';
import {DetailViewSelectLanguageComponent} from '../detail-view-select-language/detail-view-select-language.component';
import {DetailViewSelectLanguageModule} from '../detail-view-select-language/detail-view-select-language.module';
import {StdgTooltipComponent} from '../stdg-tooltip/stdg-tooltip.component';
import { StdgTooltipModule } from '../stdg-tooltip/stdg-tooltip.module';

@NgModule({
    imports: [
        IonicModule,
        CommonModule,
        FormsModule,
        RouterModule.forChild([
            {path: '', component: SettingsPage},
            {path: 'language', component: DetailViewSelectLanguageComponent, }
        ]),
        DetailViewSelectLanguageModule,
        StdgTooltipModule
    ],
  declarations: [SettingsPage]
})
export class SettingsPageModule {
    constructor() {
        (new StdgTooltipComponent('dark', 5, 0)).init();
    }
}
