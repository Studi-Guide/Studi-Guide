import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class RequestBuildingDataService {

  private async:boolean = true;

  constructor() { }

  public fetchDiscoverFloorData(method, url:string, dataToSend, callback) {
    let xhr = new XMLHttpRequest();

    xhr.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
        if (callback!=undefined) {
          callback(xhr.responseText);
        } else {
          console.info("no request callback passed");
        }
      }
    };

    let data;

    if (method === 'GET') {
      // TODO adapt url part dataToSend
      url = url + '/' + dataToSend;
      data = null;
    } else if (method === 'POST') {
      data = dataToSend;
    }

    xhr.open(method, url, this.async);
    xhr.send(data);
  }
}