import {Component, Output, EventEmitter, OnInit, ViewChild} from '@angular/core';

@Component({
  selector: 'app-search-input',
  templateUrl: './search-input.component.html',
  styleUrls: ['./search-input.component.scss'],
})
export class SearchInputComponent implements OnInit {

  @Output() discovery = new EventEmitter<string>();
  @Output() route = new EventEmitter<string[]>();

  @ViewChild('discoverySearchbar') discoverySearchbar;
  @ViewChild('routeSearchBar') routeSearchBar;

  public searchBtnIsVisible = true;
  public routeInputIsVisible = false;
  public closeRouteBtnIsVisible = false;
  private routeSearchBarValue: string;
  private discoverySearchbarValue: string;

  constructor() { }

  ngOnInit() {}

  public showRouteSearchBar() {
      this.routeInputIsVisible = true;
      const searchbars = document.querySelector('ion-item');
      searchbars.setAttribute('color', 'primary');
      this.searchBtnIsVisible = false;
      this.closeRouteBtnIsVisible = true;
  }

  public hideRouteSearchBar() {
    if (this.routeInputIsVisible) {
      this.routeInputIsVisible = false;
      const searchbars = document.querySelector('ion-item');
      searchbars.setAttribute('color', 'light-tint');
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
    let isInputEmpty = this.discoverySearchbar.value === '' || this.discoverySearchbar.value === undefined;
    isInputEmpty = isInputEmpty || this.discoverySearchbar.value === null;
    if (!isInputEmpty) {
      this.discovery.emit(this.discoverySearchbar.value);
    }
  }

  public emitRouteEvent() {
    let isStartEmpty = this.discoverySearchbar.value === '' || this.discoverySearchbar.value === undefined;
    isStartEmpty = isStartEmpty || this.discoverySearchbar.value === null;
    let isDestinationEmpty = this.routeSearchBar.value === '' || this.routeSearchBar.value === undefined;
    isDestinationEmpty = isDestinationEmpty || this.routeSearchBar.value === null;
    if (!isStartEmpty && !isDestinationEmpty) {
      const route:string[] = [this.discoverySearchbar.value, this.routeSearchBar.value];
      this.route.emit(route);
    }
  }

  public setDiscoverySearchbarValue(value:string) {
    this.discoverySearchbarValue = value;
  }

  public setStartSearchbarValue(value:string) {
    this.routeSearchBarValue = value;
  }
}
