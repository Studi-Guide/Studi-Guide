import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { IonicModule } from '@ionic/angular';

import { IonicBottomDrawerComponent } from './ionic-bottom-drawer.component';

describe('IonicBottomDrawerComponent', () => {
  let component: IonicBottomDrawerComponent;
  let fixture: ComponentFixture<IonicBottomDrawerComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ IonicBottomDrawerComponent ],
      imports: [IonicModule.forRoot()]
    }).compileComponents();

    fixture = TestBed.createComponent(IonicBottomDrawerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  }));

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
