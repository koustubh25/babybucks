import { TestBed } from '@angular/core/testing';

import { BucksandgoalsService } from './bucksandgoals.service';

describe('BucksandgoalsService', () => {
  let service: BucksandgoalsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(BucksandgoalsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
