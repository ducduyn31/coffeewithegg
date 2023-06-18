import { NgModule } from '@angular/core'
import { CommonModule } from '@angular/common'
import { OpsDashboardComponent } from './ops-dashboard.component'
import { RouterModule } from '@angular/router'

@NgModule({
  imports: [
    CommonModule,
    RouterModule.forChild([
      {
        path: '',
        component: OpsDashboardComponent,
      },
    ]),
  ],
  declarations: [OpsDashboardComponent],
  exports: [OpsDashboardComponent],
})
export class OpsDashboardComponentModule {}
