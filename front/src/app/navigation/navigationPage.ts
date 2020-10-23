import {IBuilding, ICampus, ILocation, IPathNode} from '../building-objects-if';
import {AfterViewInit, Component, ElementRef, OnInit, Renderer2, ViewChild} from '@angular/core';
import {IonContent, ModalController} from '@ionic/angular';
import {DataService} from '../services/data.service';
import {AvailableFloorsPage} from '../available-floors/available-floors.page';
import {MapViewComponent} from './map-view/map-view.component';
import {HttpErrorResponse} from '@angular/common/http';
import {ActivatedRoute} from '@angular/router';
import {SearchInputComponent} from './search-input/search-input.component';
import {DrawerState} from '../../ionic-bottom-drawer/drawer-state';
import {IonicBottomDrawerComponent} from '../../ionic-bottom-drawer/ionic-bottom-drawer.component';
import { Storage } from '@ionic/storage';
import {Router} from '@angular/router';
import {CanvasTouchHelper} from '../services/CanvasTouchHelper';
import {CampusViewModel} from './campusViewModel';

@Component({
    selector: 'app-navigation',
    templateUrl: 'navigation.page.html',
    styleUrls: ['navigation.page.scss']
})

export class NavigationPage implements OnInit{

    @ViewChild(MapViewComponent) mapView: MapViewComponent;
    @ViewChild(SearchInputComponent) searchInput: SearchInputComponent;
    @ViewChild('drawerContent') drawerContent : IonContent;
    @ViewChild('searchDrawer') searchDrawer : IonicBottomDrawerComponent;
    @ViewChild('locationDrawer') locationDrawer : IonicBottomDrawerComponent;
    @ViewChild('canvasWrapper', {read: ElementRef}) private canvasWrapper: ElementRef;

    public progressIsVisible = false;
    public availableFloorsBtnIsVisible = false;
    public errorMessage: string;
    public selectedLocation:ILocation = {
        Building: '',
        Description: '',
        Floor: '',
        Id: 0,
        Name: '',
        PathNode: {
            Coordinate: {X: 0, Y: 0, Z: 0},
            Id: 0
            },
        Tags: []
    };

    private recentSearchesKey = 'searches';
    public recentSearches : string[] = [];

    public availableCampus: CampusViewModel[] = [];

    constructor(private dataService: DataService,
                private modalCtrl: ModalController,
                private  route: ActivatedRoute,
                private router: Router,
                private storage: Storage,
                private renderer: Renderer2) {
    }

    ionViewDidEnter() {
        CanvasTouchHelper.RegisterPinch(this.renderer, this.canvasWrapper);
        this.route.queryParams.subscribe(async params =>
        {
                // discover requested location
                if (params != null && params.location != null && params.location.length > 0) {
                    this.searchInput.setDiscoverySearchbarValue(params.location);
                    await this.onDiscovery(params.location);
                    return;
                }

                // launch requested navigation
                if (params.start != null && params.start.length > 0 &&
                        params.destination != null && params.destination.length >0) {
                    await this.showNavigation(params.start, params.destination);
                    return;
                }

                await this.showDiscoveryMode();
        });
    }

    async ngOnInit() {
        const searches = JSON.parse(await this.storage.get(this.recentSearchesKey));
        if (searches !== null) {
            this.recentSearches = searches[0];
            console.log(this.recentSearches);
        }

        const campusModels = await this.dataService.get_campus().toPromise()
        for (const campus of campusModels) {
            this.availableCampus.push(new CampusViewModel(campus))
        }
    }

    public async onDiscovery(searchInput: string) {
        this.errorMessage = '';
        this.progressIsVisible = true;
        try {
            const location = await this.mapView.showDiscoveryLocation(searchInput);
            this.addRecentSearch(searchInput);
            this.scrollToCoordinate(location.PathNode.Coordinate.X, location.PathNode.Coordinate.Y);

            // Wenn man hier ein await einfügt spackt der drawer komplett. Keine ahnung wieso
            this.showLocationDrawer(location);
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

        if (state === DrawerState.Top) {
            this.drawerContent.scrollY = true;
        } else {
            this.drawerContent.scrollY = false;
        }
    }

    public async showLocationDrawer(location:ILocation) {
        await this.locationDrawer.SetState(DrawerState.Hidden);
        this.selectedLocation = location;
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
            this.progressIsVisible = true;
            await this.mapView.showFloor(data.data, this.mapView.CurrentBuilding);
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
        } else {
            this.errorMessage = (ex as Error).message;
        }

        // TODO Remove this code when discovery mode is finished
        if (this.mapView != null) {
            this.mapView.clearMapCanvas();
        }
    }

    async navigationBtnClick() {
        if (this.selectedLocation != null) {
            // STDG 178 KV.001 wird als default start eingefügt
            await this.showNavigation('KV.001', this.selectedLocation.Name);
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
        CanvasTouchHelper.transistion({ x:CanvasTouchHelper.currentZoom.x - xCoordinate, y: CanvasTouchHelper.currentZoom.y - yCoordinate},
            this.canvasWrapper, this.renderer, false);
    }

    private addRecentSearch(location:string) {
        if (this.recentSearches.includes(location)) {
            this.recentSearches.splice(this.recentSearches.indexOf(location), 1);
        }

        this.recentSearches.unshift(location);

        if (this.recentSearches.length > 3) {
            this.recentSearches.pop();
        }

        this.storage.set(this.recentSearchesKey, JSON.stringify([this.recentSearches]));
        console.log(this.recentSearches);
    }

    public async recentSearchClick(locationStr:string) {
        await this.router.navigate(['tabs/navigation'], { queryParams: { location: locationStr } });
    }

    public async presentMapPage() {
        await this.router.navigate(['tabs/navigation/map']);
    }
}
