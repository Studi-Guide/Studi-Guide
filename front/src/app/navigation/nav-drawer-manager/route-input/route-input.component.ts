import {AfterViewInit, Component, OnInit, ViewChild} from '@angular/core';
import {IonInput} from '@ionic/angular';
import {IRouteLocation, NavigationModel} from '../../navigationModel';

enum MyLocationInInput {
  No,
  From,
  To
}

@Component({
  selector: 'app-route-input',
  templateUrl: './route-input.component.html',
  styleUrls: ['./route-input.component.scss'],
})
export class RouteInputComponent implements OnInit, AfterViewInit {

  @ViewChild('inputFrom') inputFrom: IonInput;
  @ViewChild('inputTo') inputTo: IonInput;

  constructor(
      public model: NavigationModel
  ) { }

  private inputToCurrentlyActive = false;
  public myLocationInInput = MyLocationInInput.From;

  private routeLocationTo: IRouteLocation;
  private routeLocationFrom: IRouteLocation;

  ngOnInit() {}

  async ngAfterViewInit() {
    this.UpdateFromNavigationModel();
    this.inputToCurrentlyActive = false;
    await this.inputFrom.setFocus();
  }

  public async SetFocus() {
    if (this.inputToCurrentlyActive) {
      await this.inputTo.setFocus();
    } else {
      await this.inputFrom.setFocus();
    }
  }

  private async toggleActiveInput() {
    this.inputToCurrentlyActive = !this.inputToCurrentlyActive;
    await this.SetFocus();
  }

  private updateInputValues() {
    this.inputFrom.value = this.routeLocationFrom.Name;
    this.inputTo.value = this.routeLocationTo.Name;
  }

  public UpdateFromNavigationModel() {
    this.routeLocationFrom = this.model.Route.Start;
    this.routeLocationTo = this.model.Route.Destination;
    this.updateInputValues();
  }

  public showMyLocation(): boolean {
    return this.myLocationInInput === MyLocationInInput.No;
  }

  public onInputFromFocus() {
    this.inputToCurrentlyActive = false;
  }

  public onInputToFocus() {
    this.inputToCurrentlyActive = true;
  }

  public swapInputs() {
    const tmp = this.routeLocationFrom;
    this.routeLocationFrom = this.routeLocationTo;
    this.routeLocationTo = tmp;

    if (this.myLocationInInput !== MyLocationInInput.No) {
      this.myLocationInInput = this.myLocationInInput === MyLocationInInput.From ? MyLocationInInput.To : MyLocationInInput.From;
    }

    this.updateInputValues();
  }

  public async listRecentSearchClick(s: string) {
    const location = {
      Name: s,
      LatLng: {lat: 0, lng: 0}
    };
    if (this.inputToCurrentlyActive) {
      this.routeLocationTo = location;
      if (this.myLocationInInput === MyLocationInInput.To) {
        this.myLocationInInput = MyLocationInInput.No;
      }
    } else {
      this.routeLocationFrom = location;
      if (this.myLocationInInput === MyLocationInInput.From) {
        this.myLocationInInput = MyLocationInInput.No;
      }
    }

    this.updateInputValues();
    await this.toggleActiveInput();
  }

  public async listMyLocationClick() {
    const location = {
      Name: 'My Location',
      LatLng: {lat: 0, lng: 0}
    };
    if (this.inputToCurrentlyActive) {
      this.routeLocationTo = location;
      this.myLocationInInput = MyLocationInInput.To;
    } else {
      this.routeLocationFrom = location;
      this.myLocationInInput = MyLocationInInput.From;
    }

    this.updateInputValues();
    await this.toggleActiveInput();
  }

  public get From(): IRouteLocation {
    return this.routeLocationFrom;
  }

  public get To(): IRouteLocation {
    return this.routeLocationTo;
  }

}
