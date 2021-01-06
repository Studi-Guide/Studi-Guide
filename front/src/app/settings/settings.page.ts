import { Storage } from '@ionic/storage';
import {AfterViewInit, Component, ViewChild} from '@angular/core';
import { MoodleService } from '../services/moodle.service';
import { MoodleToken } from '../moodle-objects-if';
import {SettingsModel} from './settings.model';
import {IonToggle} from '@ionic/angular';

@Component({
  selector: 'app-settings',
  templateUrl: 'settings.page.html',
  styleUrls: ['settings.page.scss']
})
export class SettingsPage implements AfterViewInit {

  constructor(
      private storage: Storage,
      private moodleService: MoodleService,
      public settingsModel: SettingsModel
  ) {}

  public isSignedIn: boolean;
  public moodleUserName: string;
  public persistedMoodleToken: MoodleToken;

  private readonly MOODLE_TOKEN = 'moodle_token';
  private readonly MOODLE_USER = 'moodle_user';

  @ViewChild('DrawerDockingToggle') drawerDockingToggle: IonToggle;

  actionSheetOptions: any = {
    header: 'Moodle'
  };

  ngAfterViewInit() {
    this.drawerDockingToggle.checked = this.settingsModel.DrawerDocking;
  }

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

  private async isMoodleTokenPersisted(): Promise<boolean> {
    return await this.storage.get(this.MOODLE_TOKEN).then( (value) => {
      this.persistedMoodleToken = value;
      return this.persistedMoodleToken != null;
    });
  }

  private setLoggedOutFromMoodle() {
    this.isSignedIn = false;
    this.moodleUserName = 'No user signed in.';
  }

  private async getMoodleUserName() {
    await this.storage.get(this.MOODLE_USER).then(userName => {
      this.moodleUserName = userName;
    });
  }

  public onDrawerDockingToggleChange(event: any) {
    this.settingsModel.DrawerDocking = event.detail.checked;
  }

  public onDarkModeToggleChange(event: any) {
    this.settingsModel.DarkMode = event.detail.checked;
  }

  public onAutoDarkModeToggleChange(event: any) {
    this.settingsModel.AutoDarkMode = event.detail.checked;
  }
}
