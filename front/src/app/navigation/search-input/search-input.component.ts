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

  constructor() { }

  ngOnInit() {}

  public showRouteSearchBar() {
      this.routeInputIsVisible = true;
      const searchbars = document.querySelector('ion-item');
      searchbars.setAttribute('color', 'primary');
      this.searchBtnIsVisible = false;
      this.closeRouteBtnIsVisible = true;
      document.getElementById('map-wrapper').setAttribute('style', 'height: calc(100% - 121px);');
  }

  public hideRouteSearchBar() {
    if (this.routeInputIsVisible) {
      this.routeInputIsVisible = false;
      const searchbars = document.querySelector('ion-item');
      searchbars.setAttribute('color', 'light-tint');
      this.searchBtnIsVisible = true;
      this.closeRouteBtnIsVisible = false;
      document.getElementById('map-wrapper').setAttribute('style', '');
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
    this.discovery.emit(this.discoverySearchbar.value);
  }

  public emitRouteEvent() {
    const route:string[] = [this.discoverySearchbar.value, this.routeSearchBar.value];
    this.route.emit(route);
  }

}
