import { Injectable } from '@angular/core'
import { GetAllProjectsGQL, Project } from '@coffeewithegg/data-access'
import { map, Observable } from 'rxjs'

@Injectable({
  providedIn: 'root',
})
export class ProjectService {
  constructor(private getAllProjectsGQL: GetAllProjectsGQL) {}

  public getAllProjectsMappings(): Observable<Map<string, Project>> {
    return this.getAllProjectsGQL.watch().valueChanges.pipe(
      map((result) => {
        const resultMap = new Map()
        if (!result?.data?.projects) return resultMap

        result.data.projects.forEach((project) => {
          resultMap.set(encodeURI(project.key), project)
        })

        return resultMap
      }),
    )
  }
}
