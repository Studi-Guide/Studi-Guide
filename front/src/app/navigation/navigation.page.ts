import {ILocation} from '../building-objects-if';
import {AfterViewInit, Component, OnInit, ViewChild} from '@angular/core';
import {IonContent, ModalController, Platform} from '@ionic/angular';
import {DataService} from '../services/data.service';
import {MapViewComponent} from './map-view/map-view.component';
import {HttpErrorResponse} from '@angular/common/http';
import {ActivatedRoute, Router} from '@angular/router';
import {SearchInputComponent} from './search-input/search-input.component';
import {DrawerState} from '../../ionic-bottom-drawer/drawer-state';
import {IonicBottomDrawerComponent} from '../../ionic-bottom-drawer/ionic-bottom-drawer.component';
import {CampusViewModel} from './campusViewModel';
import {IRouteLocation, NavigationModel} from './navigationModel';
import {Plugins} from '@capacitor/core';
import {IReceivedRoute} from '../route-objects-if';
import {RouteInputComponent} from './nav-drawer-manager/route-input/route-input.component';

const { Keyboard } = Plugins;

@Component({
    selector: 'app-navigation',
    templateUrl: 'navigation.page.html',
    styleUrls: ['navigation.page.scss']
})

export class NavigationPage implements OnInit, AfterViewInit{

    public static progressIsVisible = false;

    @ViewChild(MapViewComponent) mapView: MapViewComponent;
    @ViewChild(SearchInputComponent) searchInput: SearchInputComponent;
    @ViewChild('drawerContent') drawerContent: IonContent;
    @ViewChild('searchDrawer') searchDrawer: IonicBottomDrawerComponent;
    @ViewChild('locationDrawer') locationDrawer: IonicBottomDrawerComponent;
    @ViewChild('routeDrawer') routeDrawer: IonicBottomDrawerComponent;
    @ViewChild('changeRouteDrawer') changeRouteDrawer: IonicBottomDrawerComponent;
    @ViewChild('routeInput') routeInput: RouteInputComponent;
    public availableCampus: CampusViewModel[] = [];
    private isSubscribed = false;

    constructor(private dataService: DataService,
                private modalCtrl: ModalController,
                private  route: ActivatedRoute,
                private router: Router,
                public model: NavigationModel,
                private  platform: Platform) {
    }

    public get CurrentRoute(): IReceivedRoute {
        return this.mapView?.CurrentRoute;
    }

    async ngAfterViewInit()  {
        await Promise.all([
            this.locationDrawer.SetState(DrawerState.Hidden),
            this.routeDrawer.SetState(DrawerState.Hidden),
            this.changeRouteDrawer.SetState(DrawerState.Hidden),
            this.searchDrawer.SetState(IonicBottomDrawerComponent.GetRecommendedDrawerStateForDevice())
        ]);
    }

    ionViewDidEnter() {
        if (this.isSubscribed === false){
            // CanvasTouchHelper.RegisterPinch(this.renderer, this.canvasWrapper);
            this.isSubscribed = true;
            this.route.queryParams.subscribe(async params => {
                // discover requested location
                if (params != null && params.location != null && params.location.length > 0) {
                    this.searchInput.setDiscoverySearchbarValue(params.location);
                    await this.onDiscovery(params.location);
                    return;
                }

                // launch requested navigation
                if (params != null && params.start != null && params.start.length > 0 &&
                    params.destination != null && params.destination.length > 0) {
                    await this.showNavigation(params.start, params.destination);
                } else if (params != null && params.building != null && params.building.length > 0) {
                    const building = await this.dataService.get_building(params.building).toPromise();
                    if (building !== null) {
                        const floors = await this.dataService.get_building_floor(building.Name).toPromise();
                        await this.mapView.showFloor(
                            floors?.includes('EG') ? 'EG' : floors[0],
                            building.Name);
                        this.mapView.FitMap();
                    }
                } else {
                    await this.showDiscoveryMode();
                }
            });
        }
    }

    async ngOnInit() {

        if (this.model.availableCampus.length === 0)
        {
            const campus = await this.dataService.get_campus_search().toPromise();
            for (const c of campus) {
                this.model.availableCampus.push(new CampusViewModel(c));
            }
        }

    }

    public async onDiscovery(searchInput: string) {
        this.model.errorMessage = '';
        NavigationPage.progressIsVisible = true;
        try {
            const location = await this.mapView.showDiscoveryLocation(searchInput);
            const building = await this.dataService.get_building(location.Building).toPromise();
            await this.model.addRecentSearchLocation(location, {lat: 0, lng: 0}, building);

            await this.showLocationDrawer(location);
        } catch (ex) {
            this.handleInputError(ex, searchInput);
        } finally {
            NavigationPage.progressIsVisible = false;
        }
    }

    public async onRoute(routeInput: string[]) {
        this.model.errorMessage = '';
        NavigationPage.progressIsVisible = true;
        try {
            const startLocation = await this.dataService.get_location(routeInput[0]).toPromise<ILocation>();
            const route = await this.dataService.get_route(routeInput[0], routeInput[1]).toPromise();
            route.Distance = Math.round(route.Distance / MapViewComponent.DISTANCE_SCALAR);
            route.Start.Name = startLocation.Name;
            route.End.Name = routeInput[1];
            await this.mapView.showRoute(route, startLocation);
            this.mapView.CenterMap(startLocation.PathNode.Coordinate.X, startLocation.PathNode.Coordinate.Y);
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
            NavigationPage.progressIsVisible = false;
        }
    }

