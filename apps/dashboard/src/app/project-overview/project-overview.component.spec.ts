import { ComponentFixture, TestBed } from '@angular/core/testing'

import { ProjectOverviewComponent } from './project-overview.component'
import { RouterModule } from '@angular/router'
import { APP_BASE_HREF } from '@angular/common'

describe('ProjectOverviewComponent', () => {
  let component: ProjectOverviewComponent
  let fixture: ComponentFixture<ProjectOverviewComponent>

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ProjectOverviewComponent],
      imports: [RouterModule.forRoot([])],
      providers: [{ provide: APP_BASE_HREF, useValue: '/' }],
    }).compileComponents()

    fixture = TestBed.createComponent(ProjectOverviewComponent)
    component = fixture.componentInstance
    fixture.detectChanges()
  })

  it('should create', () => {
    expect(component).toBeTruthy()
  })
})
