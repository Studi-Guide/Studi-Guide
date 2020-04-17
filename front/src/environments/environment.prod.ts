import {Inject} from "@angular/core";
import {DOCUMENT} from "@angular/common";

export class Env {

  serverUrl:string;
  production = true;

  constructor(@Inject(DOCUMENT) private document: Document) {
    this.serverUrl = document.location.href;
  }

}
