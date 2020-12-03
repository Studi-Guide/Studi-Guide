import {Component} from '@angular/core';
import {LoadingController} from '@ionic/angular';


@Component({
  selector: 'app-schedule',
  templateUrl: 'news.page.html',
  styleUrls: ['news.page.scss']
})
export class NewsPage {
  public calenderEvents: Event[] = [];
  public isMoodleUserSignedIn: boolean;

  constructor(
      public loadingController: LoadingController,
  ) {}

  public async fetchNewsData() {
    const loading = await this.loadingController.create({
      message: 'Collecting news data...'
    });
    await loading.present();

    // TODO add load logic
    await loading.dismiss();
  }

  public async doRefreshEvents(event) {
    await this.fetchNewsData();
  }
}
