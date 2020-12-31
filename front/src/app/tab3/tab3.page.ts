import {AfterViewInit, Component, ElementRef, OnInit, ViewChild} from '@angular/core';
import {LoadingController} from '@ionic/angular';
import {RssFeedService} from '../services/rss-feed.service';
import {FeedItem} from '../news/rssElement';

@Component({
  selector: 'app-tab3',
  templateUrl: 'tab3.page.html',
  styleUrls: ['tab3.page.scss']
})
export class Tab3Page implements OnInit {
  private mensaFeed = 'Mensateria-Ohm';
  isMensaExpanded = false;
  mensaFeedContent = new FeedItem('', '', '', '', Date.prototype, '');

  constructor(private loadingController: LoadingController,
              private rssFeedService: RssFeedService
  ) {}

  async ngOnInit() {
    await this.fetchMensaData();
  }

  onMensaClick() {
      this.isMensaExpanded = !this.isMensaExpanded;
    }

  private async fetchMensaData() {
    const loading = await this.loadingController.create({
      message: 'Loading mensa data...'
    });
    await loading.present();
    try {
      const rssFeedItems = await this.rssFeedService.getArticles(this.mensaFeed);
      if (rssFeedItems) {
        console.log('Received ' + rssFeedItems.length + ' itens');
        this.mensaFeedContent = rssFeedItems[0];
      }
    }catch (e) {
      console.log(e);
    }
    await loading.dismiss();
  }
}
