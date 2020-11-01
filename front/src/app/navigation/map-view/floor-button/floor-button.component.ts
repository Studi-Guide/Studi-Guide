import {Component, OnInit} from '@angular/core';

@Component({
  selector: 'app-floor-button',
  templateUrl: './floor-button.component.html',
  styleUrls: ['./floor-button.component.scss'],
})
export class FloorButtonComponent implements OnInit {

  public availableFloors:string[];

  constructor() { }

  ngOnInit() {}

}
