import {IBuilding, ILocation} from '../building-objects-if';
import {AfterViewInit, Component, ElementRef, OnInit, Renderer2, ViewChild} from '@angular/core';
import {IonContent, ModalController} from '@ionic/angular';
import {DataService} from '../services/data.service';
import {AvailableFloorsPage} from '../available-floors/available-floors.page';
import {MapViewComponent} from './map-view/map-view.component';
import {HttpErrorResponse} from '@angular/common/http';
import {ActivatedRoute, Router} from '@angular/router';
import {SearchInputComponent} from './search-input/search-input.component';
import {DrawerState} from '../../ionic-bottom-drawer/drawer-state';
import {IonicBottomDrawerComponent} from '../../ionic-bottom-drawer/ionic-bottom-drawer.component';
import {Storage} from '@ionic/storage';
import {CanvasTouchHelper} from '../services/CanvasTouchHelper';
import {CampusViewModel} from './campusViewModel';
import {NavigationModel} from './navigationModel';
import {SearchResultProvider} from '../services/searchResultProvider';

@Component({
    selector: 'app-navigation',
    templateUrl: 'navigation.page.html',
    styleUrls: ['navigation.page.scss']
})

export class NavigationPage implements OnInit, AfterViewInit{

    @ViewChild(MapViewComponent) mapView: MapViewComponent;
    @ViewChild(SearchInputComponent) searchInput: SearchInputComponent;
    @ViewChild('drawerContent') drawerContent : IonContent;
    @ViewChild('searchDrawer') searchDrawer : IonicBottomDrawerComponent;
    @ViewChild('locationDrawer') locationDrawer : IonicBottomDrawerComponent;
    @ViewChild('canvasWrapper', {read: ElementRef}) private canvasWrapper: ElementRef;

    public progressIsVisible = false;
    public availableFloorsBtnIsVisible = false;
    public errorMessage: string;
    public availableCampus: CampusViewModel[] = [];
    private isSubscripted = false;

    constructor(private dataService: DataService,
                private modalCtrl: ModalController,
                private  route: ActivatedRoute,
                private router: Router,
                private storage: Storage,
                private renderer: Renderer2,
                public model: NavigationModel) {;
    }

    ngAfterViewInit(): void {
        this.locationDrawer.SetState(DrawerState.Hidden);
    }

    ionViewDidEnter() {
        if (this.isSubscripted === false){
            CanvasTouchHelper.RegisterPinch(this.renderer, this.canvasWrapper);
            this.isSubscripted = true;
            this.route.queryParams.subscribe(async params => {
                // discover requested location
                if (params != null && params.location != null && params.location.length > 0) {
                    this.searchInput.setDiscoverySearchbarValue(params.location);
                    await this.onDiscovery(params.location);
                    return;
                }

                // launch requested navigation
                if (params.start != null && params.start.length > 0 &&
                    params.destination != null && params.destination.length > 0) {
                    await this.showNavigation(params.start, params.destination);
                    return;
                }

                if (params.building != null && params.building.length > 0) {
                    const building = await this.dataService.get_building(params.building).toPromise()
                    if (building !== null) {
                        await this.mapView.showFloor(
                            building.Floors?.includes('EG') ? 'EG' : building.Floors[0],
                            building.Name);
                        return;
                    }
                }

                await this.showDiscoveryMode();
            });
        }
    }

    async ngOnInit() {
        if (this.model.recentSearches.length === 0) {
            const searches = await SearchResultProvider.readRecentSearch(this.storage);
            if (searches !== null) {
                this.model.recentSearches = searches;
                console.log(this.model.recentSearches);
            }
        }

        if (this.model.availableCampus.length === 0)
        {
            this.model.availableCampus = await this.dataService.get_campus().toPromise()
        }

        for (const campus of this.model.availableCampus) {
            this.availableCampus.push(new CampusViewModel(campus))
        }
    }

    public async onDiscovery(searchInput: string) {
        this.model.errorMessage = '';
        this.progressIsVisible = true;
        try {
            const location = await this.mapView.showDiscoveryLocation(searchInput);
            SearchResultProvider.addRecentSearch(searchInput, this.model, this.storage);
            this.scrollToCoordinate(location.PathNode.Coordinate.X, location.PathNode.Coordinate.Y);

            await this.showLocationDrawer(location);
            this.availableFloorsBtnIsVisible = true;
        } catch (ex) {
            this.handleInputError(ex, searchInput);
        } finally {
            this.progressIsVisible = false;
        }
    }

