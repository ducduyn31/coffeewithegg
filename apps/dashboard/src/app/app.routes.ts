import { Routes } from '@angular/router'
import {
  WebComponentWrapper,
  WebComponentWrapperOptions,
} from '@angular-architects/module-federation-tools'
import { environment } from '../environments/environment'

export const AppRoutes: Routes = [
  {
    path: 'ops',
    loadChildren: () => import('ops/Module').then((m) => m.RemoteEntryModule),
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
]
