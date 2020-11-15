import {Injectable} from '@angular/core';
import {Storage} from '@ionic/storage';
import {IonicBottomDrawerComponent} from '../../ionic-bottom-drawer/ionic-bottom-drawer.component';


@Injectable({
    providedIn: 'root'
})
export class SettingsModel {

    private static settingsKey = 'settings';

    public get DrawerDocking() {
        return IonicBottomDrawerComponent.DrawerDocking;
    }

    public set DrawerDocking(value:boolean) {
        IonicBottomDrawerComponent.DrawerDocking = value;
        // this.storage... persist setting
        this.persist();
    }

    constructor(private storage:Storage) {
        this.storage.get(SettingsModel.settingsKey).then(v => {
            const settings = JSON.parse(v);
            if (settings === null) {
                // set initial value
                this.DrawerDocking = IonicBottomDrawerComponent.DrawerDocking;
            } else {
                IonicBottomDrawerComponent.DrawerDocking = settings.DrawerDocking;
            }
        })
    }

    private persist() {
        this.storage.set(SettingsModel.settingsKey,
            JSON.stringify({DrawerDocking: this.DrawerDocking})
        );
    }
}