import { Component, OnInit, ElementRef, ViewChild, Inject } from '@angular/core';

@Component({
  selector: 'app-stdg-tooltip',
  templateUrl: './stdg-tooltip.component.html',
  styleUrls: ['./stdg-tooltip.component.scss'],
})
export class StdgTooltipComponent implements OnInit {

  public tooltipText: string;
  public isTooltipVisible: boolean;
  public positionLeft: string;
  public positionTop: string;

  public readonly theme: string;
  private readonly delay: number;
  private readonly dist:number;
  private body: HTMLElement;

  @ViewChild('stdgTooltip')
  stdgTooltip: ElementRef;

  constructor(@Inject(String) theme, @Inject(Number) delay, @Inject(Number) dist) {
    this.theme = (theme === undefined || theme === null) ? 'dark' : theme;
    this.delay = (delay === undefined || delay === null) ? 0 : delay;
    this.dist = (dist === undefined || dist === null) ? 10 : dist;
  }

  /**
   * Positions the tooltip.
   * @param parent - The trigger of the tooltip.
   * @param tooltip - The tooltip itself.
   * @param posHorizontal - Desired horizontal position of the tooltip relatively to the trigger (left/center/right)
   * @param posVertical - Desired vertical position of the tooltip relatively to the trigger (top/center/bottom)
   *
   */
  private positionAt(parent:MouseEvent, posHorizontal:string, posVertical:string) {
    console.log(parent);
    console.log(this.stdgTooltip.nativeElement);
    // const parentCoords = parent.getBoundingClientRect();
    const parentCoords = {
      top: parent.y,
      left: parent.x,
      // @ts-ignore
      bottom: parent.y + parent.srcElement.clientHeight,
      // @ts-ignore
      right: parent.x + parent.srcElement.clientWidth,
      // @ts-ignore
      width: parent.srcElement.clientWidth
    };
    let left, top;

    console.log('tooltip pos vertical: '+posVertical)

    switch (posHorizontal) {
      case 'left':
        left = parentCoords.left - this.dist - this.stdgTooltip.nativeElement.offsetWidth;
        if (parentCoords.left - this.stdgTooltip.nativeElement.offsetWidth < 0) {
          left = this.dist;
        }
        break;

      case 'right':
        left = parentCoords.right + this.dist;
        if (parentCoords.right + this.stdgTooltip.nativeElement.offsetWidth > document.documentElement.clientWidth) {
          left = document.documentElement.clientWidth - this.stdgTooltip.nativeElement.offsetWidth - this.dist;
        }
        break;

      default:
      case 'center':
        left = parentCoords.left + ((parentCoords.width - this.stdgTooltip.nativeElement.offsetWidth) / 2);
    }

    switch (posVertical) {
      case 'center':
        top = (parentCoords.top + parentCoords.bottom) / 2 - this.stdgTooltip.nativeElement.offsetHeight / 2;
        break;

      case 'bottom':
        top = parentCoords.bottom + this.dist;
        break;

      default:
      case 'top':
        top = parentCoords.top - this.stdgTooltip.nativeElement.offsetHeight - this.dist;
    }

    left = (left < 0) ? parentCoords.left : left;
    top  = (top < 0) ? parentCoords.bottom + this.dist : top;

    this.positionLeft = left + 'px';
    this.positionTop  = top + pageYOffset + 'px';
  }

  private showTooltipByMouseOver(event:MouseEvent) {
    if (!event.target.hasAttribute('data-tooltip')) {
      return
    };
    console.log('show tooltip: '+this.tooltipText);

    this.isTooltipVisible = true;
    this.tooltipText = event.target.getAttribute('data-tooltip');

    const pos = event.target.getAttribute('data-position') || 'center top';
    const posHorizontal = pos.split(' ')[0];
    const posVertical = pos.split(' ')[1];

    this.positionAt(event, posHorizontal, posVertical);
  }

  private hideTooltipByMouseOut(event:MouseEvent) {
    if (event.target.hasAttribute('data-tooltip')) {
      console.log('hide tooltip');
      setTimeout(() => {
        this.isTooltipVisible = false;
      }, this.delay);
    }
  }

  public init() {
    /** Attaching one mouseover and one mouseout listener to the document instead of listeners for each trigger */
    this.body = document.body as HTMLElement;
    this.body.addEventListener('mouseover', event => {
      this.showTooltipByMouseOver(event);
    });
    this.body.addEventListener('mouseout', event => {
      this.hideTooltipByMouseOut(event);
    });
  }

  ngOnInit() {}
}