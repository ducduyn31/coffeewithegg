import { Component, ViewEncapsulation, NgModule, OnInit } from '@angular/core'
import { CommonModule } from '@angular/common'
import { map, Observable } from 'rxjs'
import { GetAllProjectsGQL, Project } from '@coffeewithegg/data-access'

@Component({
  selector: 'coffeewithegg-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss'],
  encapsulation: ViewEncapsulation.Emulated,
})
export class DashboardComponent implements OnInit {
  projects$: Observable<Project[]> = new Observable<Project[]>()

  constructor(private getAllProjectsGQL: GetAllProjectsGQL) {}

  ngOnInit(): void {
    this.projects$ = this.getAllProjectsGQL
      .watch()
      .valueChanges.pipe(map((result) => result.data.projects))
  }
}

@NgModule({
  imports: [CommonModule],
  declarations: [DashboardComponent],
  exports: [DashboardComponent],
})
export class DashboardComponentModule {}
