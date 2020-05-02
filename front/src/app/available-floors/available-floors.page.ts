import { Component, OnInit } from '@angular/core';
import {NavParams} from '@ionic/angular';
import {Components} from '@ionic/core';

@Component({
  selector: 'app-available-floors',
  templateUrl: './available-floors.page.html',
  styleUrls: ['./available-floors.page.scss'],
})
export class AvailableFloorsPage implements OnInit {

  public building: JSON;
  public floors:Array<string>;
  modal: Components.IonModal;
  constructor(navParams: NavParams) {
    this.building = JSON.parse(navParams.get('building'));
    this.floors = this.building['Floors'];
  }

  ngOnInit() {
  }

  cancel(floor: number) {
    this.modal.dismiss(floor);
  }

}
