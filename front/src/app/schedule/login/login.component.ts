import { Storage } from '@ionic/storage';
import {Component, EventEmitter, OnInit, Output, ViewChild} from '@angular/core';
import {MoodleService} from '../../services/moodle.service';
import { MoodleToken } from 'src/app/moodle-objects-if';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss'],
})
export class LoginComponent implements OnInit {
  @ViewChild('userNameInput') userNameInput;
  @ViewChild('passwordInput') passwordInput;

  @Output() isSignedIn = new EventEmitter<boolean>();
  @Output() moodleToken = new EventEmitter<MoodleToken>();

  public isUserLoggedIn: boolean;
  public token: MoodleToken;

  private readonly MOODLE_TOKEN = 'moodle_token';
  private readonly MOODLE_USER = 'moodle_user';

  constructor(
      private storage: Storage,
      private moodleService: MoodleService
  ) {}

  ngOnInit() {}

  public async checkMoodleLoginState() {
    console.log('ionViewWillEnter - schedule.page');
    this.storage.ready().then(async () => {
      await this.getPersistedToken();
      this.token == null ? this.isUserLoggedIn = false : this.isUserLoggedIn = true;
      this.isSignedIn.emit(this.isUserLoggedIn);
      if (this.isUserLoggedIn) {
        this.moodleToken.emit(this.token);
      }
      console.log('login::ngAfterViewInit - isTokenPersisted: '+ this.token);
      console.log('login::ngAfterViewInit - isUserLoggedIn: '+ this.isUserLoggedIn);
    });
  }

  public async fetchAndPersistMoodleToken() {
    const userName = this.userNameInput.value;
    const password = this.passwordInput.value;

    const tokenToPersist = await this.moodleService.getLoginToken(userName, password).toPromise();

    if (tokenToPersist.token != undefined || tokenToPersist.token != null) {
      this.isUserLoggedIn = true;
      this.moodleToken.emit(tokenToPersist);
      console.log('login::fetchAndPersistMoodleToken - isUserLoggedIn: '+ this.isUserLoggedIn);
      this.isSignedIn.emit(this.isUserLoggedIn);
      await this.storage.set(this.MOODLE_USER, userName);
      await this.storage.set(this.MOODLE_TOKEN, tokenToPersist);
    } else {
      // if login fails moodle response contains: "errorcode":"invalidlogin"
      this.isUserLoggedIn = false;
      this.isSignedIn.emit(this.isUserLoggedIn);
      alert('invalid credentials');
    }
  }

  private async getPersistedToken() {
    await this.storage.get(this.MOODLE_TOKEN).then(value => {
      this.token = value;
      console.log('login::getPersistedToken - token: ' + this.token);
    });
  }

}
