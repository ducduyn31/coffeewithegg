import { Component, OnDestroy, OnInit } from '@angular/core'
import { Subscription } from 'rxjs'
import { RouterService } from './router.service'

@Component({
  selector: 'coffeewithegg-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent implements OnInit, OnDestroy {
  shouldShowSidebar = false
  private routerSubscription: Subscription | undefined

  constructor(private routerService: RouterService) {}

  ngOnInit(): void {
    this.routerSubscription = this.routerService
      .onDashboardChange()
      .subscribe(([isDashboard]) => {
        this.shouldShowSidebar = !isDashboard
      })
  }

  ngOnDestroy() {
    this.routerSubscription?.unsubscribe()
  }
}
