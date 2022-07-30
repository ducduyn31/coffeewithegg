import { Component, OnInit } from '@angular/core'
import { NavigationEnd, Router } from '@angular/router'

@Component({
  selector: 'coffeewithegg-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent implements OnInit {
  shouldShowSidebar = false

  constructor(private router: Router) {}

  ngOnInit(): void {
    this.router.events.subscribe((ev) => {
      if (ev instanceof NavigationEnd) {
        this.shouldShowSidebar = !['/', '/dashboard'].includes(ev.url)
      }
    })
  }
}
