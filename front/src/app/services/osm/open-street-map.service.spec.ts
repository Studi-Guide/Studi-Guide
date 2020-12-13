import { TestBed } from '@angular/core/testing';

import { OpenStreetMapService } from './open-street-map.service';

describe('OpenStreetMapService', () => {
  let service: OpenStreetMapService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(OpenStreetMapService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
