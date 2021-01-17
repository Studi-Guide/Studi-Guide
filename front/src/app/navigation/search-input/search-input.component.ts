import {Component, Output, EventEmitter, ViewChild, OnInit} from '@angular/core';
import {DataService} from '../../services/data.service';
import {NavigationModel} from '../navigationModel';

interface IListObject {
  Name: string;
  Description: string;
}

@Component({
  selector: 'app-search-input',
  templateUrl: './search-input.component.html',
  styleUrls: ['./search-input.component.scss'],
})
export class SearchInputComponent implements OnInit {

  @Output() discovery = new EventEmitter<string>();
  @Output() route = new EventEmitter<string[]>();
  @Output() searchBarFocus = new EventEmitter<string>();

  @ViewChild('destinationSearchbar') destinationSearchbar: HTMLIonSearchbarElement;
  @ViewChild('startSearchBar') startSearchBar: HTMLIonSearchbarElement;

  public searchBtnIsVisible = true;
  public routeInputIsVisible = false;
  public closeRouteBtnIsVisible = false;
  public startSearchBarValue: string;
  public destinationSearchbarValue: string;

  public locations: IListObject[];

  constructor(private dataService: DataService,
              private model: NavigationModel) {}

  ngOnInit() {}

  public async onDestinationInputEvent(e: any) {
    const tmpLocations = await this.dataService.get_locations_all().toPromise();
    const query = (e.target as HTMLInputElement).value.toLowerCase();

    if (query.length < 2) {
      this.locations = [];
      return;
    }

    this.locations = [];

    this.filterIListObjects(this.model.recentSearches, query, 2);
    this.filterIListObjects(tmpLocations, query, 5);

  }

  private filterIListObjects(objects: IListObject[], query: string, maxLength: number) {
    for (const l of objects) {
      let shouldShow = l.Name.toLowerCase().indexOf(query) > -1 || l.Description.toLowerCase().indexOf(query) > -1;
      for (const o of this.locations) {
        if (o.Name === l.Name) {
          shouldShow = false;
          break;
        }
      }
      if (shouldShow) {
        this.locations.push(l);
      }
      if (this.locations.length > maxLength) {
        return;
      }
    }
  }

  public showRouteSearchBar() {
      this.routeInputIsVisible = true;
      // const searchbars = document.querySelector('ion-item');
      // searchbars.setAttribute('color', 'primary');
      this.searchBtnIsVisible = false;
      this.closeRouteBtnIsVisible = true;
  }

  public hideRouteSearchBar() {
    if (this.routeInputIsVisible) {
      this.routeInputIsVisible = false;
      // const searchbars = document.querySelector('ion-item');
      // searchbars.setAttribute('color', 'light-tint');
      this.searchBtnIsVisible = true;
      this.closeRouteBtnIsVisible = false;
    }
  }

  public routeBtnClick() {
    if (!this.routeInputIsVisible) {
      this.showRouteSearchBar();
    } else {
      this.emitRouteEvent();
    }
  }

  public emitDiscoveryEvent() {
    let isInputEmpty = this.destinationSearchbar.value === '' || this.destinationSearchbar.value === undefined;
    isInputEmpty = isInputEmpty || this.destinationSearchbar.value === null;
    if (!isInputEmpty) {
      // Workaround for https://github.com/ionic-team/ionic-v3/issues/217
      const activeElement = document.activeElement as HTMLElement;
      if (activeElement && activeElement.blur){
        activeElement.blur();
      }

      this.discovery.emit(this.destinationSearchbar.value);
    }
  }

  public emitRouteEvent() {
    let isStartEmpty = this.destinationSearchbar.value === '' || this.destinationSearchbar.value === undefined;
    isStartEmpty = isStartEmpty || this.destinationSearchbar.value === null;
    let isDestinationEmpty = this.startSearchBar.value === '' || this.startSearchBar.value === undefined;
    isDestinationEmpty = isDestinationEmpty || this.startSearchBar.value === null;
    if (!isStartEmpty && !isDestinationEmpty) {
      const route: string[] = [this.startSearchBar.value, this.destinationSearchbar.value];

      // Workaround for https://github.com/ionic-team/ionic-v3/issues/217
      const activeElement = document.activeElement as HTMLElement;
      if (activeElement && activeElement.blur){
        activeElement.blur();
      }

      this.route.emit(route);
    }
  }

  public clearDestinationInput() {
    this.destinationSearchbarValue = '';
    this.locations = [];
  }

  public clearStartInput() {
    this.startSearchBarValue = '';
  }

  public setDiscoverySearchbarValue(value: string) {
    this.destinationSearchbarValue = value;
  }

  public setStartSearchbarValue(value: string) {
    this.startSearchBarValue = value;
  }

  onSearchBarHasFocus(searchBar: string) {
    this.searchBarFocus.emit(searchBar);
  }

  onKey(e: KeyboardEvent, inputElement: string) {
    if (e.key === 'Enter') {
      if (inputElement === 'destinationSearchbar') {
          this.emitDiscoveryEvent();
      }

      if (inputElement === 'startSearchBar') {
          this.emitRouteEvent();
      }
    }
  }
}
