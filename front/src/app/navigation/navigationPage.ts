import {BuildingData, Location, MapItem, PathNode} from '../building-objects-if';
import {Component, ViewChild} from '@angular/core';
import {ModalController} from '@ionic/angular';
import {DataService} from '../services/data.service';
import {AvailableFloorsPage} from '../available-floors/available-floors.page';
import {MapViewComponent} from './map-view/map-view.component';
import {HttpErrorResponse} from '@angular/common/http';

@Component({
  selector: 'app-navigation',
  templateUrl: 'navigation.page.html',
  styleUrls: ['navigation.page.scss']
})

export class NavigationPage {

  @ViewChild(MapViewComponent) mapView: MapViewComponent;

  public progressIsVisible = false;
  public availableFloorsBtnIsVisible = false;
  public errorMessage: string;

  constructor(private dataService: DataService,
              private modalCtrl: ModalController) {
  }


  public async onDiscovery(searchInput: string) {
    this.errorMessage = '';
    this.progressIsVisible = true;
    try {
      await this.mapView.showDiscoveryLocation(searchInput);
      this.availableFloorsBtnIsVisible = true;
    } catch (ex) {
      this.handleInputError(ex, searchInput);
    } finally {
      this.progressIsVisible = false;
    }
  }

  public async onRoute(routeInput: string[]) {
    this.progressIsVisible = true;
    try {
      await this.mapView.showRoute(routeInput[0], routeInput[1]);
      this.availableFloorsBtnIsVisible = true;
    } catch (ex) {
      let inputError = '';
      if (ex instanceof HttpErrorResponse) {
        const errorString = (ex as HttpErrorResponse).error.message;
        if (errorString.includes(routeInput[0])) {
          inputError = routeInput[0];
        }

        if (errorString.includes(routeInput[1])) {
          if (inputError.length > 0) {
            inputError += ' and ';
          }

          inputError += routeInput[1];
        }
      }

      this.handleInputError(ex, inputError.length === 0 ? routeInput.toString() : inputError);
    } finally {
      this.progressIsVisible = false;
    }
  }

  async presentAvailableFloorModal() {
    let floors = new Array<string>();

    if (this.mapView.CurrentRoute == null) {
      const building = await this.dataService.get_building(this.mapView.CurrentBuilding).toPromise<BuildingData>();
      floors = floors.concat(building.Floors);
    } else {
      // get all floors from all buildings on the route
      for (const routeSection of this.mapView.CurrentRoute.RouteSections) {
        const building = await this.dataService.get_building(routeSection.Building).toPromise<BuildingData>();
        floors = floors.concat(building.Floors);
      }
      // distinct array
      floors = floors.filter((n, i) => floors.indexOf(n) === i);
    }

    const availableFloorModal = await this.modalCtrl.create({
      component: AvailableFloorsPage,
      cssClass: 'floor-modal',
      componentProps: {
        floors
      }
    })
    await availableFloorModal.present();

    const data = await availableFloorModal.onDidDismiss()
    if (data.data) {
      this.progressIsVisible = true;
      await this.mapView.showFloor(this.mapView.CurrentBuilding, data.data);
      this.progressIsVisible = false;
    }
  }

  private handleInputError(ex, searchInput: string) {
    this.availableFloorsBtnIsVisible = false;
    if (ex instanceof HttpErrorResponse) {
      const httpError = ex as HttpErrorResponse;
      if (httpError.status === 400) {
        this.errorMessage = 'Studi-Guide can\'t find ' + searchInput;
      } else {
        this.errorMessage = httpError.message;
      }
      // Remove this code when discovery mode is finished
      this.mapView.clearMapCanvas();
    }
  }
}
