import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { environment } from '../environments/environment';
import { User } from './user.model';

declare const gapi: any;

interface SignInResponse {
  userId: string;
  sessionToken: string;
}

@Injectable()
export class AuthService {

  private googleAuth: any;
  private user: User;

  constructor(
    private http: HttpClient
  ) {
    gapi.load('client:auth2', () => {
      gapi.client.init({
        'clientId': environment.clientId,
        'scope': 'profile'
      }).then(() => {
        this.googleAuth = gapi.auth2.getAuthInstance();
        const googleUser = this.googleAuth.currentUser.get();
        // Get user's data if he's signed in
        if (googleUser) {
          this.user = JSON.parse(localStorage.getItem('user'));
        }
      });
    });
  }

  signIn(): Promise<User> {
    return new Promise<User>((resolve, reject) => {
      this.googleAuth.signIn({
        'prompt': 'consent'
      }).then(googleUser => {
        const token = googleUser.getAuthResponse().id_token;
        this.http.post<SignInResponse>(`${environment.apiUrl}/signin`, null, {
          headers: new HttpHeaders().set('Authorization', token)
        }).subscribe(res => {
          const profile = googleUser.getBasicProfile();
          this.user = new User(res.userId, res.sessionToken, profile.getName());
          localStorage.setItem('user', JSON.stringify(this.user));
          resolve(this.user);
        }, reject);
      }, reject);
    });
  }

  signOut(): void {
    localStorage.removeItem('user');
    this.user = null;
    this.googleAuth.signOut();
  }

  get currentUser(): User {
    if (!this.user) {
      this.user = JSON.parse(localStorage.getItem('user'));
    }
    return this.user;
  }

  get isSignedIn(): boolean {
    return this.currentUser != null;
  }

}
