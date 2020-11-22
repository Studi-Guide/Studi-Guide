import {Component, EventEmitter, Input, OnChanges, OnInit, Output, SimpleChanges, ViewChild} from '@angular/core';
import {DataService} from '../../services/data.service';

@Component({
  selector: 'app-floor-button',
  templateUrl: './floor-button.component.html',
  styleUrls: ['./floor-button.component.scss'],
})
export class FloorButtonComponent implements OnInit, OnChanges {

  @Input() currentBuilding:string;
  @Input() currentFloor:string;
  @ViewChild('fab') fab;
  @Output() floorWithBuilding = new EventEmitter<object>();

  public availableFloors:string[];

  constructor(private _dataService: DataService) {}

  ngOnInit() {}

  async ngOnChanges(changes: SimpleChanges) {
    if (changes.currentBuilding !== undefined && this.currentBuilding !== undefined) {
      const building = await this._dataService.get_building(this.currentBuilding).toPromise();
      this.availableFloors = building.Floors;
    }
  }

  public async emitAnotherFloorToShow(index:number) {
    this.floorWithBuilding.emit({
      floor: this.availableFloors[index],
      building: this.currentBuilding
    });
    this.currentFloor = this.availableFloors[index];
  }

}
