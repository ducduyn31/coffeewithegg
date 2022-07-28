import { Component, OnInit, ViewEncapsulation } from '@angular/core'
import { map, Observable } from 'rxjs'
import { GetAllProjectsGQL, Project } from '@coffeewithegg/data-access'

@Component({
  selector: 'coffeewithegg-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.scss'],
  encapsulation: ViewEncapsulation.Emulated,
})
export class SidebarComponent implements OnInit {
  projects$: Observable<Project[]> = new Observable<Project[]>()

  constructor(private getAllProjects: GetAllProjectsGQL) {}

  ngOnInit(): void {
    this.projects$ = this.getAllProjects
      .watch()
      .valueChanges.pipe(map((result) => result.data.projects))
  }
}
