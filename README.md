# flyme-backend

### 実行方法

```
$ PORT=xxx \
  GOOGLE_APPLICATION_CREDENTIALS=xxx \
  PROJECT_ID=xxx \
  YOLP_APPID=xxx \
  BUCKET_ID=xxx \
  MODE=xxx \
  go run cmd/main.go
```

### 環境変数

`PORT`: ポート番号<br>
`GOOGLE_APPLICATION_CREDENTIALS`: FirebaseのCredentialsの入ったファイルへのパス<br>
`PROJECT_ID`: FirebaseのプロジェクトID<br>
`YOLP_APPID`: [YOLP](https://developer.yahoo.co.jp/webapi/map/)のアプリケーションID<br>
`BUCKET_ID`: Firebase Storageアクセス用のID<br>
`MODE`: 実行モード

<!-- ### クレジット

**YOLP**<br>
[Webサービス by Yahoo! JAPAN](https://developer.yahoo.co.jp/sitemap/) -->
