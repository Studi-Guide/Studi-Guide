import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { IonicModule } from '@ionic/angular';

import { IonicHorizontalScrollableComponent } from './ionic-horizontal-scrollable.component';

describe('IonicHorizontalScrollableComponent', () => {
  let component: IonicHorizontalScrollableComponent;
  let fixture: ComponentFixture<IonicHorizontalScrollableComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ IonicHorizontalScrollableComponent ],
      imports: [IonicModule.forRoot()]
    }).compileComponents();

    fixture = TestBed.createComponent(IonicHorizontalScrollableComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  }));

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
