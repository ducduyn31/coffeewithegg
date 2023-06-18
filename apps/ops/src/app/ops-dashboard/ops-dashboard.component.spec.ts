import { ComponentFixture, TestBed } from '@angular/core/testing'
import { OpsDashboardComponent } from './ops-dashboard.component'
import { of } from 'rxjs'
import { AuthService } from '@auth0/auth0-angular'

describe('OpsDashboardComponent', () => {
  let component: OpsDashboardComponent
  let fixture: ComponentFixture<OpsDashboardComponent>
  const mockedAuthService = {
    loginWithRedirect: jest.fn(),
    isAuthenticated$: of(true),
    user$: of({
      name: 'Test User',
      picture: 'https://picsum.photos/200',
      email: 'test@example.com',
    }),
  }

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [OpsDashboardComponent],
      providers: [
        {
          provide: AuthService,
          useValue: mockedAuthService,
        },
      ],
    }).compileComponents()

    fixture = TestBed.createComponent(OpsDashboardComponent)
    component = fixture.componentInstance
    fixture.detectChanges()
  })

  it('should create', () => {
    expect(component).toBeTruthy()
  })

  it('should have logout button', async () => {
    const compiled = fixture.nativeElement
    expect(compiled.querySelector('button').textContent).toContain('Logout')
  })
})

describe('OpsDashboardComponent with admin', () => {
  let fixture: ComponentFixture<OpsDashboardComponent>
  const mockedAuthService = {
    loginWithRedirect: jest.fn(),
    isAuthenticated$: of(true),
    user$: of({
      name: 'Test User',
      picture: 'https://picsum.photos/200',
      email: 'admin@example.com',
      user_roles: ['admin'],
    }),
  }

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [OpsDashboardComponent],
      providers: [
        {
          provide: AuthService,
          useValue: mockedAuthService,
        },
      ],
    }).compileComponents()

    fixture = TestBed.createComponent(OpsDashboardComponent)
    fixture.detectChanges()
  })

  it('should have projects section', async () => {
    const compiled = fixture.nativeElement
    expect(compiled.querySelector('#projects').textContent).toContain(
      'Projects',
    )
  })
})
