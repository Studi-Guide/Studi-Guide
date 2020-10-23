import {
  AfterViewInit,
  Component,
  ElementRef,
  EventEmitter,
  Input,
  OnChanges,
  OnInit,
  Output,
  Renderer2,
  SimpleChanges
} from '@angular/core';
import {Animation, AnimationController, DomController, Gesture, GestureController, Platform} from '@ionic/angular';
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

  constructor(
      private element: ElementRef,
      private renderer: Renderer2,
      private domCtrl: DomController,
      private platform: Platform,
      private gestureCtrl: GestureController
  ) { }

  ngOnInit() {

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

    this.renderer.setStyle(this.element.nativeElement, 'transition', 'transform '+this.duration+'ms ease-in-out');
  }

  ngAfterViewInit() {
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

    if (this.state !== DrawerState.Hidden) {
      this.element.nativeElement.hidden = false;
    }

    switch (this.state) {
       case DrawerState.Top:
         await this.animateTo(this.platform.height() - this.distanceTop);
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

    if (this.state === DrawerState.Hidden) {
      this.element.nativeElement.hidden = true;
    }

     this.stateChange.emit(this.state);
  }

  private onStart(detail) {
    this.renderer.setStyle(this.element.nativeElement, 'transition', 'transform 0s ease-in-out');
    this.startPositionTop = this.element.nativeElement.getBoundingClientRect().top;
    const step = (this.startPositionTop)/this.platform.height();
    console.log(detail.startY, detail.currentY, detail.deltaY, step);
  }

  private onMove(detail) {

    if ((this.shouldBounce && detail.currentY < this.distanceTop - this.bounceDelta)
        || (!this.shouldBounce && detail.currentY < this.distanceTop)) {
      return;
    } else if ((this.shouldBounce && detail.currentY > (this.platform.height() - this.minimumHeight) + this.bounceDelta)
        || (!this.shouldBounce && detail.currentY > (this.platform.height() - this.minimumHeight))) {
      return;
    }

    const translate = 'translateY('+detail.currentY+'px)';
    this.renderer.setStyle(this.element.nativeElement, 'transform', translate);
  }

  private onEnd(detail) {
    this.renderer.setStyle(this.element.nativeElement, 'transition', 'transform '+this.duration+'ms ease-in-out');
    const step = (this.startPositionTop+detail.deltaY)/this.platform.height();
    console.log(detail.startY, detail.currentY, detail.deltaY, step);


    const newTop = detail.currentY;
    const deltaTop = Math.abs(this.distanceTop - newTop);
    const deltaDock = Math.abs(this.platform.height() - this.dockedHeight - newTop);
    const deltaBot = Math.abs(this.platform.height() - this.minimumHeight - newTop);

    if (deltaTop < deltaDock && deltaTop < deltaBot) {
      this.SetState(DrawerState.Top);
    } else if (deltaBot < deltaDock && deltaBot < deltaTop) {
      this.SetState(DrawerState.Bottom);
    } else {
      this.SetState(DrawerState.Docked);
    }
  }

  private async animateTo(positionY:number) {
    console.log(this.state, positionY, this.platform.height(), positionY/this.platform.height());
    const translate = 'translateY('+(this.platform.height()-positionY)+'px)';
    this.renderer.setStyle(this.element.nativeElement, 'transform', translate);

    await this.delay(this.duration);

  }

  private async delay(ms: number) {
    await new Promise(resolve => setTimeout(()=>resolve(), ms)).then(()=>console.log("fired"));
  }

}
