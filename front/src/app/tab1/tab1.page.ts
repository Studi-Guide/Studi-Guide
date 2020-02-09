import { Component } from '@angular/core';
import { floor } from '../building-objects-if';
import { room } from '../building-objects-if';
import { corridor } from '../building-objects-if';
import { testDataRooms } from './building-data';

@Component({
  selector: 'app-tab1',
  templateUrl: 'tab1.page.html',
  styleUrls: ['tab1.page.scss']
})
export class Tab1Page {
  public mapIsVisible:boolean = true;
  public startRoom:room;
  public destinationRoom:room;
  public testRooms:room[] = testDataRooms;
  
  // todo: these values should be sent from backend
  public svgWidth:number = this.calcSvgWidth();
  public svgHeight:number = this.calcSvgHeight();
  
  public calcSvgWidth() {
    let sum:number = 0;
    this.testRooms.forEach(room => {
      if ( room.x + room.width > sum ) {
        sum = room.x + room.width;
      }
    });
    return sum;
  }

  public calcSvgHeight() {
    let sum:number = 0;
    this.testRooms.forEach(room => {
      if ( room.y + room.height > sum ) {
        sum = room.y + room.height;
      }
    });
    return sum;
  }

  constructor() {}

}