import {AfterViewInit, Component} from '@angular/core';
import {MoodleService} from '../services/moodle.service';
import {Event, MoodleToken} from '../moodle-objects-if';
import {LoadingController} from '@ionic/angular';
import {Router} from '@angular/router';

@Component({
  selector: 'app-schedule',
  templateUrl: 'schedule.page.html',
  styleUrls: ['schedule.page.scss']
})
export class SchedulePage implements AfterViewInit{
  private token: MoodleToken;
  public calenderEvents: Event[] = [];

  constructor(private moodleService: MoodleService, public loadingController: LoadingController, private router: Router) {
  }

  async ngAfterViewInit() {
    const loading = await this.loadingController.create({
      message: 'Collecting moodle data...'
    });
    const task = loading.present();
    this.token = await this.moodleService.getLoginToken('admin', 'administrator').toPromise();
    const calenderRequestData = await this.moodleService.getCalenderEventsWeek(this.token).toPromise();
    this.calenderEvents = calenderRequestData.events;

    // add dummy location to KA.206
    for(const event of this.calenderEvents) {
      event.location = 'KA.206'
    }

    await loading.dismiss();
  }

    async doRefreshEvents(event) {
      const calenderRequestData = await this.moodleService.getCalenderEventsWeek(this.token).toPromise();
      this.calenderEvents = calenderRequestData.events;

      // add dummy location to KA.206
      for(const calenderEvent of this.calenderEvents) {
        calenderEvent.location = 'KA.206'
      }

      event.target.complete();
    }

  async onLocationClick(locationName: string) {
    await this.router.navigate(['tabs/navigation'], { queryParams: { location: locationName } });
  }
}
