import { NgModule } from '@angular/core'
import { BrowserModule } from '@angular/platform-browser'

import { AppComponent } from './app.component'
import { RouterModule } from '@angular/router'
import { ProjectOverviewComponent } from './project-overview/project-overview.component'
import { CommonAngularModule } from '@coffeewithegg/common-angular'
import { CommonModule } from '@angular/common'
import { AppRoutes } from './app.routes'
import { SidebarComponent } from './sidebar/sidebar.component'
import { DashboardComponentModule } from './dashboard/dashboard.component'
import { APOLLO_OPTIONS, ApolloModule } from 'apollo-angular'
import { HttpLink } from 'apollo-angular/http'
import { InMemoryCache } from '@apollo/client'
import { environment } from '../environments/environment'
import { HttpClientModule } from '@angular/common/http'

@NgModule({
  declarations: [AppComponent, ProjectOverviewComponent, SidebarComponent],
  imports: [
    ApolloModule,
    BrowserModule,
    CommonModule,
    CommonAngularModule,
    DashboardComponentModule,
    HttpClientModule,
    RouterModule.forRoot(AppRoutes, { initialNavigation: 'enabledBlocking' }),
  ],
  providers: [
    {
      provide: APOLLO_OPTIONS,
      useFactory: (httpLink: HttpLink) => {
        return {
          cache: new InMemoryCache(),
          link: httpLink.create({
            uri: environment.backendServer,
          }),
        }
      },
      deps: [HttpLink],
    },
  ],
  bootstrap: [AppComponent],
})
export class AppModule {}
