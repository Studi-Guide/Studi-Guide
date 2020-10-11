import {
  AfterViewInit,
  Component,
  ElementRef,
  EventEmitter,
  Input,
  OnChanges, OnInit,
  Output,
  Renderer2,
  SimpleChanges
} from '@angular/core';
import {AnimationController, DomController, Gesture, Animation, GestureController, Platform} from '@ionic/angular';
import {DrawerState} from './drawer-state';

@Component({
  selector: 'app-ionic-bottom-drawer',
  templateUrl: './ionic-bottom-drawer.component.html',
  styleUrls: ['./ionic-bottom-drawer.component.scss'],
})
export class IonicBottomDrawerComponent implements OnInit, AfterViewInit, OnChanges {

  @Input() distanceTop = 50;

  @Input() dockedHeight = 250;

  @Input() minimumHeight = 150;

  @Input() shouldBounce = true;

  @Input() disableDrag = false;

  @Input() easing = 'ease-in-out';

  @Input() duration = 150;

  @Input() state: DrawerState = DrawerState.Docked;

  @Input() bounceDelta = 30;

  @Output() stateChange: EventEmitter<DrawerState> = new EventEmitter<DrawerState>();

  private startPositionTop: number;
  private gesture: Gesture;
  private animation: Animation;

  constructor(
      private element: ElementRef,
      private renderer: Renderer2,
      private domCtrl: DomController,
      private platform: Platform,
      private gestureCtrl: GestureController,
      private animationCtrl: AnimationController
  ) { }

  ngOnInit() {

    this.animation = this.animationCtrl.create()
        .addElement(this.element.nativeElement)
        .easing(this.easing)
        .duration(5000)
        .fromTo('transform', 'translateY(0px)'/*begin distance to top */, 'translateY('+this.platform.height()+'px)' /* end distance to top pos */);

    this.gesture = this.gestureCtrl.create({
      el: this.element.nativeElement,
      threshold: 0,
      gestureName: 'push-pull-drawer',
      direction: 'y',
      onStart: (detail => {  this.onStart(detail); }),
      onMove: (detail => { this.onMove(detail); }),
      onEnd: (detail => { this.onEnd(detail); })
    });
    this.gesture.enable();
  }

  ngAfterViewInit() {

    //this.animation.progressStart(false);
    //this.animation.progressStep(this.minimumHeight/this.platform.height());
    //this.animation.progressEnd(1, 0, 5000);

    this.SetState(this.state);
  }

  ngOnChanges(changes: SimpleChanges) {
    if (!changes.state) {
      return;
    }

    this.SetState(changes.state.currentValue);
  }

  public async SetState(newState:DrawerState) {
    this.state = newState;

    switch (this.state) {
       case DrawerState.Top:
         await this.animateTo(this.distanceTop);
         break;
       case DrawerState.Bottom:
         await this.animateTo(this.minimumHeight);
         break;
       case DrawerState.Docked:
         await this.animateTo(this.dockedHeight);
         break;
       case DrawerState.Hidden:
         await this.animateTo(0);
         break;
     }

     this.stateChange.emit(this.state);
  }

  private onStart(detail) {
    this.startPositionTop = detail.startY//this.element.nativeElement.getBoundingClientRect().top;
    const step = (this.startPositionTop)/this.platform.height();

    this.animation.progressStart(false, step);
    console.log(detail.startY, detail.currentY, detail.deltaY, step);
  }

  private onMove(detail) {
    const step = (this.startPositionTop+detail.deltaY)/this.platform.height();
    this.animation.progressStep(step);
    console.log(detail.startY, detail.currentY, detail.deltaY, step);
  }

  private onEnd(detail) {
    //this.animation.to('transform', 'translateY('+(this.distanceTop)+'px)');
    const step = (this.startPositionTop+detail.deltaY)/this.platform.height();
    this.animation.progressStep(step);
    this.animation.pause();
    //this.animation.progressEnd(0, 1);
    console.log(detail.startY, detail.currentY, detail.deltaY, step);
    //this.animateTo(detail.currentY);
    //this.animation.stop();  resets the animation

  }

  private async animateTo(positionY:number) {
    console.log(this.state, positionY, this.platform.height(), positionY/this.platform.height());
    await this.animationCtrl.create()
        .addElement(this.element.nativeElement)
        .easing(this.easing)
        .duration(this.duration)
        .to('transform', 'translateY(' + (this.platform.height()-positionY) + 'px)')
        .iterations(1)
        .play();
  }

  // private onMove(detail) {
  //   const newTop = this.startPositionTop + detail.deltaY;
  //
  //   if ((this.shouldBounce && newTop < this.distanceTop - this.bounceDelta)
  //       || (!this.shouldBounce && newTop < this.distanceTop)) {
  //     return;
  //   } else if ((this.shouldBounce && newTop > (this.platform.height() - this.minimumHeight) + this.bounceDelta)
  //       || (!this.shouldBounce && newTop > (this.platform.height() - this.minimumHeight))) {
  //     return;
  //   }
  //
  //   this.animate(newTop);
  // }

  // private onStart(detail) {
  //   this.startPositionTop = this.element.nativeElement.getBoundingClientRect().top;
  // }
  //
  // private onEnd(detail) {
  //   const newTop = this.startPositionTop + detail.deltaY;
  //   const deltaTop = Math.abs(this.distanceTop - newTop);
  //   const deltaDock = Math.abs(this.platform.height() - this.dockedHeight - newTop);
  //   const deltaBot = Math.abs(this.platform.height() - this.minimumHeight - newTop);
  //
  //   if (deltaTop < deltaDock && deltaTop < deltaBot) {
  //     this.SetState(DrawerState.Top);
  //   } else if (deltaBot < deltaDock && deltaBot < deltaTop) {
  //     this.SetState(DrawerState.Bottom);
  //   } else {
  //     this.SetState(DrawerState.Docked);
  //   }
  // }
  //
  // private async animate(position:number) {
  //   await this.animationCtrl.create()
  //       .addElement(this.element.nativeElement)
  //       .easing(this.easing)
  //       .duration(this.duration)
  //       .to('transform', 'translateY(' + position + 'px)')
  //       .iterations(1)
  //       .play();
  // }

}
