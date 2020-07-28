import {
  AfterViewInit,
  Component,
  ElementRef,
  EventEmitter,
  Input,
  OnChanges,
  OnInit,
  Output,
  Renderer2, SimpleChanges
} from '@angular/core';
import {DomController, Platform, Gesture, GestureController} from "@ionic/angular";
import {DrawerState} from "./drawer-state";

@Component({
  selector: 'app-ionic-bottom-drawer',
  templateUrl: './ionic-bottom-drawer.component.html',
  styleUrls: ['./ionic-bottom-drawer.component.scss'],
})
export class IonicBottomDrawerComponent implements AfterViewInit, OnChanges {

  @Input() dockedHeight = 250;

  @Input() shouldBounce = true;

  @Input() disableDrag = false;

  @Input() distanceTop = 100;

  @Input() transition = '0.25s ease-in-out';

  @Input() state: DrawerState = DrawerState.Bottom;

  @Input() minimumHeight = 50;

  @Output() stateChange: EventEmitter<DrawerState> = new EventEmitter<DrawerState>();

  private _startPositionTop: number;
  private readonly _BOUNCE_DELTA = 30;
  private gesture: Gesture;

  constructor(
      private element: ElementRef,
      private renderer: Renderer2,
      private domCtrl: DomController,
      private platform: Platform,
      private gestureCtrl: GestureController
  ) { }

  ngAfterViewInit() {


    this.gesture = this.gestureCtrl.create({
      el: this.element.nativeElement,
      threshold: 15,
      gestureName: 'swipe-up',
      direction: "y",
      onMove: (detail => { this.onMove(detail); })
    });



    // Codepen Snippet for Gesture
    //
    // JS
    //   import { createGesture } from 'https://cdn.jsdelivr.net/npm/@ionic/core/dist/esm/index.mjs';
    //
    //   let p = document.querySelector('p');
    //   const gesture = createGesture({
    //     el: document.querySelector('.rectangle'),
    //     direction: "y",
    //     onMove: (detail) => { onMove(detail); }
    //   })
    //
    //   gesture.enable(true);
    //
    //   const onMove = (detail) => {
    //     const type = detail.type;
    //     const currentX = detail.currentX;
    //     const deltaX = detail.deltaX;
    //     const velocityX = detail.velocityX;
    //     const currentY = detail.currentY;
    //     const deltaY = detail.deltaY;
    //     const velocitY = detail.velocityY;
    //
    //     p.innerHTML = `
    //   <div>Type: ${type}</div>
    //   <div>Current X: ${currentX}</div>
    //   <div>Delta X: ${deltaX}</div>
    //   <div>Velocity X: ${velocityX}</div>
    //   <div>Current Y: ${currentY}</div>
    //   <div>Delta Y: ${deltaY}</div>
    //   <div>Velocity Y: ${velocitY}</div>
    // `
    //   }
    //
    // HTML
    //
    // <html>
    //     <head></head>
    // <body>
    // <div class="rectangle"></div>
    //     <p>
    //     Swipe to start tracking
    // </p>
    // </body>
    // </html>
    //
    // CSS
    //
    // .rectangle {
    //     width: 100px;
    //     height: 400px;
    //     background: rgba(0, 0, 255, 0.5);
    //   }




    this.renderer.setStyle(this.element.nativeElement.querySelector('.ion-bottom-drawer-scrollable-content :first-child'),
        'touch-action', 'none');
    this._setDrawerState(this.state);

    const hammer = new Hammer(this.element.nativeElement);
    hammer.get('pan').set({ enable: true, direction: Hammer.DIRECTION_VERTICAL });
    hammer.on('pan panstart panend', (ev: any) => {
      if (this.disableDrag) {
        return;
      }

      switch (ev.type) {
        case 'panstart':
          this._handlePanStart();
          break;
        case 'panend':
          this._handlePanEnd(ev);
          break;
        default:
          this._handlePan(ev);
      }
    });
  }

  ngOnChanges(changes: SimpleChanges) {
    if (!changes.state) {
      return;
    }

    this._setDrawerState(changes.state.currentValue);
  }

  private onMove(detail) {
    console.log(detail.type, detail.currentX, detail.deltaX,  detail.velocityX, detail.currentY, detail.velocityY);
  }

  private _setDrawerState(state: DrawerState) {
    this.renderer.setStyle(this.element.nativeElement, 'transition', this.transition);
    switch (state) {
      case DrawerState.Bottom:
        this._setTranslateY('calc(100vh - ' + this.minimumHeight + 'px)');
        break;
      case DrawerState.Docked:
        this._setTranslateY((this.platform.height() - this.dockedHeight) + 'px');
        break;
      default:
        this._setTranslateY(this.distanceTop + 'px');
    }
  }

  private _handlePanStart() {
    this._startPositionTop = this.element.nativeElement.getBoundingClientRect().top;
  }

  private _handlePanEnd(ev) {
    if (this.shouldBounce && ev.isFinal) {
      this.renderer.setStyle(this.element.nativeElement, 'transition', this.transition);

      switch (this.state) {
        case DrawerState.Docked:
          this._handleDockedPanEnd(ev);
          break;
        case DrawerState.Top:
          this._handleTopPanEnd(ev);
          break;
        default:
          this._handleBottomPanEnd(ev);
      }
    }
    this.stateChange.emit(this.state);
  }

  private _handleTopPanEnd(ev) {
    if (ev.deltaY > this._BOUNCE_DELTA) {
      this.state = DrawerState.Docked;
    } else {
      this._setTranslateY(this.distanceTop + 'px');
    }
  }

  private _handleDockedPanEnd(ev) {
    const absDeltaY = Math.abs(ev.deltaY);
    if (absDeltaY > this._BOUNCE_DELTA && ev.deltaY < 0) {
      this.state = DrawerState.Top;
    } else if (absDeltaY > this._BOUNCE_DELTA && ev.deltaY > 0) {
      this.state = DrawerState.Bottom;
    } else {
      this._setTranslateY((this.platform.height() - this.dockedHeight) + 'px');
    }
  }

  private _handleBottomPanEnd(ev) {
    if (-ev.deltaY > this._BOUNCE_DELTA) {
      this.state = DrawerState.Docked;
    } else {
      this._setTranslateY('calc(100vh - ' + this.minimumHeight + 'px)');
    }
  }

  private _handlePan(ev) {
    const pointerY = ev.center.y;
    this.renderer.setStyle(this.element.nativeElement, 'transition', 'none');
    if (pointerY > 0 && pointerY < this.platform.height()) {
      if (ev.additionalEvent === 'panup' || ev.additionalEvent === 'pandown') {
        const newTop = this._startPositionTop + ev.deltaY;
        if (newTop >= this.distanceTop) {
          this._setTranslateY(newTop + 'px');
        } else if (newTop < this.distanceTop) {
          this._setTranslateY(this.distanceTop + 'px');
        }
        if (newTop > this.platform.height() - this.minimumHeight) {
          this._setTranslateY((this.platform.height() - this.minimumHeight) + 'px');
        }
      }
    }
  }

  private _setTranslateY(value) {
    this.domCtrl.write(() => {
      this.renderer.setStyle(this.element.nativeElement, 'transform', 'translateY(' + value + ')');
    });
  }

}
