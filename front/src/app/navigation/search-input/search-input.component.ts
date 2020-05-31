import {Component, Output, EventEmitter, OnInit, ViewChild} from '@angular/core';

@Component({
  selector: 'app-search-input',
  templateUrl: './search-input.component.html',
  styleUrls: ['./search-input.component.scss'],
})
export class SearchInputComponent implements OnInit {

  @Output() search = new EventEmitter<string>();
  @Output() route = new EventEmitter<string[]>();

  @ViewChild('discoverySearchbar') discoverySearchbar;
  @ViewChild('routeSearchBar') routeSearchBar;

  private searchBtnIsVisible = true;
  private routeInputIsVisible = false;
  private closeRouteBtnIsVisible = false;

  constructor() { }

  ngOnInit() {}

  private showRouteSearchBar() {
      this.routeInputIsVisible = true;
      const searchbars = document.querySelector('ion-item');
      searchbars.setAttribute('color', 'primary');
      this.searchBtnIsVisible = false;
      this.closeRouteBtnIsVisible = true;
  }

  private hideRouteSearchBar() {
    if (this.routeInputIsVisible) {
      this.routeInputIsVisible = false;
      const searchbars = document.querySelector('ion-item');
      searchbars.setAttribute('color', 'light-tint');
      this.searchBtnIsVisible = true;
      this.closeRouteBtnIsVisible = false;
    }
  }

  private routeBtnClick() {
    if (!this.routeInputIsVisible) {
      this.showRouteSearchBar();
    } else {
      this.emitRouteEvent();
    }
  }

  private emitSearchEvent() {
    this.search.emit(this.discoverySearchbar.value);
  }

  private emitRouteEvent() {
    const route:string[] = [this.discoverySearchbar.value, this.routeSearchBar.value];
    this.route.emit(route);
  }

}
