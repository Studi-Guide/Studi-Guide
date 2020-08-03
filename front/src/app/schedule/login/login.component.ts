import { Storage } from '@ionic/storage';
import {Component, EventEmitter, OnInit, Output, ViewChild} from '@angular/core';
import {MoodleService} from '../../services/moodle.service';

class MoodleUser {
  private name: string;
  private password: string;
  private signedIn: boolean;

  setName(value: string) {
    this.name = value;
  }
  setPassword(value: string) {
    this.password = value;
  }
  getPassword(): string {
    return this.password;
  }
  isSignedIn(): boolean {
    return this.signedIn;
  }
}

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss'],
})
export class LoginComponent implements OnInit {
  @ViewChild('userNameInput') userNameInput;
  @ViewChild('passwordInput') passwordInput;

  @Output() isSignedIn = new EventEmitter<boolean>();

  private user: MoodleUser;
  public moodleUserIsLoggedIn;

  constructor(private storage: Storage, private moodleService: MoodleService) {
    this.user = new MoodleUser();
  }

  ngOnInit() {}

  // tslint:disable-next-line:use-lifecycle-interface
  ngAfterViewInit() {
    this.moodleUserIsLoggedIn = this.isUserSignedIn();
  }

  signIn() {
    this.user.setName(this.userNameInput.value);
    this.user.setPassword(this.passwordInput.value);
    // TODO add MoodleTokenPersistence
    const tokenToPersist = this.moodleService.getLoginToken(this.userNameInput.value, this.passwordInput.value).toPromise();
    // this.storage.set(this.userNameInput.value, tokenToPersist);

    this.moodleUserIsLoggedIn = this.isUserSignedIn();
    this.isSignedIn.emit(this.isUserSignedIn());
  }

  public isUserSignedIn() {
    // TODO check if a token with user name is persisted
    return true;
  }

}
