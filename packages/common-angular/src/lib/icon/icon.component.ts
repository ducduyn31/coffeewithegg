import { Component, Input, ViewEncapsulation } from '@angular/core'

@Component({
  selector: 'cwe-icon',
  templateUrl: './icon.component.html',
  styleUrls: ['./icon.component.scss'],
  encapsulation: ViewEncapsulation.Emulated,
})
export class IconComponent {
  @Input() name = ''
}
