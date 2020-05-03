import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

import { IonicModule } from '@ionic/angular';

import { AvailableFloorsPageRoutingModule } from './available-floors-routing.module';

import { AvailableFloorsPage } from './available-floors.page';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    IonicModule,
    AvailableFloorsPageRoutingModule
  ],
  declarations: [AvailableFloorsPage]
})
export class AvailableFloorsPageModule {}
