import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { SignInComponent } from './sign-in/sign-in.component';
import { ListComponent } from './list/list.component';
import { AuthGuard } from './auth.guard';

const routes: Routes = [
  {
    path: '',
    children: [{
      path: 'signin',
      component: SignInComponent
    }, {
      path: '',
      component: ListComponent,
      canActivate: [AuthGuard]
    }]
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
  providers: [AuthGuard]
})
export class AppRoutingModule { }
