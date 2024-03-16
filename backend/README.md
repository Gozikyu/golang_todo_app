# Todo App API ドキュメント

## ログイン API

### 概要

ログインを行う

### エンドポイント

POST /login

### サンプルリクエスト

```bash
curl -X POST \
  http://localhost:8888/login \
  -H 'Content-Type: application/json' \
  -d '{"email":"ichiro@example.com","password":"password"}'
```

## ユーザータスク取得 API

### 概要

特定のユーザーの全タスクを取得

### エンドポイント

GET /:userId/tasks

### サンプルリクエスト

```bash
curl http://localhost:8888/1/tasks
```

## ユーザータスク作成 API

### 概要

新規タスクを作成

### エンドポイント

POST /:userId/tasks

### サンプルリクエスト

```bash
curl -X POST \
 http://localhost:8888/restricted/1/tasks \
 -H  'Content-Type: application/json' \
 -H  'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxIiwiZXhwIjoxNzEwNTkyNTA5fQ.Kj-UjyIuAXitKFBd_XK4mSF1t9Z-KR3yx311D5--P3A' \
 -d '{
"UserId": "1",
"Title": "Sample Task",
"Description": "これはサンプルタスクです。",
"Status": "NOT_STARTED"
}'
```

## ユーザータスク削除 API

### 概要

タスクを削除

### エンドポイント

DELETE /:userId/tasks/:taskId

### サンプルリクエスト

```bash
curl -X DELETE http://localhost:8888/1/tasks/101
```

## ユーザータスク更新 API

### 概要

タスク内容を更新

### エンドポイント

PUT /:userId/tasks/:taskId

### サンプルリクエスト

```bash
curl -X PUT \
  http://localhost:8888/your_user_id/tasks/your_task_id \
  -H 'Content-Type: application/json' \
  -d '{
    "TaskId": "102",
    "UserId": "2",
    "Title": "Updated Task",
    "Description": "タスクが更新されました。",
    "Status": "IN_PROGRESS"
  }'
```

## 全ユーザー取得 API

### 概要

全ユーザーを取得

### エンドポイント

GET /users

### サンプルリクエスト

```bash
curl http://localhost:8888/users
```

## ユーザー取得 API

### 概要

特定ユーザーを取得

### エンドポイント

GET /users/:userId

### サンプルリクエスト

```bash
curl http://localhost:8888/users/1
```

## ユーザー作成 API

### 概要

ユーザーを作成

### エンドポイント

POST /users

### サンプルリクエスト

```bash
curl -X POST \
  http://localhost:8888/users \
  -H 'Content-Type: application/json' \
  -d '{
    "UserId": "1",
    "Name": "John Doe",
    "Email": "john@example.com"
  }'
```

## ユーザー削除 API

### 概要

ユーザーを削除

### エンドポイント

DELETE /users/:userId

### サンプルリクエスト

```bash
curl -X DELETE http://localhost:8888/users/1
```

## ユーザー更新 API

### 概要

ユーザー内容を更新

### エンドポイント

PUT /users/:userId

### サンプルリクエスト

```bash
curl -X PUT \
  http://localhost:8888/users/1 \
  -H 'Content-Type: application/json' \
  -d '{
    "Name": "Updated Name",
    "Email": "updated@example.com"
  }'
```
