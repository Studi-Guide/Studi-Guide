import { IonicModule } from '@ionic/angular';
import { RouterModule } from '@angular/router';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { SchedulePage } from './schedule.page';
import {LoginComponent} from './login/login.component';

@NgModule({
  imports: [
    IonicModule,
    CommonModule,
    FormsModule,
    RouterModule.forChild([
      { path: '', component: SchedulePage },
      { path: 'login/', component: SchedulePage }])
  ],
  declarations: [SchedulePage, LoginComponent]
})
export class SchedulePageModule {}
