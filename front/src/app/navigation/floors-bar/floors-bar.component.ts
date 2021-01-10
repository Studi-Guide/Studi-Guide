import {Component, EventEmitter, Input, OnChanges, OnInit, Output, SimpleChanges, ViewChild} from '@angular/core';
import {DataService} from '../../services/data.service';

@Component({
  selector: 'floors-bar',
  templateUrl: './floors-bar.component.html',
  styleUrls: ['./floors-bar.component.scss'],
})
export class FloorsBarComponent implements OnInit, OnChanges {

  @Input() currentBuilding: string;
  @Input() currentFloor: string;
  @ViewChild('fab') fab;
  @Output() floorWithBuilding = new EventEmitter<object>();

  public availableFloors: string[];

  constructor(private dataService: DataService) {
    // this.hack();
  }

  ngOnInit() {}

  async ngOnChanges(changes: SimpleChanges) {
    if (changes.currentBuilding !== undefined && this.currentBuilding !== undefined) {
      const building = await this.dataService.get_building(this.currentBuilding).toPromise();
      this.availableFloors = building.Floors.reverse();
    }
    // this.hack();
  }

  public async emitAnotherFloorToShow(index: number) {
    this.floorWithBuilding.emit({
      floor: this.availableFloors[index],
      building: this.currentBuilding
    });
    this.currentFloor = this.availableFloors[index];
  }

  // private hack() {
  //   this.currentFloor = '1';
  //   this.currentBuilding = 'KA';
  //   this.availableFloors = ['0',
  //     '1',
  //     '2',
  //     '3',
  //     '4',
  //     '5',
  //     '6',
  //     '7',
  //     '8',
  //     '9'].reverse();
  // }
}
