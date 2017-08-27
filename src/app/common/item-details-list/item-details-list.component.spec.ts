import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ItemDetailsListComponent } from './item-details-list.component';

describe('ItemDetailsListComponent', () => {
  let component: ItemDetailsListComponent;
  let fixture: ComponentFixture<ItemDetailsListComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ItemDetailsListComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ItemDetailsListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
