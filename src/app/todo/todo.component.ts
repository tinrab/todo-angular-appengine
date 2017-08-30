import {
  Component,
  Input,
  Output,
  EventEmitter
} from '@angular/core';

import { Todo } from '../todo.model';

@Component({
  selector: 'app-todo',
  templateUrl: './todo.component.html'
})
export class TodoComponent {

  @Input()
  item: Todo;
  @Output()
  updateRequest: EventEmitter<string>;
  @Output()
  deleteRequest: EventEmitter<void>;
  editMode: boolean;
  editableTitle: string;

  constructor() {
    this.updateRequest = new EventEmitter<string>();
    this.deleteRequest = new EventEmitter<void>();
  }

  update(): void {
    if (this.editMode) {
      this.editMode = false;
      this.updateRequest.emit(this.editableTitle);
    }
  }

  delete(): void {
    this.deleteRequest.emit();
  }

  onTitleClick(): void {
    this.editMode = true;
    this.editableTitle = this.item.title;
  }

  onTitleInputBlur(): void {
    this.editMode = false;
  }

}
