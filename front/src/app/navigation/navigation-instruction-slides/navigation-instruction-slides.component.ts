import {Component, ElementRef, EventEmitter, Input, OnInit, Output, ViewChild} from '@angular/core';
import {INavigationInstruction} from './navigation-instruction-if';
import {IonSlides} from "@ionic/angular";


@Component({
  selector: 'app-navigation-instruction-slides',
  templateUrl: './navigation-instruction-slides.component.html',
  styleUrls: ['./navigation-instruction-slides.component.scss']
})
export class NavigationInstructionSlidesComponent implements OnInit {

  @Input() instructions: INavigationInstruction[];

  @Output() nextSlide : EventEmitter<number> = new EventEmitter<number>();
  @Output() previousSlide : EventEmitter<number> = new EventEmitter<number>();

  @ViewChild('ionSlides') ionSlides : IonSlides;

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

  public async onPrevSlide() {
    this.previousSlide.emit(await this.ionSlides.getActiveIndex());
  }

  public async onNextSlide() {
    this.nextSlide.emit(await this.ionSlides.getActiveIndex());
  }

  public async hide() {
    this.element.nativeElement.hidden = true;
  }

  public async show() {
    this.element.nativeElement.hidden = false;
  }

}
