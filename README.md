# apuroda

## setup

まだ docker-compose に乗ってないし、まだオンメモリでしか動かない

## how to run

```bash
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
- 永続化機構(レイヤードなので簡単にできる想定)
- まいかいgin.Hでuserを渡さないようにするやつ
  - やり方がわからん