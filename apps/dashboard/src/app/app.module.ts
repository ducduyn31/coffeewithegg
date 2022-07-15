import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppComponent } from './app.component';
import { RouterModule } from '@angular/router';
import { ProjectOverviewComponent } from './project-overview/project-overview.component';
import { CommonAngularModule } from '@coffeewithegg/common-angular';
import { CommonModule } from '@angular/common';
import { WebComponentWrapper, WebComponentWrapperOptions } from '@angular-architects/module-federation-tools';

@NgModule({
  declarations: [AppComponent, ProjectOverviewComponent],
  imports: [
    BrowserModule,
    CommonModule,
    CommonAngularModule,
    RouterModule.forRoot(
      [
        {
          path: 'ops',
          loadChildren: () =>
            import('ops/Module').then((m) => m.RemoteEntryModule),
        },
        {
          path: 'sunnystream',
          component: WebComponentWrapper,
          data: {
            remoteEntry: 'http://localhost:4202/remoteEntry.js',
            remoteName: 'sunnystream',
            exposedModule: './Module',
            elementName: 'app-sunnystream',
          } as WebComponentWrapperOptions
        }
      ],
      { initialNavigation: 'enabledBlocking' }
    ),
  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
