import {AfterViewInit, Component, OnInit, ViewChild} from '@angular/core';
import {IonInput} from '@ionic/angular';
import {IRouteLocation, NavigationModel} from '../../navigationModel';

@Component({
  selector: 'app-route-input',
  templateUrl: './route-input.component.html',
  styleUrls: ['./route-input.component.scss'],
})
export class RouteInputComponent implements OnInit, AfterViewInit {

  @ViewChild('inputFrom') inputFrom: IonInput;
  @ViewChild('inputTo') inputTo: IonInput;

  constructor(
      private model: NavigationModel
  ) { }

  ngOnInit() {}

  ngAfterViewInit() {
    this.UpdateFromNavigationModel();
  }

  public UpdateFromNavigationModel() {
    this.inputFrom.value = this.model.Route.Start.Name;
    this.inputTo.value = this.model.Route.Destination.Name;
  }

  public swapInputs() {
    const tmp = this.inputFrom.value;
    this.inputFrom.value = this.inputTo.value;
    this.inputTo.value = tmp;
  }

  public get From(): IRouteLocation {
    return this.model.Route.Start;
  }

  public get To(): IRouteLocation {
    return this.model.Route.Destination;
  }

}
