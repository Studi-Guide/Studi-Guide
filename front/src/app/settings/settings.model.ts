import {Injectable} from '@angular/core';
import {Storage} from '@ionic/storage';
import {IonicBottomDrawerComponent} from '../../ionic-bottom-drawer/ionic-bottom-drawer.component';
import {SettingsEntity} from './settings.entity';
import {DarkModeService} from '../services/dark-mode.service';

@Injectable({
    providedIn: 'root'
})
export class SettingsModel {
    private static settingsKey = 'settings';
    private entity: SettingsEntity;

    constructor(private storage: Storage, private darkMode: DarkModeService) {
        this.storage.get(SettingsModel.settingsKey).then(v => {
            this.entity = Object.assign(new SettingsEntity(), JSON.parse(v));

            IonicBottomDrawerComponent.DrawerDocking = this.entity.DrawerDocking;
            if (this.entity.AutoDarkMode) {
                darkMode.enableAutoDarkMode();
            } else {
                darkMode.isDarkMode = this.entity.DarkMode;
            }
        });
    }

    public get DrawerDocking() {
        return IonicBottomDrawerComponent.DrawerDocking;
    }

    public set DrawerDocking(value: boolean) {
        this.entity.DrawerDocking = value;
        IonicBottomDrawerComponent.DrawerDocking = value;
        this.persist();
    }

    public get DarkMode() {
        return this.darkMode.isDarkMode;
    }

    public set DarkMode(value: boolean) {
        this.entity.DarkMode = value;
        this.darkMode.isDarkMode = value;
        this.persist();
    }

    public get AutoDarkMode() {
        return this.darkMode.isAutoDarkMode;
    }

    public set AutoDarkMode(value: boolean) {
        this.entity.AutoDarkMode = value;
        if (value){
            this.darkMode.enableAutoDarkMode();
        } else {
            this.darkMode.disableAutoDarkMode();
        }

        this.persist();
    }

    private persist() {
        this.storage.set(SettingsModel.settingsKey,
            JSON.stringify(this.entity)
        );
    }
}


export class SettingsChangedEventArgs {
    constructor(
        public key: string,
        public value: string
    ) {
    }
}

