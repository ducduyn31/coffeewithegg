import { Component, ElementRef, OnInit, ViewChild, ViewEncapsulation } from '@angular/core';

@Component({
  selector: 'cwe-card',
  templateUrl: './card.component.html',
  styleUrls: ['./card.component.scss'],
  encapsulation: ViewEncapsulation.Emulated,
})
export class CardComponent implements OnInit {
  @ViewChild('card', {read: ElementRef}) cardElementRef: ElementRef | undefined;

  constructor(
    private elementRef: ElementRef
  ) {}

  ngOnInit(): void {
    if (!this.cardElementRef) return;
    const attributes = this.elementRef.nativeElement.attributes;
    const cardAttributes = this.cardElementRef.nativeElement.attributes;

    for (const attribute of attributes) {
      if (attribute.name === 'ngModel' || cardAttributes.getNamedItemNS(attribute.namespaceURI, attribute.name)) {
        continue;
      }
      this.cardElementRef.nativeElement.setAttributeNS(attribute.namespaceURI, attribute.name, attribute.value);
    }
  }
}
