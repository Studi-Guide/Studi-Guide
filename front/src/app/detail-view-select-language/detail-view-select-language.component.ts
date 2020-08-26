import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'detail-view-select-language',
  templateUrl: './detail-view-select-language.component.html',
  styleUrls: ['./detail-view-select-language.component.scss'],
})
export class DetailViewSelectLanguageComponent implements OnInit {

  @Input() PageTitle:string;

  constructor() { }

  ngOnInit() {}

}
