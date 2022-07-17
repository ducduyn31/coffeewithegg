import { Component, Input, OnInit, ViewEncapsulation } from '@angular/core'
import { Project } from '../app.types'
import { Router } from '@angular/router'

@Component({
  selector: 'coffeewithegg-project-overview',
  templateUrl: './project-overview.component.html',
  styleUrls: ['./project-overview.component.scss'],
  encapsulation: ViewEncapsulation.Emulated,
})
export class ProjectOverviewComponent implements OnInit {
  @Input('project') project!: Project

  constructor(private router: Router) {}

  ngOnInit(): void {}

  onClick() {
    if (this.project.path) this.router.navigateByUrl(this.project.path)
  }
}
