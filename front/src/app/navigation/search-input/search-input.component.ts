import {Component, Output, EventEmitter, OnInit, ViewChild} from '@angular/core';

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

  constructor() { }

  ngOnInit() {}

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
