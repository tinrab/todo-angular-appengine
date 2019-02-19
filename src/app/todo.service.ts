import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { map } from 'rxjs/operators';

import { environment } from '../environments/environment';
import { Todo } from './todo.model';
import { User } from './user.model';
import { AuthService } from './auth.service';

@Injectable()
export class TodoService {

  private user: User;
  private headers: HttpHeaders;

  constructor(
    private authService: AuthService,
    private http: HttpClient,
  ) {
    this.user = authService.currentUser;
    this.headers = new HttpHeaders({
      'Accept': 'application/json',
      'Content-Type': 'application/json',
    });

    if (this.user) {
      this.headers.set('Authorization', this.user.sessionToken);
    }
  }

  createTodo(title: string): Promise<Todo> {
    return this.http.post<Todo>(
      `${environment.apiUrl}/todos`,
      { title },
      { headers: this.headers },
    ).toPromise();
  }

  listTodos(): Promise<Todo[]> {
    return this.http.get<Todo[]>(
      `${environment.apiUrl}/todos`,
      { headers: this.headers },
    ).toPromise();
  }

  updateTodo(id: string, title: string): Promise<Todo> {
    return this.http.post<Todo>(
      `${environment.apiUrl}/todos/${id}`,
      { title },
      { headers: this.headers },
    ).toPromise();
  }

  deleteTodo(id: string): Promise<void> {
    return this.http.delete<Todo>(
      `${environment.apiUrl}/todos/${id}`,
      { headers: this.headers },
    ).pipe(
      map(() => {
      }),
    ).toPromise();
  }

}
