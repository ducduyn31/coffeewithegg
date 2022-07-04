import { Component, Input, OnInit, ViewEncapsulation } from '@angular/core';

@Component({
  selector: 'cwe-icon',
  templateUrl: './icon.component.html',
  styleUrls: ['./icon.component.scss'],
  encapsulation: ViewEncapsulation.Emulated,
})
export class IconComponent implements OnInit {
  @Input() name = '';

  constructor() {}

  ngOnInit(): void {
  }
}