    public async onRoute(routeInput: string[]) {
        this.errorMessage = '';
        this.progressIsVisible = true;
        try {
            const startLocation = await this.dataService.get_location(routeInput[0]).toPromise<ILocation>();
            const route = await this.dataService.get_route(routeInput[0], routeInput[1]).toPromise();
            await this.mapView.showRoute(route, startLocation);
            this.scrollToCoordinate(startLocation.PathNode.Coordinate.X, startLocation.PathNode.Coordinate.Y);
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

    public onDrawerStateChange(state:DrawerState) {
        // in case the view is not initialized
        if (this.drawerContent === undefined) {
            return;
        }

        this.drawerContent.scrollY = state === DrawerState.Top;
    }

    public async showLocationDrawer(location:ILocation) {
        await this.locationDrawer.SetState(DrawerState.Hidden);
        this.model.selectedLocation = location;
        await this.searchDrawer.SetState(DrawerState.Hidden);
        await this.locationDrawer.SetState(DrawerState.Docked);
    }

    public async onCloseLocationDrawer(event:any) {
        await this.locationDrawer.SetState(DrawerState.Hidden);
        await this.searchDrawer.SetState(DrawerState.Docked);
    }

    async presentAvailableFloorModal() {
        let floors = new Array<string>();

        // STDG-138 discovery mode ... get all floor of all displayed buildings
        let buildings = await this.dataService.get_buildings().toPromise<IBuilding[]>();
        buildings = buildings.filter((n, i) => buildings.indexOf(n) === i);

        for (const building of buildings) {
            const buildingData = await this.dataService.get_building(building.Name).toPromise<IBuilding>();
            floors = floors.concat(buildingData.Floors);
        }

        // distinct array
        floors = floors.filter((n, i) => floors.indexOf(n) === i);

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
            this.showAnotherFloorOfCurrentBuilding(data.data, this.mapView.CurrentBuilding)
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
        } else {
            this.errorMessage = (ex as Error).message;
        }

        // TODO Remove this code when discovery mode is finished
        if (this.mapView != null) {
            this.mapView.clearMapCanvas();
        }
    }

    async navigationBtnClick() {
        if (this.model.selectedLocation != null) {
            // STDG 178 KV.001 wird als default start eingefügt
            await this.showNavigation('KV.001', this.model.selectedLocation.Name);
        }
    }

    private async showNavigation(start: string, destination: string) {
        this.searchInput.showRouteSearchBar();
        this.searchInput.setDiscoverySearchbarValue(destination);
        this.searchInput.setStartSearchbarValue(start);
        await this.onCloseLocationDrawer(null);
        await this.onRoute([start, destination])
    }

    private async showDiscoveryMode() {
        if (this.mapView.CurrentRoute == null && this.mapView.CurrentBuilding == null) {
            // STDG-138 load base map
            await this.mapView.showDiscoveryMap('', 'EG')
            this.availableFloorsBtnIsVisible = true;

            // Coordinates of KA.013
            this.scrollToCoordinate(310, 550);
        }
    }

    private scrollToCoordinate(xCoordinate: number, yCoordinate:number) {
        const availableSize = {width: window.innerWidth, height: window.innerHeight};

        CanvasTouchHelper.transistion(
            { x: CanvasTouchHelper.currentZoom.x - xCoordinate,
                y: CanvasTouchHelper.currentZoom.y - yCoordinate},
            this.canvasWrapper, this.renderer, false);
    }

    public async recentSearchClick(locationStr:string) {
        await this.router.navigate(['tabs/navigation/detail'], { queryParams: { location: locationStr } });
    }

    public async presentMapPage() {
        await this.router.navigate(['tabs/navigation/']);
    }

    public async onFloorChangeByFloorButton(floorAndBuildingInput: object) {
        // @ts-ignore
        await this.showAnotherFloorOfCurrentBuilding(floorAndBuildingInput.floor, floorAndBuildingInput.building);
    }

    private async showAnotherFloorOfCurrentBuilding(floor: string, building: string) {
        this.progressIsVisible = true;
        await this.mapView.showFloor(floor, building);
        this.progressIsVisible = false;
    }
}
