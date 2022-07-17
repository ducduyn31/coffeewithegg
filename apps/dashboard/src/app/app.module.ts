import { NgModule } from '@angular/core'
import { BrowserModule } from '@angular/platform-browser'

import { AppComponent } from './app.component'
import { RouterModule } from '@angular/router'
import { ProjectOverviewComponent } from './project-overview/project-overview.component'
import { CommonAngularModule } from '@coffeewithegg/common-angular'
import { CommonModule } from '@angular/common'
import { AppRoutes } from './app.routes'

@NgModule({
  declarations: [AppComponent, ProjectOverviewComponent],
  imports: [
    BrowserModule,
    CommonModule,
    CommonAngularModule,
    RouterModule.forRoot(AppRoutes, { initialNavigation: 'enabledBlocking' }),
  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
