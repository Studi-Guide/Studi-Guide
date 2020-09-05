import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { IonicModule } from '@ionic/angular';

import { DetailViewSelectLanguageComponent } from './detail-view-select-language.component';

describe('DetailViewSelectLanguageComponent', () => {
  let component: DetailViewSelectLanguageComponent;
  let fixture: ComponentFixture<DetailViewSelectLanguageComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ DetailViewSelectLanguageComponent ],
      imports: [IonicModule.forRoot()]
    }).compileComponents();

    fixture = TestBed.createComponent(DetailViewSelectLanguageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  }));

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
