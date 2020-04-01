import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class RequestBuildingDataService {

  constructor() { }

  private buildingDataUrl:string = "http://127.0.0.1:8080/roomlist/"; // "https://example.com/"

  // TODO uncomment dataToSend when API is built
  public fetchDiscoverFloorData(method, /*dataToSend,*/ callback) {
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
    // TODO solve blocked Cross-Origin request caused by different ports of the dev servers (8100 & 8080)
    const async:boolean = true;
    xhr.open(method, this.buildingDataUrl, async);
    xhr.send(/*dataToSend*/);
  }
}