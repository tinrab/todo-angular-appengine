import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../auth.service';

@Component({
  selector: 'app-sign-in',
  templateUrl: './sign-in.component.html'
})
export class SignInComponent {

  constructor(
    private router: Router,
    private authService: AuthService
  ) { }

  signIn(): void {
    this.authService.signIn()
      .then(user => this.router.navigateByUrl(''))
      .catch(error => console.log(error));
  }

}
