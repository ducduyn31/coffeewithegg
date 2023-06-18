import { Component, OnInit, ViewEncapsulation } from '@angular/core'
import { AuthService } from '@auth0/auth0-angular'
import { map, Observable } from 'rxjs'

@Component({
  selector: 'coffeewithegg-ops-dashboard',
  templateUrl: './ops-dashboard.component.html',
  styleUrls: ['./ops-dashboard.component.scss'],
  encapsulation: ViewEncapsulation.Emulated,
})
export class OpsDashboardComponent implements OnInit {
  protected isAdmin$: Observable<boolean> = new Observable<boolean>()

  constructor(protected auth: AuthService) {}

  ngOnInit() {
    this.isAdmin$ = this.auth.user$.pipe(
      map((user) => {
        return user?.['user_roles']?.includes('admin')
      }),
    )
  }

  logout() {
    this.auth.logout()
  }
}
