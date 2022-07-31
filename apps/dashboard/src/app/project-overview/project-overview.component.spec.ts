import { ComponentFixture, TestBed } from '@angular/core/testing'

import { ProjectOverviewComponent } from './project-overview.component'
import { Project } from '@coffeewithegg/data-access'
import { CommonAngularModule } from '@coffeewithegg/common-angular'
import { RouterTestingModule } from '@angular/router/testing'
import { Router } from '@angular/router'

describe('ProjectOverviewComponent', () => {
  let component: ProjectOverviewComponent
  let fixture: ComponentFixture<ProjectOverviewComponent>
  let router: Router

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ProjectOverviewComponent],
      imports: [CommonAngularModule, RouterTestingModule.withRoutes([])],
    }).compileComponents()

    fixture = TestBed.createComponent(ProjectOverviewComponent)
    component = fixture.componentInstance
    fixture.detectChanges()

    router = TestBed.inject(Router)
  })

  it('should create', () => {
    expect(component).toBeTruthy()
  })

  it('should redirect to the correct path', () => {
    const project: Project = {
      id: '1',
      key: 'test',
      name: 'Test Project',
      technologies: [{ id: '123', key: 'angular', name: 'Angular' }],
    }
    const redirect = 'test'
    component.project = project
    component.redirectPath = redirect
    fixture.detectChanges()

    const mockedFn = jest.fn()
    jest.spyOn(router, 'navigateByUrl').mockImplementation(mockedFn)

    const card = fixture.nativeElement.querySelector('cwe-card')
    card.click()

    expect(mockedFn).toBeCalledWith('test')
  })
})
