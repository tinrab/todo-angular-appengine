import { Component, AfterViewInit } from '@angular/core';
import { Router } from '@angular/router';

import { AuthService } from '../auth.service';
import { User } from '../user.model';
import { TodoService } from '../todo.service';
import { Todo } from '../todo.model';

@Component({
  selector: 'app-list',
  templateUrl: './list.component.html'
})
export class ListComponent implements AfterViewInit {

  user: User;
  todos: Todo[];
  newTodoTitle: string;

  constructor(
    private router: Router,
    private authService: AuthService,
    private todoService: TodoService
  ) {
    this.user = this.authService.currentUser;
  }

  ngAfterViewInit(): void {
    this.todoService.listTodos()
    .then(todos => this.todos = todos);
  }

  signOut(): void {
    this.authService.signOut();
    this.router.navigateByUrl('signin');
  }

  createTodo(): void {
    const title = this.newTodoTitle.trim();
    if (title) {
      this.todoService.createTodo(title)
      .then(todo => this.todos.unshift(todo));
      this.newTodoTitle = '';
    }
  }

}
