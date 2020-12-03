import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {FeedItem} from '../news/rssElement';
import {Parser} from 'xml2js';
@Injectable({
  providedIn: 'root'
})
export class RssFeedService {

  constructor(private http: HttpClient) { }

  public async getArticlesForUrl(feedUrl: string) {
    const res = await this.http.get(feedUrl,
        {headers:new HttpHeaders()
              .set('Access-Control-Allow-Origin', '*')})
        .toPromise();

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

    });
    return articles;
  }
}
