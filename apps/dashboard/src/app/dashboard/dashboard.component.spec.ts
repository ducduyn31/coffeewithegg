import { ComponentFixture, TestBed } from '@angular/core/testing'

import { DashboardComponent } from './dashboard.component'
import {
  ApolloTestingController,
  ApolloTestingModule,
} from 'apollo-angular/testing'
import { ProjectService } from '../project.service'
import { Observable, of } from 'rxjs'
import { Project } from '@coffeewithegg/data-access'
import { mockProject } from '../mocks'
import { CommonAngularModule } from '@coffeewithegg/common-angular'
import { ProjectOverviewComponent } from '../project-overview/project-overview.component'
import { RouterTestingModule } from '@angular/router/testing'

describe('DashboardComponent', () => {
  let component: DashboardComponent
  let fixture: ComponentFixture<DashboardComponent>
  let controller: ApolloTestingController
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
      declarations: [DashboardComponent, ProjectOverviewComponent],
      imports: [
        CommonAngularModule,
        ApolloTestingModule,
        RouterTestingModule.withRoutes([]),
      ],
      providers: [
        {
          provide: ProjectService,
          useValue: mockProjectService,
        },
      ],
    }).compileComponents()

    fixture = TestBed.createComponent(DashboardComponent)
    component = fixture.componentInstance
    fixture.detectChanges()

    controller = TestBed.inject(ApolloTestingController)
  })

  afterEach(() => {
    controller.verify()
  })

  it('should create', () => {
    expect(component).toBeTruthy()
  })

  it('should show all projects', () => {
    const projects = fixture.nativeElement.querySelectorAll(
      'coffeewithegg-project-overview',
    )
    expect(projects).toHaveLength(1)
  })

  it('should have an about section', async () => {
    const about = fixture.nativeElement.querySelector('coffeewithegg-about')
    expect(about).toBeTruthy()
  })
})
