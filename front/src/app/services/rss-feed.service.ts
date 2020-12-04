import { Injectable } from '@angular/core';
import {FeedItem} from '../news/rssElement';
import {Parser} from 'xml2js';
import {DataService} from './data.service';
@Injectable({
  providedIn: 'root'
})
export class RssFeedService {

  constructor(private dataService: DataService) { }

  public async getArticles(feed: string) {
    const res = await this.dataService.get_rssFeed(feed);
    return this.parseXml(res);
  }

  private parseXml(xmlStr) {
    const articles: FeedItem[] = [];
    const parser = new Parser (
        {
          trim: true,
          explicitArray: true
        });

    parser.parseString(xmlStr, (err, result) =>
    {
        if (result?.rss?.channel){
          for (const channel of result.rss.channel) {
            if (channel.item) {
              for (const item of channel.item) {
                articles.push(
                    new FeedItem(
                        item.description[0],
                        item.link[0],
                        item.title[0],
                        new Date(item.pubDate[0]),
                        item['content:encoded'][0]));
              }
            }
          }
        }
    });
    return articles;
  }
}
