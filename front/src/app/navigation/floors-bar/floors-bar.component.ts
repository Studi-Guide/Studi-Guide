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

  constructor(private dataService: DataService) {}

  ngOnInit() {}

  async ngOnChanges(changes: SimpleChanges) {
    if (changes.currentBuilding !== undefined && this.currentBuilding !== undefined) {
      const building = await this.dataService.get_building(this.currentBuilding).toPromise();
      const floors = await this.dataService.get_building_floor(building.Name).toPromise();
      this.availableFloors = floors.reverse();
    }
  }

  public async emitAnotherFloorToShow(index: number) {
    this.floorWithBuilding.emit({
      floor: this.availableFloors[index],
      building: this.currentBuilding
    });
    this.currentFloor = this.availableFloors[index];
  }
}
