import {
  AfterViewInit,
  Component,
  ElementRef,
  EventEmitter,
  Input,
  OnChanges,
  Output,
  Renderer2,
  SimpleChanges
} from '@angular/core';
import {AnimationController, DomController, Gesture, GestureController, Platform} from '@ionic/angular';
import {DrawerState} from './drawer-state';

@Component({
  selector: 'app-ionic-bottom-drawer',
  templateUrl: './ionic-bottom-drawer.component.html',
  styleUrls: ['./ionic-bottom-drawer.component.scss'],
})
export class IonicBottomDrawerComponent implements AfterViewInit, OnChanges {

  @Input() distanceTop = 50;

  @Input() dockedHeight = 250;

  @Input() minimumHeight = 150;

  @Input() shouldBounce = true;

  @Input() disableDrag = false;

  @Input() easing = 'ease-in-out';

  @Input() duration = 300;

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
      private gestureCtrl: GestureController,
      private animationCtrl: AnimationController
  ) { }

  ngAfterViewInit() {

    this.gesture = this.gestureCtrl.create({
      el: this.element.nativeElement,
      threshold: 15,
      gestureName: 'swipe-up',
      direction: 'y',
      onMove: (detail => { this.onMove(detail); }),
      onStart: (detail => {  this.onStart(detail); }),
      onEnd: (detail => { this.onEnd(detail); })
    });
    this.gesture.enable();

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
        await this.animate(this.distanceTop);
        break;
      case DrawerState.Bottom:
        await this.animate(this.platform.height() - this.minimumHeight);
        break;
      case DrawerState.Docked:
        await this.animate(this.platform.height() - this.dockedHeight);
        break;
      case DrawerState.Hidden:
        await this.animate(this.platform.height());
        break;
    }

    this.stateChange.emit(this.state);
  }

  private onMove(detail) {
    const newTop = this.startPositionTop + detail.deltaY;

    if ((this.shouldBounce && newTop < this.distanceTop - this.bounceDelta)
        || (!this.shouldBounce && newTop < this.distanceTop)) {
      return;
    } else if ((this.shouldBounce && newTop > (this.platform.height() - this.minimumHeight) + this.bounceDelta)
        || (!this.shouldBounce && newTop > (this.platform.height() - this.minimumHeight))) {
      return;
    }

    this.animate(newTop);
  }

  private onStart(detail) {
    this.startPositionTop = this.element.nativeElement.getBoundingClientRect().top;
  }

  private onEnd(detail) {
    const newTop = this.startPositionTop + detail.deltaY;
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

  private async animate(position:number) {
    await this.animationCtrl.create()
        .addElement(this.element.nativeElement)
        .easing(this.easing)
        .duration(this.duration)
        .to('transform', 'translateY(' + position + 'px)')
        .iterations(1)
        .play();
  }

}
