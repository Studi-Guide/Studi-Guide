import {Inject, Injectable} from "@angular/core";
import {DOCUMENT} from "@angular/common";

export const environment = {
  production: true
}

@Injectable()
export class Env {

  serverUrl:string;
  production = true;

  constructor(@Inject(DOCUMENT) private document: Document) {
    this.serverUrl = document.location.origin;
  }

}
