import { Injectable } from '@angular/core'
import { NavigationEnd, Router } from '@angular/router'
import { filter, map, Observable } from 'rxjs'

@Injectable({
  providedIn: 'root',
})
export class RouterService {
  private isDashboard$: Observable<[boolean, string]> | undefined

  constructor(private readonly router: Router) {}

  public onDashboardChange(): Observable<[boolean, string]> {
    if (!this.isDashboard$) {
      this.isDashboard$ = this.router.events.pipe(
        filter((ev) => ev instanceof NavigationEnd),
        map((ev) => {
          const url = this.router.parseUrl((ev as NavigationEnd).url)
          const isDashboard =
            !url.root.hasChildren() ||
            url.root.children['primary'].segments[0].path === 'dashboard'
          const { page } = url.queryParams
          return [isDashboard, page]
        }),
      )
    }

    return this.isDashboard$
  }
}
