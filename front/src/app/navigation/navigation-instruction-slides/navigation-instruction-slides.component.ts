import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-navigation-instruction-slides',
  templateUrl: './navigation-instruction-slides.component.html',
  styleUrls: ['./navigation-instruction-slides.component.scss'],
})
export class NavigationInstructionSlidesComponent implements OnInit {

  // Optional parameters to pass to the swiper instance.
  // See http://idangero.us/swiper/api/ for valid options.
  slideOpts = {
    initialSlide: 1,
    speed: 400
  };

  constructor() { }

  ngOnInit() {}

}
