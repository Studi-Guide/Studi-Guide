import {AfterViewInit, Component, ElementRef, Input, Renderer2, ViewChild} from '@angular/core';

@Component({
  selector: 'app-expander',
  templateUrl: './ionic-expander.component.html',
  styleUrls: ['./ionic-expander.component.scss'],
})
export class IonicExpanderComponent implements AfterViewInit {
  @ViewChild('expandWrapper', { read: ElementRef }) expandWrapper: ElementRef;
  @Input() expanded = false;
  @Input() expandHeight = '150px';
  @Input() header = '';
  constructor(public renderer: Renderer2) {}

  ngAfterViewInit() {
    this.renderer.setStyle(this.expandWrapper.nativeElement, 'max-height', this.expandHeight);
  }

    expand() {
        this.expanded = !this.expanded;
    }
}
