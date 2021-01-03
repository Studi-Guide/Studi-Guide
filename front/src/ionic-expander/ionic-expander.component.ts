import {
  AfterViewInit,
  Component,
  ElementRef,
  Input,
  OnChanges,
  Renderer2,
  SimpleChanges,
  ViewChild
} from '@angular/core';

@Component({
  selector: 'app-expander',
  templateUrl: './ionic-expander.component.html',
  styleUrls: ['./ionic-expander.component.scss'],
})
export class IonicExpanderComponent implements OnChanges {
  @ViewChild('expandWrapper', { read: ElementRef }) expandWrapper: ElementRef;
  @ViewChild('expandContent', { read: ElementRef }) expandContent: ElementRef;
  @Input() expanded = false;
  @Input() header = '';
  constructor(public renderer: Renderer2) {
  }

  expand() {
      this.expanded = !this.expanded;
  }

  ngOnChanges(changes: SimpleChanges): void {
    if (changes.expanded.currentValue &&
        (changes.expanded.currentValue !== changes.expanded.previousValue)) {
      const rect = (this.expandContent.nativeElement as HTMLDivElement).getBoundingClientRect();
      this.renderer.setStyle(this.expandWrapper.nativeElement, 'max-height', rect.height);
      this.renderer.setStyle(this.expandWrapper.nativeElement, 'transition', 'max-height 0.4s ease-in-out');
    }
  }
}
