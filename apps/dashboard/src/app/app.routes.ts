import { Routes } from '@angular/router'
import {
  WebComponentWrapper,
  WebComponentWrapperOptions,
} from '@angular-architects/module-federation-tools'
import { environment } from '../environments/environment'
import { DashboardComponent } from './dashboard/dashboard.component'
import { loadRemoteModule } from '@nx/angular/mf'

export const AppRoutes: Routes = [
  {
    path: 'ops',
    loadChildren: () =>
      loadRemoteModule('ops', './Module').then((m) => m.RemoteEntryModule),
  },
  {
    path: 'sunnystream',
    component: WebComponentWrapper,
    data: {
      remoteEntry: environment.sunnyStreamEntry,
      remoteName: 'sunnystream',
      exposedModule: './Module',
      elementName: 'app-sunnystream',
    } as WebComponentWrapperOptions,
  },
  {
    path: 'dashboard',
    component: DashboardComponent,
  },
  {
    path: '',
    component: DashboardComponent,
  },
]
