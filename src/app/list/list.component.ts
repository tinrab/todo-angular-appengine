import { Component, AfterViewInit } from '@angular/core';
import { Router } from '@angular/router';

import { AuthService } from '../auth.service';
import { User } from '../user.model';

@Component({
  selector: 'app-list',
  templateUrl: './list.component.html'
})
export class ListComponent implements AfterViewInit {

  user: User;

  constructor(
    private router: Router,
    private authService: AuthService
  ) {
    this.user = this.authService.currentUser;
  }

  ngAfterViewInit(): void {
  }

  signOut(): void {
    this.authService.signOut();
    this.router.navigateByUrl('signin');
  }

}
