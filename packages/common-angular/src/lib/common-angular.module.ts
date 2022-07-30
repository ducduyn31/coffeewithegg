import { NgModule } from '@angular/core'
import { CommonModule } from '@angular/common'
import { CardComponent } from './card/card.component'
import { IconComponent } from './icon/icon.component'
import { IconButtonComponent } from './icon-button/icon-button.component'

@NgModule({
  imports: [CommonModule],
  declarations: [CardComponent, IconComponent, IconButtonComponent],
  exports: [CardComponent, IconComponent, IconButtonComponent],
})
export class CommonAngularModule {}
