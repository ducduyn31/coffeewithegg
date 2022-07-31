import {
  Component,
  EventEmitter,
  Input,
  Output,
  ViewEncapsulation,
} from '@angular/core'

type ButtonVariant = 'default' | 'transparent'

@Component({
  selector: 'cwe-icon-button',
  templateUrl: './icon-button.component.html',
  styleUrls: ['./icon-button.component.scss'],
  encapsulation: ViewEncapsulation.Emulated,
})
export class IconButtonComponent {
  @Input() name = ''
  @Input() variant: ButtonVariant = 'default'
  @Output() btnClick = new EventEmitter<MouseEvent>()

  onClick(event: MouseEvent) {
    this.btnClick.emit(event)
  }
}
