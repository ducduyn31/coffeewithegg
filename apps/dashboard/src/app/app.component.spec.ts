import { TestBed } from '@angular/core/testing'
import { AppComponent } from './app.component'
import { APP_BASE_HREF } from '@angular/common'
import { RouterModule } from '@angular/router'
import { RouterTestingModule } from '@angular/router/testing'

describe('AppComponent', () => {
  it('should create the app', () => {
    const fixture = TestBed.configureTestingModule({
      imports: [RouterTestingModule.withRoutes([])],
    }).createComponent(AppComponent)
    const app = fixture.componentInstance
    expect(app).toBeTruthy()
  })

  it('should not show sidebar when path is / or /dashboard', () => {
    const fixture = TestBed.configureTestingModule({
      imports: [RouterModule.forRoot([])],
      providers: [{ provide: APP_BASE_HREF, useValue: '/' }],
    }).createComponent(AppComponent)
    const app = fixture.componentInstance
    expect(app.shouldShowSidebar).toEqual(false)
  })
})
