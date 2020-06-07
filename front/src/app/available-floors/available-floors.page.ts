import { Component, OnInit } from '@angular/core';
import {NavParams} from '@ionic/angular';
import {Components} from '@ionic/core';

@Component({
  selector: 'app-available-floors',
  templateUrl: './available-floors.page.html',
  styleUrls: ['./available-floors.page.scss'],
})
export class AvailableFloorsPage implements OnInit {

  public floors:string[];
  modal: Components.IonModal;
  constructor(navParams: NavParams) {
    this.floors = navParams.get('floors');
  }

  ngOnInit() {
  }

  cancel(floor: number) {
    this.modal.dismiss(floor);
  }

}
