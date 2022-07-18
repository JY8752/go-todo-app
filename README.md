# go-todo-app
golangでTodoアプリの作成。

```
erDiagram

users ||--o{ articles: ""

users {
  string name
  string email
  integer age
}

articles {
  string title
  text text
}
```