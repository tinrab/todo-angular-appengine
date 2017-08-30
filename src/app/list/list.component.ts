import { Component, AfterViewInit } from '@angular/core';
import { Router } from '@angular/router';

import { AuthService } from '../auth.service';
import { User } from '../user.model';
import { TodoService } from '../todo.service';

@Component({
  selector: 'app-list',
  templateUrl: './list.component.html'
})
export class ListComponent implements AfterViewInit {

  user: User;

  constructor(
    private router: Router,
    private authService: AuthService,
    private todoService: TodoService
  ) {
    this.user = this.authService.currentUser;
  }

  ngAfterViewInit(): void {
    this.todoService.listTodos()
    .then(todos => console.log(todos));
  }

  signOut(): void {
    this.authService.signOut();
    this.router.navigateByUrl('signin');
  }

}
