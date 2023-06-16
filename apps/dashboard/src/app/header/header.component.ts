import { Component, OnDestroy, OnInit, ViewEncapsulation } from '@angular/core'
import { RouterService } from '../router.service'
import { Subscription } from 'rxjs'
import { Router } from '@angular/router'

@Component({
  selector: 'coffeewithegg-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css'],
  encapsulation: ViewEncapsulation.Emulated,
})
export class HeaderComponent implements OnInit, OnDestroy {
  private routerServiceSubscription: Subscription | undefined
  currentPage = 'profile'
  pages = [
    { page: 'profile', label: 'Profile' },
    { page: 'projects', label: 'Projects' },
    { page: 'contact', label: 'Contact' },
  ]

  constructor(private routerService: RouterService, private router: Router) {}

  ngOnInit(): void {
    this.routerServiceSubscription = this.routerService
      .onDashboardChange()
      .subscribe(([, page]) => {
        this.currentPage = page === '' ? 'profile' : page
      })
  }

  ngOnDestroy() {
    this.routerServiceSubscription?.unsubscribe()
  }

  async routeTo(page: string) {
    await this.router.navigate([], {
      queryParams: {
        page,
      },
    })
    const el = document.getElementById(`page-${page}`)
    el?.scrollIntoView({ behavior: 'smooth' })
  }
}
