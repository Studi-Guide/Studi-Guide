import {AfterViewInit, Component} from '@angular/core';
import {LoadingController} from '@ionic/angular';
import {FeedItem} from './rssElement';
import {RssFeedService} from '../services/rss-feed.service';

@Component({
  selector: 'app-schedule',
  templateUrl: 'news.page.html',
  styleUrls: ['news.page.scss']
})
export class NewsPage implements AfterViewInit{

  private rssFeeds:string[] = [
      'https://www.th-nuernberg.de/news-archiv/rss.xml',
      'https://www.th-nuernberg.de/calendarRSS.xml'
  ]

  public isMoodleUserSignedIn: boolean;
  rssFeed: FeedItem[] = [];

  constructor(
      public loadingController: LoadingController,
      private rssFeedService: RssFeedService
  ) {}

  async ngAfterViewInit() {
    await this.fetchNewsData();
  }

  public async fetchNewsData() {
    const loading = await this.loadingController.create({
      message: 'Collecting news data...'
    });
    await loading.present();

    const items: FeedItem[] = [];
    for(const feed of this.rssFeeds)
    {
      try {
        const rssFeedItems = await this.rssFeedService.getArticlesForUrl(feed);
        items.push(...rssFeedItems);
      }catch (e) {
        console.log(e);
      }
    }

    this.rssFeed = items;
    await loading.dismiss();

  }

  public async doRefreshEvents(event) {
    await this.fetchNewsData();
    event.target.complete();
  }

  onFeedClick(element: FeedItem) {
    element.isExpanded = !element.isExpanded;
  }
}
