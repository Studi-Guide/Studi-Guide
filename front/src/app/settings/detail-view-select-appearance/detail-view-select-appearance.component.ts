import { Component, OnInit } from '@angular/core';
import {SettingsModel} from '../settings.model';

@Component({
  selector: 'app-detail-view-select-appearance',
  templateUrl: './detail-view-select-appearance.component.html',
  styleUrls: ['./detail-view-select-appearance.component.scss'],
})
export class DetailViewSelectAppearanceComponent implements OnInit {
  selectedAppearance: string;

  constructor(private settingsModel: SettingsModel) {
    if (settingsModel.AutoDarkMode) {
      this.selectedAppearance = 'system';
    } else {
      this.selectedAppearance = settingsModel.DarkMode ? 'dark' : 'light';
    }
  }

  ngOnInit() {}

  onAppearanceChanged(event: any) {
    if (event?.detail?.value) {
      switch (event.detail.value) {
        case 'system':
          this.settingsModel.AutoDarkMode = true;
          break;
        case 'light':
          if (this.settingsModel.AutoDarkMode) {
            this.settingsModel.AutoDarkMode = false;
          }

          this.settingsModel.DarkMode = false;
          break;
        case 'dark':
          if (this.settingsModel.AutoDarkMode) {
            this.settingsModel.AutoDarkMode = false;
          }

          this.settingsModel.DarkMode = true;
          break;
      }
    }
  }
}
