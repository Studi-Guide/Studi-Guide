import {BuildingData, Location} from '../building-objects-if';
import {AfterViewInit, Component, OnInit, ViewChild} from '@angular/core';
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

@Component({
    selector: 'app-navigation',
    templateUrl: 'navigation.page.html',
    styleUrls: ['navigation.page.scss']
})

export class NavigationPage implements  AfterViewInit, OnInit{

    @ViewChild(MapViewComponent) mapView: MapViewComponent;
    @ViewChild(SearchInputComponent) searchInput: SearchInputComponent;
    @ViewChild('drawerContent') drawerContent : IonContent;
    @ViewChild('searchDrawer') searchDrawer : IonicBottomDrawerComponent;
    @ViewChild('locationDrawer') locationDrawer : IonicBottomDrawerComponent;

    public progressIsVisible = false;
    public availableFloorsBtnIsVisible = false;
    public errorMessage: string;
    public clickedLocation:Location = {
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

    constructor(private dataService: DataService,
                private modalCtrl: ModalController,
                private  route: ActivatedRoute,
                private storage: Storage) {
    }

    async ngAfterViewInit(): Promise<void> {
        this.route.params.subscribe(async params =>
        {
            if (params != null && params.location != null && params.location.length > 0) {
                this.searchInput.setDiscoverySearchbarValue(params.location);
                await this.onDiscovery(params.location);
            }
            else {
                if (this.mapView.CurrentRoute == null && this.mapView.CurrentBuilding == null) {
                    // STDG-138 load base map
                    await this.mapView.showDiscoveryMap('', 'EG')
                    this.availableFloorsBtnIsVisible = true;

                    // Scroll to mid
                    const div = document.getElementById('canvas-wrapper');

                    // Coordinates of KA.013
                    div.scrollBy(345 - 50,600 - 125);
                }
            }
        });
    }

    async ngOnInit() {
        const searches = await this.storage.get(this.recentSearchesKey);
        if (searches !== undefined || searches !== ' ') {
            this.recentSearches = JSON.parse(searches)[0];
            console.log(this.recentSearches);
        }
    }

    public async onDiscovery(searchInput: string) {
        this.errorMessage = '';
        this.progressIsVisible = true;
        try {
            await this.mapView.showDiscoveryLocation(searchInput);
            this.addRecentSearch(searchInput);
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

    public async onMapViewLocationClick(location:Location) {
        await this.locationDrawer.SetState(DrawerState.Hidden);
        this.clickedLocation = location;
        this.searchDrawer.SetState(DrawerState.Hidden);
        this.locationDrawer.SetState(DrawerState.Docked);
    }

    public async onCloseLocationDrawer(event:any) {
        await this.locationDrawer.SetState(DrawerState.Hidden);
        this.searchDrawer.SetState(DrawerState.Docked);
    }

    async presentAvailableFloorModal() {
        let floors = new Array<string>();

        if (this.mapView.CurrentRoute == null) {
            if (this.mapView.CurrentBuilding != null) {
                const building = await this.dataService.get_building(this.mapView.CurrentBuilding).toPromise<BuildingData>();
                floors = floors.concat(building.Floors);
            }
            else {
                // STDG-138 discovery mode ... get all floor of all displayed buildings
                let buildings = this.mapView.floorMapRenderer.objectsToBeVisualized.map(x => x.Building);
                buildings = buildings.filter((n, i) => buildings.indexOf(n) === i);

                for (const building of buildings) {
                    const buildingData = await this.dataService.get_building(building).toPromise<BuildingData>();
                    floors = floors.concat(buildingData.Floors);
                }
            }
        } else {
            // get all floors from all buildings on the route
            for (const routeSection of this.mapView.CurrentRoute.RouteSections) {
                const building = await this.dataService.get_building(routeSection.Building).toPromise<BuildingData>();
                floors = floors.concat(building.Floors);
            }
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
}
