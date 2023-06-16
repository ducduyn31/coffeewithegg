import { setRemoteDefinitions } from '@nx/angular/mf'
import { environment } from './environments/environment'

setRemoteDefinitions({
  ops: environment.opsEntry,
})

import('./bootstrap').catch((err) => console.error(err))
