import { Component, OnInit, ViewEncapsulation } from '@angular/core'
import { Project } from '@coffeewithegg/data-access'
import { Router } from '@angular/router'
import { ProjectService } from '../singletons/project/project.service'

@Component({
  selector: 'coffeewithegg-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.scss'],
  encapsulation: ViewEncapsulation.Emulated,
})
export class SidebarComponent implements OnInit {
  projectsMapper: Map<string, Project> = new Map()
  expand = false

  constructor(private projectService: ProjectService, private router: Router) {}

  onClick(target: string) {
    this.router.navigateByUrl(target)
  }

  toggleExpand() {
    this.expand = !this.expand
  }

  ngOnInit(): void {
    this.projectService.getAllProjectsMappings().subscribe((result) => {
      this.projectsMapper = result
    })
  }
}
