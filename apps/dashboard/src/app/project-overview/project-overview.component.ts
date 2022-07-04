import { Component, Input, OnInit, ViewEncapsulation } from '@angular/core';
import { Project } from '../app.types';

@Component({
  selector: 'coffeewithegg-project-overview',
  templateUrl: './project-overview.component.html',
  styleUrls: ['./project-overview.component.scss'],
  encapsulation: ViewEncapsulation.Emulated,
})
export class ProjectOverviewComponent implements OnInit {
  @Input('project') project!: Project;

  constructor() {}

  ngOnInit(): void {}
}
