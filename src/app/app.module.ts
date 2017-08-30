import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { SignInComponent } from './sign-in/sign-in.component';
import { ListComponent } from './list/list.component';
import { AuthService } from './auth.service';
import { TodoService } from './todo.service';
import { TodoComponent } from './todo/todo.component';

@NgModule({
  declarations: [
    AppComponent,
    SignInComponent,
    ListComponent,
    TodoComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    AppRoutingModule
  ],
  providers: [AuthService, TodoService],
  bootstrap: [AppComponent]
})
export class AppModule { }
