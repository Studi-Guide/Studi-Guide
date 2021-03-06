import {AfterViewInit, Component, ViewChild} from '@angular/core';
import {MoodleService} from '../services/moodle.service';
import {Event, MoodleToken} from '../moodle-objects-if';
import {LoadingController} from '@ionic/angular';
import {Router} from '@angular/router';
import {LoginComponent} from './login/login.component';


@Component({
  selector: 'app-schedule',
  templateUrl: 'schedule.page.html',
  styleUrls: ['schedule.page.scss']
})
export class SchedulePage implements AfterViewInit {

  constructor(
      private moodleService: MoodleService,
      public loadingController: LoadingController,
      private router: Router
  ) {}
  private token: MoodleToken;
  public calenderEvents: Event[] = [];
  public isMoodleUserSignedIn: boolean;

  @ViewChild(LoginComponent) login: LoginComponent;

  async ionViewWillEnter() {
    await this.login.checkMoodleLoginState();
  }

  public async onSignIn(isSignedIn: boolean) {
    isSignedIn ? this.isMoodleUserSignedIn = true : this.isMoodleUserSignedIn = false;
  }

  public async fetchMoodleData(moodleToken: MoodleToken) {
    const loading = await this.loadingController.create({
      message: 'Collecting moodle data...'
    });
    await loading.present();

    this.token = moodleToken;
    const calenderRequestData = await this.moodleService.getCalenderEventsWeek(moodleToken).toPromise();

    if (this.moodleService.containsEvents(calenderRequestData)) {
      // remove images => needs cookies
      this.calenderEvents = this.CleanupEvents(calenderRequestData.events);

      // add dummy location to KA.206
      for (const event of this.calenderEvents) {
        event.location = 'KA.206';
      }
      await loading.dismiss();
    } else {
      this.isMoodleUserSignedIn = false;
      this.login.isUserLoggedIn = false;
      await loading.dismiss();
    }
  }

  public async doRefreshEvents(event) {
    if (this.isMoodleUserSignedIn) {
      const calenderRequestData = await this.moodleService.getCalenderEventsWeek(this.token).toPromise();
      if (this.moodleService.containsEvents(calenderRequestData)) {
        this.calenderEvents = calenderRequestData.events;

        // add dummy location to KA.206
        for (const calenderEvent of this.calenderEvents) {
          calenderEvent.location = 'KA.206';
        }

        event.target.complete();
      } else {
        this.isMoodleUserSignedIn = false;
      }
    }
  }

  public async onLocationClick(locationName: string) {
    await this.router.navigate(['tabs/navigation/detail'], { queryParams: { location: locationName } });
  }

  ngAfterViewInit(): void {}

   private CleanupEvents(events: Event[]) {
     const imgRegex = new RegExp('<img[^>]*?>', 'g');
     for (const event of events) {
        if (imgRegex.test(event.description)) {
          for (const match of event.description.match(imgRegex)) {
            event.description = event.description.replace(match, '');
          }
        }
     }

     return events;
  }
}
