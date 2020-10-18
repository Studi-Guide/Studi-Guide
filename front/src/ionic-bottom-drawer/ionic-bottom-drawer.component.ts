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
      threshold: 1,
      gestureName: 'push-pull-drawer',
      direction: 'y',
      onStart: (detail => {  this.onStart(detail); }),
      onMove: (detail => { this.onMove(detail); }),
      onEnd: (detail => { this.onEnd(detail); })
    });
    this.gesture.enable();

    this.renderer.setStyle(this.element.nativeElement, 'transition', 'transform 1ms ease-in-out');
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

    if (this.state === DrawerState.Hidden) {
      this.element.nativeElement.hidden = true;
    }

     this.stateChange.emit(this.state);
  }

  private onStart(detail) {
    this.startPositionTop = this.element.nativeElement.getBoundingClientRect().top;
    const step = (this.startPositionTop)/this.platform.height();
    console.log(detail.startY, detail.currentY, detail.deltaY, step);
  }

  private onMove(detail) {
    const translate = 'translateY('+detail.currentY+'px)';
    this.renderer.setStyle(this.element.nativeElement, 'transform', translate);
  }

  private onEnd(detail) {
    const step = (this.startPositionTop+detail.deltaY)/this.platform.height();
    console.log(detail.startY, detail.currentY, detail.deltaY, step);
  }

  private async animateTo(positionY:number) {
    console.log(this.state, positionY, this.platform.height(), positionY/this.platform.height());
    const translate = 'translateY('+(this.platform.height()-positionY)+'px)';
    this.renderer.setStyle(this.element.nativeElement, 'transform', translate);

  }

}
