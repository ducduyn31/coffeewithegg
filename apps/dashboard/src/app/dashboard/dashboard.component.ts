import { Component, NgModule, OnInit, ViewEncapsulation } from '@angular/core'
import { CommonModule } from '@angular/common'
import { Project } from '@coffeewithegg/data-access'
import { ProjectService } from '../singletons/project/project.service'
import { ProjectOverviewComponent } from '../project-overview/project-overview.component'
import { CommonAngularModule } from '@coffeewithegg/common-angular'
import { ProfileComponent } from '../profile/profile.component'
import { HeaderComponent } from '../header/header.component'
import { ContactComponent } from '../contact/contact.component'
import { RouterModule } from '@angular/router'
import { SharedModule } from '../shared/shared.module'

@Component({
  selector: 'coffeewithegg-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss'],
  encapsulation: ViewEncapsulation.Emulated,
})
export class DashboardComponent implements OnInit {
  projects: Map<string, Project> = new Map<string, Project>()

  constructor(private projectService: ProjectService) {}

  ngOnInit(): void {
    this.projectService
      .getAllProjectsMappings()
      .subscribe((result) => (this.projects = result))
  }
}

@NgModule({
  imports: [CommonModule, CommonAngularModule, RouterModule, SharedModule],
  declarations: [
    DashboardComponent,
    ProjectOverviewComponent,
    ProfileComponent,
    HeaderComponent,
    ContactComponent,
  ],
  exports: [DashboardComponent],
})
export class DashboardComponentModule {}
