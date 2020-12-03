import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {FeedItem} from '../news/rssElement';
import {Parser} from 'xml2js';
import {DataService} from './data.service';
@Injectable({
  providedIn: 'root'
})
export class RssFeedService {

  constructor(private dataService: DataService) { }

  public async getArticlesForUrl(feedUrl: string) {
    const res = await this.dataService.get_proxy_request_asText(btoa(feedUrl), {responseType: 'text'});
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
                        ''));
              }
            }
          }
        }
    });
    return articles;
  }
}
