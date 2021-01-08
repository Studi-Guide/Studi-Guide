import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { IonicModule } from '@ionic/angular';

import { DetailViewSelectAppearanceComponent } from './detail-view-select-appearance.component';

describe('DetailViewSelectAppearanceComponent', () => {
  let component: DetailViewSelectAppearanceComponent;
  let fixture: ComponentFixture<DetailViewSelectAppearanceComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ DetailViewSelectAppearanceComponent ],
      imports: [IonicModule.forRoot()]
    }).compileComponents();

    fixture = TestBed.createComponent(DetailViewSelectAppearanceComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  }));

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
