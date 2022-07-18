import { NgModule } from '@angular/core'
import { CommonModule } from '@angular/common'
import { CardComponent } from './card/card.component'
import { IconComponent } from './icon/icon.component'

@NgModule({
  imports: [CommonModule],
  declarations: [CardComponent, IconComponent],
  exports: [CardComponent, IconComponent],
})
export class CommonAngularModule {}
