import {Component, OnInit, ViewChild} from '@angular/core';

@Component({
  selector: 'app-floor-button',
  templateUrl: './floor-button.component.html',
  styleUrls: ['./floor-button.component.scss'],
})
export class FloorButtonComponent implements OnInit {

  @ViewChild('fab') fab;
  @ViewChild('currentBuilding') currentBuilding;

  public availableFloors:string[];

  constructor() { }

  ngOnInit() {}

  loadAvailableFloors() {
    if(this.fab.activated) {
      // TODO load available floors for selected building
      console.log(this.currentBuilding);
    }
  }

}
