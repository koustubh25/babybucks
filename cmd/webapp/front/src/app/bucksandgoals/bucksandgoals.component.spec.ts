import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BucksandgoalsComponent } from './bucksandgoals.component';

describe('BucksandgoalsComponent', () => {
  let component: BucksandgoalsComponent;
  let fixture: ComponentFixture<BucksandgoalsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ BucksandgoalsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(BucksandgoalsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
