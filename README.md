# flyme-backend

### 実行方法

```
$ PORT=xxx \
  GOOGLE_APPLICATION_CREDENTIALS=xxx \
  PROJECT_ID=xxx \
  go run cmd/main.go
```

### 環境変数

`PORT`: ポート番号<br>
`GOOGLE_APPLICATION_CREDENTIALS`:FirebaseのCredentialsの入ったファイルへのパス<br>
`PROJECT_ID`:FirebaseのプロジェクトID