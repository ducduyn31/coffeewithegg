import { Component, Input, ViewEncapsulation } from '@angular/core'
import { Router } from '@angular/router'
import { Project } from '@coffeewithegg/data-access'

@Component({
  selector: 'coffeewithegg-project-overview',
  templateUrl: './project-overview.component.html',
  styleUrls: ['./project-overview.component.scss'],
  encapsulation: ViewEncapsulation.Emulated,
})
export class ProjectOverviewComponent {
  @Input() project!: Project
  @Input() redirectPath!: string

  constructor(private router: Router) {}

  onClick() {
    if (this.project && this.redirectPath)
      this.router.navigateByUrl(this.redirectPath)
  }
}
