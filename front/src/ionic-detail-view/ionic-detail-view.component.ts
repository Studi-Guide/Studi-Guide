import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-ionic-detail-view',
  templateUrl: './ionic-detail-view.component.html',
  styleUrls: ['./ionic-detail-view.component.scss'],
})
export class IonicDetailViewComponent implements OnInit {

  @Input() PageTitle:string;

  constructor() { }

  ngOnInit() {}

}
