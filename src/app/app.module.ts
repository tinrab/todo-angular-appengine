import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';

import {
  MdButtonModule,
  MdInputModule,
  MdListModule,
  MdIconModule,
  MdToolbarModule
} from '@angular/material';

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
    BrowserAnimationsModule,
    FormsModule,
    HttpClientModule,
    MdButtonModule,
    MdInputModule,
    MdListModule,
    MdIconModule,
    MdToolbarModule,
    AppRoutingModule
  ],
  providers: [AuthService, TodoService],
  bootstrap: [AppComponent]
})
export class AppModule { }
