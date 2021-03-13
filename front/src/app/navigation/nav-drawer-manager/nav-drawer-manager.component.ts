import {AfterViewInit, Component, EventEmitter, Input, Output, ViewChild} from '@angular/core';
import {IonContent, Platform} from '@ionic/angular';
import {IonicBottomDrawerComponent} from '../../../ionic-bottom-drawer/ionic-bottom-drawer.component';
import {IRouteLocation, NavigationModel} from '../navigationModel';
import {DrawerState} from '../../../ionic-bottom-drawer/drawer-state';
import {CampusViewModel} from '../campusViewModel';
import {INavigationInstruction} from '../navigation-instruction-slides/navigation-instruction-if';
import {RouteInputComponent} from './route-input/route-input.component';

export enum NavDrawerState {
  SearchView,
  LocationView,
  RouteView,
  InNavigationView,
  ChangeRouteView
}

@Component({
  selector: 'app-nav-drawer-manager',
  templateUrl: './nav-drawer-manager.component.html',
  styleUrls: ['./nav-drawer-manager.component.scss'],
})
export class NavDrawerManagerComponent implements AfterViewInit {

  @Input() state: NavDrawerState = NavDrawerState.SearchView;

  @Output() stateChange: EventEmitter<NavDrawerState> = new EventEmitter<NavDrawerState>();

  @Output() campusClick: EventEmitter<CampusViewModel> = new EventEmitter<CampusViewModel>();

  @Output() detailsClick: EventEmitter<any> = new EventEmitter<any>();

  @Output() navInstructionClick: EventEmitter<INavigationInstruction> = new EventEmitter<INavigationInstruction>();

  @Output() search: EventEmitter<string> = new EventEmitter<string>();

  @Output() newRoute: EventEmitter<IRouteLocation[]> = new EventEmitter<IRouteLocation[]>();

  @ViewChild('drawerContent') drawerContent: IonContent;
  @ViewChild('searchDrawer') searchDrawer: IonicBottomDrawerComponent;
  @ViewChild('locationDrawer') locationDrawer: IonicBottomDrawerComponent;
  @ViewChild('routeDrawer') routeDrawer: IonicBottomDrawerComponent;
  @ViewChild('inNavigationDrawer') inNavigationDrawer: IonicBottomDrawerComponent;
  @ViewChild('changeRouteDrawer') changeRouteDrawer: IonicBottomDrawerComponent;

  @ViewChild('routeInput') routeInput: RouteInputComponent;

  constructor(
      public model: NavigationModel,
      private platform: Platform) { }

  async ngAfterViewInit() {
    await Promise.all([
      this.locationDrawer.SetState(DrawerState.Hidden),
      this.routeDrawer.SetState(DrawerState.Hidden),
      this.inNavigationDrawer.SetState(DrawerState.Hidden),
      this.changeRouteDrawer.SetState(DrawerState.Hidden, false),
      this.searchDrawer.SetState(DrawerState.Hidden)]);
  }

  public async SetState(newState: NavDrawerState, shouldEmit = false) {

    switch (this.state) {
      case NavDrawerState.SearchView:
        await this.searchDrawer.SetState(DrawerState.Hidden);
        break;
      case NavDrawerState.LocationView:
        await this.locationDrawer.SetState(DrawerState.Hidden);
        break;
      case NavDrawerState.RouteView:
        await this.routeDrawer.SetState(DrawerState.Hidden);
        break;
      case NavDrawerState.InNavigationView:
        await this.inNavigationDrawer.SetState(DrawerState.Hidden);
        break;
      case NavDrawerState.ChangeRouteView:
        await this.changeRouteDrawer.SetState(DrawerState.Hidden, false);
        break;
    }

    this.state = newState;

    switch (this.state) {
      case NavDrawerState.SearchView:
        await this.searchDrawer.SetState(IonicBottomDrawerComponent.GetRecommendedDrawerStateForDevice());
        break;
      case NavDrawerState.LocationView:
        await this.locationDrawer.SetState(IonicBottomDrawerComponent.GetRecommendedDrawerStateForDevice());
        break;
      case NavDrawerState.RouteView:
        await this.routeDrawer.SetState(IonicBottomDrawerComponent.GetRecommendedDrawerStateForDevice());
        break;
      case NavDrawerState.InNavigationView:
        await this.inNavigationDrawer.SetState(IonicBottomDrawerComponent.GetRecommendedDrawerStateForDevice());
        break;
      case NavDrawerState.ChangeRouteView:
        await this.changeRouteDrawer.SetState(IonicBottomDrawerComponent.GetRecommendedDrawerStateForDevice());
        break;
    }

    if (shouldEmit) {
      this.stateChange.emit(this.state);
    }
  }

  public async onRouteClick() {
    this.stateChange.emit(NavDrawerState.RouteView);
  }

  public async onLaunchNavigationClick() {
    this.stateChange.emit(NavDrawerState.InNavigationView);
  }

  public async onEndNavigationClick() {
    this.stateChange.emit(NavDrawerState.LocationView);
  }

  // custom event handler
  async onSearchFocus($event: string) {
    if (this.platform.is('hybrid')) {
      await this.searchDrawer.SetState(DrawerState.Top);
    }
  }

  public emitRoute(s: string[]) {
  }

  public async onCloseLocationDrawer() {
    this.stateChange.emit(NavDrawerState.SearchView);
  }

  public async onCloseRouteDrawer() {
    this.stateChange.emit(NavDrawerState.LocationView);
  }

  public async onChangeRouteStartEndClick() {
    this.routeInput.UpdateFromNavigationModel();
    await this.SetState(NavDrawerState.ChangeRouteView, true);
    await this.routeInput.SetFocus();
  }

  // drawer state change handler

  public onSearchDrawerStateChange(state: DrawerState) {
    // in case the view is not initialized
    if (this.drawerContent === undefined) {
      return;
    }

    this.drawerContent.scrollY = state === DrawerState.Top;
  }

  public async onChangeRouteDrawerStateChange(state: DrawerState) {
    console.log(state);
    if (state === DrawerState.Hidden) {
      await this.SetState(NavDrawerState.RouteView, true);
    }
  }

  public async onCancelChangeRoute() {
    await this.SetState(NavDrawerState.RouteView, false);
  }

  public UseDrawerForNavigation(): boolean {
    return (IonicBottomDrawerComponent.GetRecommendedDrawerStateForDevice() === DrawerState.Top);
  }

}
