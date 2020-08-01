import {AfterViewInit, Component} from '@angular/core';
import {MoodleService} from '../services/moodle.service';
import {Event, MoodleToken} from '../moodle-objects-if';
import {LoadingController} from '@ionic/angular';

@Component({
  selector: 'app-schedule',
  templateUrl: 'schedule.page.html',
  styleUrls: ['schedule.page.scss']
})
export class SchedulePage implements AfterViewInit {
  private token: MoodleToken;
  public calenderEvents: Event[] = [];

  constructor(private moodleService: MoodleService, public loadingController: LoadingController) {}

  public async onSignIn(isSignedIn: boolean) {
    if (isSignedIn) {
      await this.fetchMoodleData();
    }
  }

  private async fetchMoodleData() {
    const loading = await this.loadingController.create({
      message: 'Collecting moodle data...'
    });
    const task = loading.present();
    this.token = await this.moodleService.getLoginToken('admin', 'administrator').toPromise();
    const calenderRequestData = await this.moodleService.getCalenderEventsWeek(this.token).toPromise();
    this.calenderEvents = calenderRequestData.events;
    await loading.dismiss();
  }

  async doRefreshEvents(event) {
    const calenderRequestData = await this.moodleService.getCalenderEventsWeek(this.token).toPromise();
    this.calenderEvents = calenderRequestData.events;
    event.target.complete();
  }

  ngAfterViewInit(): void {}
}
