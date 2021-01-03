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

  @ViewChild('stdgTooltip') stdgTooltip: ElementRef;

  constructor() {}

  /**
   * Positions the tooltip.
   * @param parent - The trigger of the tooltip.
   * @param posHorizontal - Desired horizontal position of the tooltip relatively to the trigger (left/center/right)
   * @param posVertical - Desired vertical position of the tooltip relatively to the trigger (top/center/bottom)
   */
  private positionAt(parent: MouseEvent, posHorizontal: string, posVertical: string) {
    console.log(parent);
    const target = (parent.target as HTMLElement);
    const parentCoords = {
      top: target.getBoundingClientRect().y,
      left: target.getBoundingClientRect().x,
      bottom: target.getBoundingClientRect().y + (parent.target as HTMLElement).clientHeight,
      right: target.getBoundingClientRect().x + (parent.target as HTMLElement).clientWidth,
      width: (parent.target as HTMLElement).clientWidth
    };
    let left, top;

    // TODO const offsetWidth = this.stdgTooltip.nativeElement.offsetWidth;
    // TODO const offsetHeight = this.stdgTooltip.nativeElement.offsetHeight
    const offsetWidth = this.tooltipText.length * 2.4; // settings drawer docking: 118
    const offsetHeight = this.tooltipText.length + 24/*(200 / this.tooltipText.length)*/; // settings drawer docking: 54

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

      this.positionAt(event, posHorizontal, posVertical);
    }, this.mouseOverDelay);
  }

  private showTooltip() {
    this.tooltipClass = 'stdg-tooltip stdg-tooltip-' + this.theme;
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

  private init(theme, mouseOverDelay, mouseOutDelay, dist) {
    this.theme = (theme === undefined || theme === null) ? 'dark' : theme;
    this.mouseOverDelay = (mouseOverDelay === undefined || mouseOverDelay === null) ? 0 : mouseOverDelay;
    this.mouseOutDelay = (mouseOutDelay === undefined || mouseOutDelay === null) ? 0 : mouseOutDelay;
    this.dist = (dist === undefined || dist === null) ? 10 : dist;

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
    this.init('dark', 0, 0, 0);
  }
}
