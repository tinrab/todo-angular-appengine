# Todo List

<p align="center">
  <img src="https://github.com/outcrawl/site/blob/master/data/posts/todo-list-angular-google-app-engine-part-1/finished.gif"/>
</p>

A short tutorial series about building a To-Do List using Angular and Google App Engine.

1. [Build a Todo List with Angular and Google App Engine - Part 1](https://outcrawl.com/todo-list-angular-google-app-engine-part-1)
2. [Build a Todo List with Angular and Google App Engine - Part 2](https://outcrawl.com/todo-list-angular-google-app-engine-part-2)

## Build

Replace `[CLIENT_ID]` with your own Google client ID inside `server/app.yaml`, `src/environments/environment.ts` and `src/environments/environment.prod.ts`.

Run the following commands to build and deploy.

```
export PROJECT_ID=[PROJECT_ID]
npm run build && npm run deploy
```

Replace `[PROJECT_ID]` with your Google API Console project ID.

Finally, navigate to `https://[PROJECT_ID].appspot.com`.
