import {Component, Inject, LOCALE_ID, OnInit} from '@angular/core';
import {Router} from '@angular/router';

@Component({
  selector: 'detail-view-select-language',
  templateUrl: './detail-view-select-language.component.html',
  styleUrls: ['./detail-view-select-language.component.scss'],
})
export class DetailViewSelectLanguageComponent implements OnInit {

  public readonly languages = [
      {Language: 'English',  Identifier: 'en-US'},
      {Language: 'German', Identifier: 'de'}
  ];

  constructor(
      private router:Router,
      @Inject(LOCALE_ID) private locale:string
  ) { }

  ngOnInit() {
    console.log(this.router.url, this.locale);
  }

  public SelectLanguageClick(identifier:string) {
    console.log(identifier);
    if (identifier !== this.locale) {
      this.router.navigate(['/' + identifier + this.router.url]);
    } else {
      console.log('locale', this.locale, 'not found in current URL');
    }
  }

}
