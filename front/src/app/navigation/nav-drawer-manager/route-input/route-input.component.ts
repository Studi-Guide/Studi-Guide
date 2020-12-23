import {Component, OnInit, ViewChild} from '@angular/core';
import {IonInput} from '@ionic/angular';

@Component({
  selector: 'app-route-input',
  templateUrl: './route-input.component.html',
  styleUrls: ['./route-input.component.scss'],
})
export class RouteInputComponent implements OnInit {

  @ViewChild('inputFrom') inputFrom:IonInput;
  @ViewChild('inputTo') inputTo:IonInput;

  constructor() { }

  ngOnInit() {}

  public swapInputs() {
    const tmp = this.inputFrom.value;
    this.inputFrom.value = this.inputTo.value;
    this.inputTo.value = tmp;
  }

}
