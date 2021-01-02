import { TestBed } from '@angular/core/testing';

import { RecentSearchesService } from './recent-searches.service';

describe('RecentSearchesService', () => {
  let service: RecentSearchesService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(RecentSearchesService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
