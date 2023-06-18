import { NgModule } from '@angular/core'
import { CommonModule } from '@angular/common'
import { RouterModule } from '@angular/router'

import { RemoteEntryComponent } from './entry.component'
import { LoginComponent } from '../login/login.component'
import { AuthGuard } from '@auth0/auth0-angular'

@NgModule({
  declarations: [RemoteEntryComponent, LoginComponent],
  imports: [
    CommonModule,
    RouterModule.forChild([
      {
        path: '',
        component: RemoteEntryComponent,
        pathMatch: 'full',
      },
      {
        path: 'dashboard',
        canActivate: [AuthGuard],
        loadChildren: () =>
          import('../ops-dashboard/ops-dashboard.module').then(
            (m) => m.OpsDashboardComponentModule,
          ),
      },
    ]),
  ],
  providers: [],
})
export class RemoteEntryModule {}
