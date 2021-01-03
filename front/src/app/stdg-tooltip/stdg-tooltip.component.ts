import {Component, ElementRef, ViewChild, AfterViewInit} from '@angular/core';

@Component({
  selector: 'stdg-tooltip',
  templateUrl: './stdg-tooltip.component.html',
  styleUrls: ['./stdg-tooltip.component.scss'],
})
export class StdgTooltipComponent implements AfterViewInit {

  public tooltipText: string;
  public tooltipClass = 'stdg-tooltip-hide';

  public theme: string;
  private mouseOverDelay: number;
  private mouseOutDelay: number;
  private dist: number;
  private rounded: boolean;

  @ViewChild('stdgTooltip') stdgTooltip: ElementRef;

  constructor() {}

  /**
   * Positions the tooltip.
   * @param parent - The trigger of the tooltip.
   * @param posHorizontal - Desired horizontal position of the tooltip relatively to the trigger (left/center/right)
   * @param posVertical - Desired vertical position of the tooltip relatively to the trigger (top/center/bottom)
   */
  private async positionAt(parent: MouseEvent, posHorizontal: string, posVertical: string) {
    const target = (parent.target as HTMLElement);
    const parentCoords = {
      top: target.getBoundingClientRect().y,
      left: target.getBoundingClientRect().x,
      bottom: target.getBoundingClientRect().y + (parent.target as HTMLElement).clientHeight,
      right: target.getBoundingClientRect().x + (parent.target as HTMLElement).clientWidth,
      width: (parent.target as HTMLElement).clientWidth
    };
    let left, top;

    await this.delayAngularRenderingCycles(1);
    const offsetWidth = this.stdgTooltip.nativeElement.offsetWidth;
    const offsetHeight = this.stdgTooltip.nativeElement.offsetHeight;

    switch (posHorizontal) {
      case 'left':
        left = parentCoords.left - this.dist - offsetWidth;
        if (parentCoords.left - offsetWidth < 0) {
          left = this.dist;
        }
        break;

      case 'right':
        left = parentCoords.right + this.dist;
        if (parentCoords.right + offsetWidth > document.documentElement.clientWidth) {
          left = document.documentElement.clientWidth - offsetWidth - this.dist;
        }
        break;

      default:
      case 'center':
        left = parentCoords.left + ((parentCoords.width - offsetWidth) / 2);
    }

    switch (posVertical) {
      case 'center':
        top = (parentCoords.top + parentCoords.bottom) / 2 - offsetHeight / 2;
        break;

      case 'bottom':
        top = parentCoords.bottom + this.dist;
        break;

      default:
      case 'top':
        top = parentCoords.top - offsetHeight - this.dist;
    }

    left = (left < 0) ? parentCoords.left : left;
    top  = (top < 0) ? parentCoords.bottom + this.dist : top;

    this.stdgTooltip.nativeElement.style.left = left + 'px';
    this.stdgTooltip.nativeElement.style.top  = top + pageYOffset + 'px';
  }

  private delayAngularRenderingCycles(cycleTimes: number) {
    const ms = cycleTimes * 16;
    return new Promise(resolve => setTimeout(resolve, ms));
  }

  private showTooltipByMouseOver(event: MouseEvent) {
    if (!event.target.hasAttribute('data-tooltip')) {
      return;
    }

    setTimeout(() => {
      this.tooltipText = event.target.getAttribute('data-tooltip');
      this.showTooltip();

      const pos = event.target.getAttribute('data-tooltip-pos') || 'center top';
      const posHorizontal = pos.split(' ')[0];
      const posVertical = pos.split(' ')[1];

      this.positionAt(event, posHorizontal, posVertical).then();
    }, this.mouseOverDelay);
  }

  private showTooltip() {
    this.tooltipClass = 'stdg-tooltip stdg-tooltip-' + this.theme;
    this.tooltipClass += this.rounded ? ' stdg-tooltip-rounded' : '';
  }

  private hideTooltipByMouseOut(event: MouseEvent) {
    if (event.target.hasAttribute('data-tooltip')) {
      setTimeout(() => {
        this.hideTooltip();
      }, this.mouseOutDelay);
    }
  }

  private hideTooltip() {
    this.tooltipClass = 'stdg-tooltip-hide';
  }

  /**
   * Initializes the tooltip.
   * @param theme - Theme of the tooltip: light, medium, or dark
   * @param mouseOverDelay - delay in ms before the tooltip is shown
   * @param mouseOutDelay - delay in ms before the tooltip is hidden
   * @param dist - distance in px between target element and tooltip
   * @param rounded - if the tooltip corners are rounded or not
   */
  private init(theme: string, mouseOverDelay: number, mouseOutDelay: number, dist: number, rounded: boolean) {
    this.theme = (theme === undefined || theme === null) ? 'dark' : theme;
    this.mouseOverDelay = (mouseOverDelay === undefined || mouseOverDelay === null) ? 0 : mouseOverDelay;
    this.mouseOutDelay = (mouseOutDelay === undefined || mouseOutDelay === null) ? 0 : mouseOutDelay;
    this.dist = (dist === undefined || dist === null) ? 10 : dist;
    this.rounded = (rounded === undefined || rounded === null) ? false : rounded;

    /** Attaching one mouseover and one mouseout listener to the document instead of listeners for each trigger */
    const body = document.body as HTMLElement;
    body.addEventListener('mouseover', event => {
      this.showTooltipByMouseOver(event);
    });
    body.addEventListener('mouseout', event => {
      this.hideTooltipByMouseOut(event);
    });
  }

  ngAfterViewInit(){
    this.init('medium', 0, 0, 0, true);
  }
}
