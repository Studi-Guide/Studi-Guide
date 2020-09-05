import { Storage } from '@ionic/storage';
import { Component } from '@angular/core';
import { MoodleService } from '../services/moodle.service';
import { MoodleToken } from '../moodle-objects-if';

@Component({
  selector: 'app-settings',
  templateUrl: 'settings.page.html',
  styleUrls: ['settings.page.scss']
})
export class SettingsPage {

  constructor(
      private storage: Storage,
      private moodleService: MoodleService
  ) {}

  public isSignedIn:boolean;
  public moodleUserName:string;
  public persistedMoodleToken:MoodleToken;

  private readonly MOODLE_TOKEN = 'moodle_token';
  private readonly MOODLE_USER = 'moodle_user';

  actionSheetOptions: any = {
    header: 'Moodle'
  };

  async ionViewWillEnter() {
    this.storage.ready().then(async () => {
      if (await this.isMoodleTokenPersisted()) {
        const data = await this.moodleService.getCalenderEventsWeek(this.persistedMoodleToken).toPromise();
        if (this.moodleService.containsEvents(data)) {
          this.isSignedIn = true;
          await this.getMoodleUserName();
          return;
        }
      }
      this.setLoggedOutFromMoodle();
    });
  }

  public async logoutFromMoodle() {
    await this.storage.remove(this.MOODLE_USER).then(() => {
      this.storage.remove(this.MOODLE_TOKEN).then( () => {
        this.setLoggedOutFromMoodle();
      });
    });
  }

  private async isMoodleTokenPersisted():Promise<boolean> {
    return await this.storage.get(this.MOODLE_TOKEN).then( (value) => {
      this.persistedMoodleToken = value;
      return this.persistedMoodleToken != null;
    });
  }

  private setLoggedOutFromMoodle() {
    this.isSignedIn = false;
    this.moodleUserName = 'No user signed in.'
  }

  private async getMoodleUserName() {
    await this.storage.get(this.MOODLE_USER).then(userName => {
      this.moodleUserName = userName;
    });
  }
}