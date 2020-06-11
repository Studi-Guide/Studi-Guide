import {BuildingData, Location, MapItem, PathNode} from '../building-objects-if';
// import {testDataRooms, testDataPathNodes} from './test-building-data';
import {Component, ViewChild} from '@angular/core';
import {ModalController} from '@ionic/angular';
import {DataService} from '../services/data.service';
import {FloorMapRenderer} from './map-view/floorMapRenderer';
import {NaviRouteRenderer, ReceivedRoute} from './map-view/naviRouteRenderer';
import {AvailableFloorsPage} from '../available-floors/available-floors.page';
import {MapViewComponent} from './map-view/map-view.component';

@Component({
  selector: 'app-navigation',
  templateUrl: 'navigation.page.html',
  styleUrls: ['navigation.page.scss']
})

export class NavigationPage {

  @ViewChild(MapViewComponent) mapView:MapViewComponent;

  public progressIsVisible = false;
  public availableFloorsBtnIsVisible = false;


  constructor(private dataService: DataService,
              private modalCtrl: ModalController) {
  }


  public async onDiscovery(searchInput: string) {
    this.progressIsVisible = true;
    await this.mapView.showDiscoveryLocation(searchInput);
    this.progressIsVisible = false;
    this.availableFloorsBtnIsVisible = true;
  }

  public async onRoute(routeInput: string[]) {
    this.progressIsVisible = true;
    await this.mapView.showRoute(routeInput[0], routeInput[1]);
    this.progressIsVisible = false;
    this.availableFloorsBtnIsVisible = true;
  }

  private isEmptyOrSpaces(str) {
    return str === null || str.match(/^ *$/) !== null;
  }

  async presentAvailableFloorModal() {
    let floors = new Array<string>();

    const building = await this.dataService.get_building(this.mapView.CurrentBuilding).toPromise<BuildingData>();
    floors = floors.concat(building.Floors);

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
}
