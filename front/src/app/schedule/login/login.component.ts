import {Component, EventEmitter, OnInit, Output, ViewChild} from '@angular/core';

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
  getName(): string {
    return this.name;
  }
  getPassword(): string {
    return this.password;
  }
  signIn() {
    // TODO implement login at moodle
    alert('sign in');
    this.signedIn = true;
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

  constructor() {
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
    console.log('user: '+ this.user.getName()+'\npassword: '+this.user.getPassword());
    this.user.signIn();
    this.moodleUserIsLoggedIn = this.isUserSignedIn();
    this.isSignedIn.emit(this.isUserSignedIn());
  }

  public isUserSignedIn() {
    return this.user.isSignedIn();
  }

}
