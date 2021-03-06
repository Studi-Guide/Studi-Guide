import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class DarkModeService {

  // tslint:disable-next-line:variable-name
  private _isDarkMode = false;
  // tslint:disable-next-line:variable-name
  private _isAutoDarkMode = false;
  private handler: (e) => void;
  private prefersDark: MediaQueryList;


  constructor() {
    this.prefersDark = window.matchMedia('(prefers-color-scheme: dark)');
  }

  get isDarkMode(): boolean {
      return this._isDarkMode;
  }

  set isDarkMode(value: boolean) {
      this.setDarkTheme(value);
  }

  get isAutoDarkMode(): boolean {
    return this._isAutoDarkMode;
  }

  set isAutoDarkMode(value: boolean) {
    if (value === true) {
      this.enableAutoDarkMode();
    } else {
      this.disableAutoDarkMode();
    }
  }

  private enableAutoDarkMode() {
    this.setDarkTheme(this.prefersDark.matches);
    this.handler = function toggleDarkTheme(e) {
        document.body.classList.toggle('dark', e.matches);
    };

    // Listen for changes to the prefers-color-scheme media query
    this. prefersDark.addListener(this.handler);
    this._isAutoDarkMode = true;
  }

  private disableAutoDarkMode() {
    this.prefersDark.removeListener(this.handler);
    this._isAutoDarkMode = false;
    this._isDarkMode = this.prefersDark.matches;
  }

  private setDarkTheme(enable: boolean) {
      document.body.classList.toggle('dark', enable);
      this._isDarkMode = enable;
  }
}
