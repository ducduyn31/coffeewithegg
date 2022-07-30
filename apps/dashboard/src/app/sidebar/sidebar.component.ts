import { Component, OnInit, ViewEncapsulation } from '@angular/core'
import { GetAllProjectsGQL, Project } from '@coffeewithegg/data-access'
import { Router } from '@angular/router'

@Component({
  selector: 'coffeewithegg-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.scss'],
  encapsulation: ViewEncapsulation.Emulated,
})
export class SidebarComponent implements OnInit {
  projects: Project[] = []
  expand = true

  constructor(
    private getAllProjects: GetAllProjectsGQL,
    private router: Router,
  ) {}

  onClick(target: string) {
    this.router.navigateByUrl(target)
  }

  toggleExpand() {
    this.expand = !this.expand
  }

  ngOnInit(): void {
    this.getAllProjects.watch().valueChanges.subscribe((result) => {
      this.projects = result.data.projects
    })
  }
}
