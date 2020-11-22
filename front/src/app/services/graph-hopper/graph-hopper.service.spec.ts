import { TestBed } from '@angular/core/testing';

import { GraphHopperService } from './graph-hopper.service';

describe('GraphHopperService', () => {
  let service: GraphHopperService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GraphHopperService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
