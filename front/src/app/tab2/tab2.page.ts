import {AfterViewInit, Component} from '@angular/core';
import {MoodleService} from '../services/moodle.service';
import {MoodleToken} from '../moodle-objects-if';

@Component({
  selector: 'app-tab2',
  templateUrl: 'tab2.page.html',
  styleUrls: ['tab2.page.scss']
})
export class Tab2Page implements AfterViewInit{
  private token: MoodleToken;

  constructor(private moodleService: MoodleService) {
  }

  async ngAfterViewInit() {
    this.token = await this.moodleService.getLoginToken('admin', 'administrator').toPromise();
  }
}
