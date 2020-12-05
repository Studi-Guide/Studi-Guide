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
import {DomController, Gesture, GestureController, Platform} from '@ionic/angular';
import {DrawerState} from './drawer-state';

@Component({
  selector: 'app-ionic-bottom-drawer',
  templateUrl: './ionic-bottom-drawer.component.html',
  styleUrls: ['./ionic-bottom-drawer.component.scss'],
})
export class IonicBottomDrawerComponent implements OnInit, AfterViewInit, OnChanges {
  constructor(
      private element: ElementRef,
      private renderer: Renderer2,
      private domCtrl: DomController,
      private platform: Platform,
      private gestureCtrl: GestureController
  ) { }

  /**
   * Application Wide Drawer Docking Setting
   */
  public static DrawerDocking = true;

  @Input() gripElementsClass = 'drawer-grip';

  @Input() distanceTop = 50;

  @Input() dockedHeight = 300;

  @Input() minimumHeight = 200;

  @Input() shouldBounce = true;

  @Input() disableDrag = false;

  @Input() easing = 'ease-in-out';

  @Input() duration = 150;

  @Input() state: DrawerState = DrawerState.Docked;

  @Input() bounceDelta = 30;

  @Output() stateChange: EventEmitter<DrawerState> = new EventEmitter<DrawerState>();

  private startPositionTop: number;
  private gesture: Gesture;

  public static GetRecommendedDrawerStateForDevice():DrawerState {
    const isSmallDevice: boolean = window.matchMedia('(max-width: 767.98px)').matches;
    const isMediumDevice: boolean = window.matchMedia('(min-width: 768px)').matches;
    const isBigDevice: boolean = window.matchMedia('(min-width: 1200px)').matches;

    if (isSmallDevice) {
      return DrawerState.Bottom;
    } else if (isMediumDevice && !isBigDevice) {
      return DrawerState.Docked;
    } else if (isMediumDevice && isBigDevice) {
      return DrawerState.Top;
    }
  }

  ngOnInit() {

    const element = this.element.nativeElement.querySelector('.'+this.gripElementsClass);

    if (element === null) {
      throw new Error('can not find any element with the class name "'
          +this.gripElementsClass+'" for gesture initialization of IonicBottomDrawer');
    }

    this.gesture = this.gestureCtrl.create({
      el: element,
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
  }

  private onMove(detail) {
    const isMovementSensibilityBufferExceeded = detail.deltaY > 10 || detail.deltaX > 10;
    if (isMovementSensibilityBufferExceeded) {
      if (IonicBottomDrawerComponent.DrawerDocking) {
        if ((this.shouldBounce && detail.currentY < this.distanceTop - this.bounceDelta)
            || (!this.shouldBounce && detail.currentY < this.distanceTop)) {
          return;
        } else if ((this.shouldBounce && detail.currentY > (this.platform.height() - this.minimumHeight) + this.bounceDelta)
            || (!this.shouldBounce && detail.currentY > (this.platform.height() - this.minimumHeight))) {
          return;
        }
      } else {
        if (detail.currentY < this.distanceTop) {
          return;
        } else if (detail.currentY > (this.platform.height() - this.minimumHeight)) {
          return;
        }
      }

      const translate = 'translateY('+detail.currentY+'px)';
      this.renderer.setStyle(this.element.nativeElement, 'transform', translate);
    }
  }

  private onEnd(detail) {
    this.renderer.setStyle(this.element.nativeElement, 'transition', 'transform '+this.duration+'ms ease-in-out');

    if (!IonicBottomDrawerComponent.DrawerDocking) {
      return;
    }


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
    const translate = 'translateY('+(this.platform.height()-positionY)+'px)';
    this.renderer.setStyle(this.element.nativeElement, 'transform', translate);

    await this.delay(this.duration);

  }

  private async delay(ms: number) {
    await new Promise(resolve => setTimeout(()=>resolve(), ms)).then();
  }

}
