import { Component, OnInit, ViewEncapsulation } from '@angular/core'
import { AuthService } from '@auth0/auth0-angular'
import { ActivatedRoute, Router } from '@angular/router'

@Component({
  selector: 'coffeewithegg-ops-entry',
  templateUrl: './entry.component.html',
  encapsulation: ViewEncapsulation.Emulated,
})
export class RemoteEntryComponent implements OnInit {
  isAuthenticated$ = this.auth.isAuthenticated$

  constructor(
    private auth: AuthService,
    private router: Router,
    private route: ActivatedRoute,
  ) {}

  ngOnInit(): void {
    this.auth.isAuthenticated$.subscribe((isAuthenticated) => {
      if (isAuthenticated) {
        const parentCommand = this.route.parent
          ? this.route.parent.snapshot.url.join('/')
          : ''
        this.router.navigate([parentCommand, 'dashboard'])
      }
    })
  }
}
