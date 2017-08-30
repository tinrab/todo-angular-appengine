import {
  Directive,
  AfterViewInit,
  ElementRef,
  Renderer
} from '@angular/core';

@Directive({
  selector: '[appAutofocus]'
})
export class AutofocusDirective implements AfterViewInit {

  constructor(
    private element: ElementRef,
    private renderer: Renderer
  ) { }

  ngAfterViewInit(): void {
    this.renderer.invokeElementMethod(
      this.element.nativeElement,
      'focus',
      []
    );
  }

}
