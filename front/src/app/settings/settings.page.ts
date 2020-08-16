import { Storage } from '@ionic/storage';
import { Component } from '@angular/core';

@Component({
  selector: 'app-settings',
  templateUrl: 'settings.page.html',
  styleUrls: ['settings.page.scss']
})
export class SettingsPage {

  public isSignedIn:boolean;
  public moodleUserName:string;

  private readonly MOODLE_TOKEN = 'moodle_token';
  private readonly MOODLE_USER = 'moodle_user';

  constructor(
      private storage: Storage
  ) {}

  async ionViewWillEnter() {
    this.storage.ready().then(async () => {
      await this.isMoodleUserLoggedIn();
    });
  }

  public async logoutFromMoodle() {
    await this.storage.remove(this.MOODLE_USER).then(async value => {
      await this.storage.remove(this.MOODLE_TOKEN).then(async value => {
        this.setLoggedOutFromMoodle();
      });
    });
  }

  private async isMoodleUserLoggedIn() {
    await this.storage.get(this.MOODLE_TOKEN).then(async value => {
      if (value != null || value != undefined) {
        console.log(value.token);
        this.isSignedIn = true;
        await this.getMoodleUserName();
      } else {
        this.setLoggedOutFromMoodle();
      }
    });
  }

  private setLoggedOutFromMoodle() {
    this.isSignedIn = false;
    this.moodleUserName = 'Kein User eingelogged.'
  }

  private async getMoodleUserName() {
    await this.storage.get(this.MOODLE_USER).then(userName => {
      this.moodleUserName = userName;
    });
  }

  actionSheetOptions: any = {
    header: 'Moodle'
  };
}