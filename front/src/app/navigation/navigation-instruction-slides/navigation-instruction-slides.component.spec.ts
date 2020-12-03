import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { IonicModule } from '@ionic/angular';

import { NavigationInstructionSlidesComponent } from './navigation-instruction-slides.component';

describe('NavigationInstructionSlidesComponent', () => {
  let component: NavigationInstructionSlidesComponent;
  let fixture: ComponentFixture<NavigationInstructionSlidesComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ NavigationInstructionSlidesComponent ],
      imports: [IonicModule.forRoot()]
    }).compileComponents();

    fixture = TestBed.createComponent(NavigationInstructionSlidesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  }));

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
