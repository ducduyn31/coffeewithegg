import { ComponentFixture, TestBed } from '@angular/core/testing'

import { SidebarComponent } from './sidebar.component'
import { RouterTestingModule } from '@angular/router/testing'
import { ProjectService } from '../project.service'
import { Observable, of } from 'rxjs'
import { mockProject } from '../mocks'
import { Project } from '@coffeewithegg/data-access'
import { CommonAngularModule } from '@coffeewithegg/common-angular'
import { Router } from '@angular/router'

describe('SidebarComponent', () => {
  let component: SidebarComponent
  let fixture: ComponentFixture<SidebarComponent>
  let router: Router
  let mockedGetAllProjects: jest.MockedFn<
    () => Observable<Map<string, Project>>
  >

  beforeEach(async () => {
    mockedGetAllProjects = jest
      .fn()
      .mockImplementation(() => of(new Map([['name', mockProject()]])))
    const mockProjectService = {
      getAllProjectsMappings: mockedGetAllProjects,
    }
    await TestBed.configureTestingModule({
      imports: [CommonAngularModule, RouterTestingModule],
      declarations: [SidebarComponent],
      providers: [
        {
          provide: ProjectService,
          useValue: mockProjectService,
        },
      ],
    }).compileComponents()

    fixture = TestBed.createComponent(SidebarComponent)
    component = fixture.componentInstance
    fixture.detectChanges()

    router = TestBed.inject(Router)
  })

  it('should create', () => {
    expect(component).toBeTruthy()
    expect(fixture.nativeElement.querySelectorAll('li')).toHaveLength(1)
  })

  it('should navigate to the correct path when click', () => {
    const mockedFn = jest.fn()
    jest.spyOn(router, 'navigateByUrl').mockImplementation(mockedFn)

    const firstPath = fixture.nativeElement
      .querySelector('li')
      .querySelector('button')
    firstPath.click()

    expect(mockedFn).toBeCalledWith('name')
  })

  it('should expand and collapse', () => {
    const toggler = fixture.nativeElement
      .querySelector('cwe-icon-button')
      .querySelector('button')
    expect(component.expand).toEqual(false)
    toggler.click()
    expect(component.expand).toEqual(true)
  })
})