    public onDrawerStateChange(state: DrawerState) {
        // in case the view is not initialized
        if (this.drawerContent === undefined) {
            return;
        }

        this.drawerContent.scrollY = state === DrawerState.Top;
    }

    public async showLocationDrawer(location: ILocation) {
        await this.locationDrawer.SetState(DrawerState.Hidden);
        this.model.latestSearchResult.Name = location.Name;
        this.model.latestSearchResult.Description = location.Description;
        this.model.latestSearchResult.Information = location.Tags ?
            [['Tags: ', location.Tags.join(',')]] :
            [];

        this.model.latestSearchResult.Images = location.Images;

        const isHybrid = this.platform.is('hybrid');
        if (isHybrid) {
            await Keyboard.hide();
        }

        await this.searchDrawer.SetState(DrawerState.Hidden);
        await this.locationDrawer.SetState(IonicBottomDrawerComponent.GetRecommendedDrawerStateForDevice());
    }

    public async onCloseLocationDrawer(event: any) {
        this.mapView.RefreshMap();
        await this.locationDrawer.SetState(DrawerState.Hidden);
        if (!this.CurrentRoute) {
            await this.searchDrawer.SetState(IonicBottomDrawerComponent.GetRecommendedDrawerStateForDevice());
        } else {
            await this.routeDrawer.SetState(IonicBottomDrawerComponent.GetRecommendedDrawerStateForDevice());
        }
    }

    private handleInputError(ex, searchInput: string) {
        console.log(ex);
        if (ex instanceof HttpErrorResponse) {
            const httpError = ex as HttpErrorResponse;
            if (httpError.status === 400) {
                this.model.errorMessage = 'Studi-Guide can\'t find ' + searchInput;
            } else {
                this.model.errorMessage = httpError.message;
            }
        } else {
            this.model.errorMessage = (ex as Error).message;
        }

        // TODO Remove this code when discovery mode is finished
        if (this.mapView != null) {
            this.mapView.clearMapCanvas();
        }
    }

    async navigationBtnClick() {
        if (this.model.latestSearchResult != null) {
            try {
                const startLocation = await this.dataService.get_location(this.model.latestSearchResult.Name).toPromise<ILocation>();
                await this.showNavigation(startLocation.Building + '.Entrance', startLocation.Name);
            }catch (ex) {
                let inputError = '';
                if (ex instanceof HttpErrorResponse) {
                    const errorString = (ex as HttpErrorResponse).error.message;
                    if (errorString.includes(this.model.latestSearchResult.Name)) {
                        inputError = this.model.latestSearchResult.Name;
                    }
                }
                this.handleInputError(ex, inputError.length === 0 ? this.model.latestSearchResult.Name : inputError);
            } finally {
                NavigationPage.progressIsVisible = false;
            }
        }
    }

    private async showNavigation(start: string, destination: string) {
        this.searchInput.setDiscoverySearchbarValue(destination);
        this.searchInput.setStartSearchbarValue(start);
        await this.onRoute([start, destination]);
        await this.onCloseLocationDrawer(null);
    }

    private async showDiscoveryMode() {
        if (this.mapView.CurrentRoute == null && this.mapView.CurrentBuilding == null) {
            // STDG-138 load base map
            await this.mapView.showDiscoveryMap('', 'EG');

            // Coordinates of KA.013
            this.mapView.CenterMap(310, 550);
        }
    }

    public async recentSearchClick(locationStr: string) {
        await this.router.navigate(['tabs/navigation/detail'], { queryParams: { location: locationStr } });
    }

    public async presentMapPage() {
        await this.router.navigate(['tabs/navigation/']);
    }

    public get ProgressIsVisible(): boolean {
        return NavigationPage.progressIsVisible;
    }

    async onFloorChanged(event: any) {
        this.searchInput.clearStartInput();
        this.searchInput.clearDestinationInput();
        await this.onCloseLocationDrawer(null);
    }

    // custom event handler
    async onSearchFocus($event: string) {
        if (this.platform.is('hybrid')) {
            await this.searchDrawer.SetState(DrawerState.Top);
        }
    }

   async onCloseRouteDrawer() {
       this.mapView.ClearRoute();
       this.mapView.RefreshMap();
       await this.searchDrawer.SetState(IonicBottomDrawerComponent.GetRecommendedDrawerStateForDevice());
       await this.routeDrawer.SetState(DrawerState.Hidden);
       await this.locationDrawer.SetState(DrawerState.Hidden);
    }

    public async onChangeRouteStartEndClick() {
        this.routeInput.myLocationEnabled = false;
        this.routeInput.routeLocationFrom = { Name: this.CurrentRoute?.Start?.Name ?? '', LatLng: { lat: 0, lng : 0}};
        this.routeInput.routeLocationTo = { Name: this.CurrentRoute?.End?.Name ?? '', LatLng: { lat: 0, lng : 0}};
        this.routeInput.updateInputValues();
        await this.routeDrawer.SetState(DrawerState.Hidden);
        await this.changeRouteDrawer.SetState(IonicBottomDrawerComponent.GetRecommendedDrawerStateForDevice());
    }

    public async onCancelChangeRoute() {
        await this.changeRouteDrawer.SetState(DrawerState.Hidden);
        await this.routeDrawer.SetState(IonicBottomDrawerComponent.GetRecommendedDrawerStateForDevice());
    }

    async onChangeRoute(route: any[]) {
        await this.onCancelChangeRoute();
        await this.onRoute( [route[0].toString(), route[1].toString()]);
    }
}
