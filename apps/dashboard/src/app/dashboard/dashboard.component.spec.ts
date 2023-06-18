import { ComponentFixture, TestBed } from '@angular/core/testing'

import { DashboardComponent } from './dashboard.component'
import {
  ApolloTestingController,
  ApolloTestingModule,
} from 'apollo-angular/testing'
import { ProjectService } from '../singletons/project/project.service'
import { Observable, of } from 'rxjs'
import { Project } from '@coffeewithegg/data-access'
import { mockProject } from '../mocks'
import { CommonAngularModule } from '@coffeewithegg/common-angular'
import { ProjectOverviewComponent } from '../project-overview/project-overview.component'
import { RouterTestingModule } from '@angular/router/testing'
import { HeaderComponent } from '../header/header.component'
import { ProfileComponent } from '../profile/profile.component'
import { ContactComponent } from '../contact/contact.component'

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
      declarations: [
        DashboardComponent,
        ProjectOverviewComponent,
        HeaderComponent,
        ProfileComponent,
        ContactComponent,
      ],
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

  it('should have a header and a contact', async () => {
    const header = fixture.nativeElement.querySelector('coffeewithegg-header')
    const profile = fixture.nativeElement.querySelector('coffeewithegg-profile')
    const contact = fixture.nativeElement.querySelector('coffeewithegg-contact')
    expect(header).toBeTruthy()
    expect(profile).toBeTruthy()
    expect(contact).toBeTruthy()
  })
})
