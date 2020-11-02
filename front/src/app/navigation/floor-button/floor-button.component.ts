import {Component, ContentChild, OnInit, ViewChild} from '@angular/core';
import {DataService} from '../../services/data.service';

@Component({
  selector: 'app-floor-button',
  templateUrl: './floor-button.component.html',
  styleUrls: ['./floor-button.component.scss'],
})
export class FloorButtonComponent implements OnInit {

  @ViewChild('fab') fab;
  // TODO
  @ViewChild('currentBuilding') currentBuilding; // @ContentChild didn't work

  public availableFloors:string[];

  constructor(private _dataService: DataService) {}

  ngOnInit() {}

  async loadAvailableFloors() {
    console.log(this.fab.activated);
    if(!this.fab.activated) {
      const building = await this._dataService.get_building(this.currentBuilding.innerText).toPromise();
      this.availableFloors = building.Floors;
      this.fab.activated = true;
    } else {
      this.fab.activated = false;
    }
  }

}
