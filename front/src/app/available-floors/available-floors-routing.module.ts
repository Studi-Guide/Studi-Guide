import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { AvailableFloorsPage } from './available-floors.page';

const routes: Routes = [
  {
    path: '',
    component: AvailableFloorsPage
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class AvailableFloorsPageRoutingModule {}
