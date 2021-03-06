import { Storage } from '@ionic/storage';
import {Component, EventEmitter, OnInit, Output, ViewChild} from '@angular/core';
import {MoodleService} from '../../services/moodle.service';
import { MoodleToken } from 'src/app/moodle-objects-if';
import {FingerprintAIO} from '@ionic-native/fingerprint-aio/ngx';
import { Platform } from '@ionic/angular';
import {SettingsModel} from '../../settings/settings.model';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss'],
})
export class LoginComponent implements OnInit {
  @Output() isSignedIn = new EventEmitter<boolean>();
  @Output() moodleToken = new EventEmitter<MoodleToken>();

  public isUserLoggedIn = true;
  public token: MoodleToken;
  public invalidCredentialsMessage: string;

  private readonly MOODLE_TOKEN = 'moodle_token';
  private readonly MOODLE_USER = 'moodle_user';
  public userInput = 'admin';
  public userPassword = 'administrator';

  constructor(
      private storage: Storage,
      private moodleService: MoodleService,
      private faio: FingerprintAIO,
      private platform: Platform,
      public  settings: SettingsModel
  ) {}

  ngOnInit() {}

  public async checkMoodleLoginState() {
    await this.storage.ready();

    let autoLogin = false;
    if (this.platform.is('hybrid') && this.userInput.length > 0 &&  this.userPassword.length > 0) {
      // Shot face id on hybrid
      try {
        await this.faio.show({});
        autoLogin = true;
        console.log('Face ID result' + autoLogin);
      }
      catch (e) {
        console.log(e);
      }
    }

    if (autoLogin) {
      await this.fetchAndPersistMoodleToken();
    }

    await this.getPersistedToken();
    this.token == null ? this.isUserLoggedIn = false : this.isUserLoggedIn = true;
    this.isSignedIn.emit(this.isUserLoggedIn);
    if (this.isUserLoggedIn) {
      this.moodleToken.emit(this.token);
    }
  }

  public async fetchAndPersistMoodleToken() {
    const userName = this.userInput;
    const password = this.userPassword;

    const tokenToPersist = await this.moodleService.getLoginToken(userName, password).toPromise();

    if (this.moodleService.containsToken(tokenToPersist)) {
      this.isUserLoggedIn = true;
      this.moodleToken.emit(tokenToPersist);
      this.isSignedIn.emit(this.isUserLoggedIn);
      await this.storage.set(this.MOODLE_USER, userName);
      await this.storage.set(this.MOODLE_TOKEN, tokenToPersist);

      // clear fields
      this.userInput = '';
      this.userPassword = '';
    } else {
      // if login fails moodle response contains: "errorcode":"invalidlogin"
      this.isUserLoggedIn = false;
      this.isSignedIn.emit(this.isUserLoggedIn);
      this.invalidCredentialsMessage = 'Invalid credentials';
    }
  }

  public clearInvalidCredentialsMsg() {
    this.invalidCredentialsMessage = '';
  }

  private async getPersistedToken() {
    await this.storage.get(this.MOODLE_TOKEN).then(value => {
      this.token = value;
    });
  }
}
