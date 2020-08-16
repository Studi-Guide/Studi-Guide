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
  private token: MoodleToken;
  public calenderEvents: Event[] = [];
  public isMoodleUserSignedIn: boolean;

  @ViewChild(LoginComponent) login:LoginComponent;

  constructor(
      private moodleService: MoodleService,
      public loadingController: LoadingController,
      private router: Router
  ) {}

  async ionViewWillEnter() {
    this.login.checkMoodleLoginState();
  }

  public async onSignIn(isSignedIn: boolean) {
    console.log('schedule::onSignIn - isSignedIn: '+isSignedIn);
    isSignedIn ? this.isMoodleUserSignedIn = true : this.isMoodleUserSignedIn = false;
  }

  public async fetchMoodleData(moodleToken: MoodleToken) {
    // TODO check response if token is expired
    //  => this.isMoodleUserSignedIn = false; this.storage.remove('moodleToken');
    console.log('schedule::fetchMoodleData');
    const loading = await this.loadingController.create({
      message: 'Collecting moodle data...'
    });
    const task = loading.present();

    this.token = moodleToken;
    const calenderRequestData = await this.moodleService.getCalenderEventsWeek(moodleToken).toPromise();

    if (calenderRequestData.events != undefined || calenderRequestData.events != null) {
      this.calenderEvents = calenderRequestData.events;

      // add dummy location to KA.206
      for (const event of this.calenderEvents) {
        event.location = 'KA.206'
      }
      await loading.dismiss();
    } else {
      // TODO
      this.isMoodleUserSignedIn = false;
      loading.dismiss();
    }
  }

  async doRefreshEvents(event) {
    if (this.isMoodleUserSignedIn) {
      const calenderRequestData = await this.moodleService.getCalenderEventsWeek(this.token).toPromise();
      if (calenderRequestData.events !== undefined || calenderRequestData.events !== null) {
        this.calenderEvents = calenderRequestData.events;

        // add dummy location to KA.206
        for (const calenderEvent of this.calenderEvents) {
          calenderEvent.location = 'KA.206'
        }

        event.target.complete();
      } else {
        this.isMoodleUserSignedIn = false;
      }
    }
  }

  async onLocationClick(location: string) {
    await this.router.navigateByUrl('tabs/navigation/search/' + location)
  }

  ngAfterViewInit(): void {}
}
