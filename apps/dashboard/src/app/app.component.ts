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
    },
    {
      name: 'Ops'
    },
  ]
}
