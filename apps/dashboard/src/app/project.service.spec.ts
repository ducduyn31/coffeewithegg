import { TestBed } from '@angular/core/testing'

import { ProjectService } from './project.service'
import {
  ApolloTestingController,
  ApolloTestingModule,
} from 'apollo-angular/testing'
import { GetAllProjectsGQL } from '@coffeewithegg/data-access'
import { Apollo } from 'apollo-angular'
import { mockProject } from './mocks'

describe('ProjectService', () => {
  let service: ProjectService
  let controller: ApolloTestingController

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [ApolloTestingModule],
      providers: [ProjectService, GetAllProjectsGQL, Apollo],
    })
    controller = TestBed.inject(ApolloTestingController)
    const apollo = TestBed.inject(Apollo)
    service = new ProjectService(new GetAllProjectsGQL(apollo))
  })

  afterEach(() => {
    controller.verify()
  })

  it('should be created', () => {
    expect(service).toBeTruthy()
  })

  it('should return a map of redirectPath to Project', () => {
    service.getAllProjectsMappings().subscribe((result) => {
      expect(result.keys()).toHaveLength(1)
    })

    controller.expectOne('GetAllProjects').flushData({
      projects: [mockProject()],
    })
  })
})
