# go_todo_app
Todo web application with AUTH by Go.

## 動作確認

カレントディレクトのmain関数実行
go の場合、標準で並列リクエスト可能
```cmd
go run .
```

## テスト容易性の確保

main関数で HTTP サーバを起動すると、戻り値が検証できません。

* テスト完了後に終了できない

* 出力を検証しにくい

* 異常時に、os.Exit が呼ばれてしまう

* ポート番号が固定されていると、サーバーの起動に失敗する可能性がある

上記を踏まえて、run 関数を用意しました。

## 動作検証

### タスク一覧の取得

*`curl`コマンドによる動作検証例*

Windowsのコマンドプロンプトの利用を推奨しない。

[WindowsでcurlコマンドでJSONを送信する](https://qiita.com/Hina_Developer/items/e583021a44a753e29dde)


```shell
$ curl -i -X POST localhost:80/tasks -d @.\handler\testdata\add_task\ok_req.json.golden 
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Fri, 16 Aug 2024 17:49:30 GMT
Content-Length: 8

{"id":4}

$ curl -i -X POST localhost:80/tasks -d @.\handler\testdata\add_task\bad_req.json.golden 
HTTP/1.1 400 Bad Request
Content-Type: application/json; charset=utf-8
Date: Fri, 16 Aug 2024 17:49:57 GMT
Content-Length: 90

$ curl -X POST localhost:80/register -d '{"name": "john2", "password":"test", "role":"user"}'

$ curl -i -X POST localhost:80/login -d '{"user_name": "budou", "password":"test", "role":"admin"}' |  tail
```
### database migration
```cmd
mysqldef.exe -u todo -p todo -h 192.168.0.20 -P 3306 todo < "E:\Dropbox\GoDev\Projects\go_todo_app\_tools\mysql\schema.sql"
```