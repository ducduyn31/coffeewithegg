import { Component } from '@angular/core';
import { Project } from './app.types';

@Component({
  selector: 'coffeewithegg-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent {
  public projects: Project[] = [
    {
      name: 'Dashboard',
      path: '/',
      description: 'A page that contains list of projects and to navigate among them',
      technologyUsed: ['angular'],
    },
    {
      name: 'Ops',
      path: '/ops',
      description: 'A page that contains list of projects and to navigate among them',
      technologyUsed: ['angular'],
    },
    {
      name: 'Sunny Stream',
      path: '/sunnystream',
      description: 'A page that contains list of projects and to navigate among them',
      technologyUsed: ['react'],
    },
  ]
}
