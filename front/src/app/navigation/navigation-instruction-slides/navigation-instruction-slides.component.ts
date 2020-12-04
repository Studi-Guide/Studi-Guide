import {Component, ElementRef, Input, OnInit} from '@angular/core';
import {INavigationInstruction} from './navigation-instruction-if';

@Component({
  selector: 'app-navigation-instruction-slides',
  templateUrl: './navigation-instruction-slides.component.html',
  styleUrls: ['./navigation-instruction-slides.component.scss'],
})
export class NavigationInstructionSlidesComponent implements OnInit {

  @Input() instructions: INavigationInstruction[];

  // Optional parameters to pass to the swiper instance.
  // See http://idangero.us/swiper/api/ for valid options.
  slideOpts = {
    initialSlide: 0,
    speed: 400
  };

  constructor(private element: ElementRef) { }

  ngOnInit() {
    this.hide();
  }

  public async hide() {
    this.element.nativeElement.hidden = true;
  }

  public async show() {
    this.element.nativeElement.hidden = false;
  }

}
