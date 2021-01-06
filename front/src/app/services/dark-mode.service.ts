import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class DarkModeService {

  // tslint:disable-next-line:variable-name
  private _isDarkMode = false;
  // tslint:disable-next-line:variable-name
  private _isAutoDarkMode = false;

  constructor() { }

  get isDarkMode(): boolean {
      return this._isDarkMode;
  }

  set isDarkMode(value: boolean) {
      this.toggleDarkTheme(value);
  }

  get isAutoDarkMode(): boolean {
    return this._isAutoDarkMode;
  }

  enableAutoDarkMode() {
    // Use matchMedia to check the user preference
    const prefersDark = window.matchMedia('(prefers-color-scheme: dark)');

    this.toggleDarkTheme(prefersDark.matches);

    // Listen for changes to the prefers-color-scheme media query
    prefersDark.addListener(this.onThemeChanged);
    this._isAutoDarkMode = true;
  }

  disableAutoDarkMode() {
    // Use matchMedia to check the user preference
    const prefersDark = window.matchMedia('(prefers-color-scheme: dark)');
    // Listen for changes to the prefers-color-scheme media query
    prefersDark.removeListener(this.onThemeChanged);
    this._isAutoDarkMode = false;
  }

  private toggleDarkTheme(shouldAdd) {
    document.body.classList.toggle('dark', shouldAdd);
    this._isDarkMode = shouldAdd;
  }

  private onThemeChanged(mediaQuery: MediaQueryListEvent) {
    this.toggleDarkTheme(mediaQuery.matches);
  }
}
