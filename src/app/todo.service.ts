import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/toPromise';

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
    private http: HttpClient
  ) {
    this.user = authService.currentUser;
    this.headers = new HttpHeaders({
      'Accept': 'application/json',
      'Content-Type': 'application/json',
      'Authorization': this.user.sessionToken
    });
  }

  create(title: string): Promise<Todo> {
    return this.http.post<Todo>(
      `${environment.apiUrl}/todos`,
      { title: title },
      { headers: this.headers }
    ).toPromise();
  }

  listTodos(): Promise<Todo[]> {
    return this.http.get<Todo[]>(
      `${environment.apiUrl}/todos`,
      { headers: this.headers }
    ).toPromise();
  }

  update(id: string, title: string): Promise<Todo> {
    return this.http.put<Todo>(
      `${environment.apiUrl}/todos/${id}`,
      { title: title },
      { headers: this.headers }
    ).toPromise();
  }

  delete(id: string): Promise<void> {
    return this.http.delete<Todo>(
      `${environment.apiUrl}/todos/${id}`,
      { headers: this.headers }
    ).map(_ => { })
      .toPromise();
  }

}
