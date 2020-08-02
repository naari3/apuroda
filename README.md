# apuroda

## setup

まだ docker-compose に乗りきってない(dbだけdocker-composeに載せた)

## how to run

```bash
$ docker-compose up -d
$ PGSSLMODE=disable psqldef apuroda --password=postgres --port=5433 < schema.sql
$ go run main.go
```

## how to use

1. http://localhost:8080/new_user
1. ユーザー作成
1. http://localhost:8080/new_file
1. アップロード
1. 各個別ページでダウンロード
1. http://localhost:8080/ でファイル一覧確認

## unimplemented

- 各ユーザーのアップロードしたファイル
- まいかいgin.Hでuserを渡さないようにするやつ
  - やり方がわからん