# sample-mvc-api-go

- 基本的なCRUD操作を行うREST API
- TODOアプリのバックエンドのイメージでTODO(タイトルと内容)の取得/追加/更新/削除が行える
- Go言語のサンプルコードによくあるクリーンアーキテクチャに影響を受けた構成ではなく、MVC風なもの
- 参考サイト: https://dev.classmethod.jp/articles/go-sample-rest-api/

# プロジェクト構成図

```text
sample-mvc-api-go/                  ルートディレクトリ
    ┣ build/                        Dockerfileなど
    ┃   ┣ db/                       動作確認用DB
    ┃   ┃   ┣ sql/                  DDLとテストデータ投入用SQL
    ┃   ┃   ┗ Dockerfile
    ┃   ┣ sample-mvc-api-go/ 
    ┃   ┃   ┗ Dockerfile            このプロジェクトの実行ファイルを含んだイメージを作成するDockerfile
    ┃   ┗ docker-compose.yml
    ┣ cmd/
    ┃   ┣ sample-mvc-api-go
    ┃   ┗ main.go                   メイン処理
    ┣ controller/
    ┃   ┣ dto/                      リクエスト/レスポンス用のDTOファイルを配置する
    ┃   ┣ router.go                 HTTPメソッドを元にコントローラの各処理へのルーティングを行う
    ┃   ┣ router_test.go            `router.go`のテストファイル
    ┃   ┣ todo_controller.go        リクエストを元にモデルの各処理を呼び出しレスポンスを返却する
    ┃   ┗ todo_controller_test.go   `todo_controller.go`のテストファイル 
    ┣ model/
    ┃   ┣ entity/　
    ┃   ┗ repository/
    ┣ test/
    ┃   ┗ mock.go                   単体テスト用のモック
    ┣ test_results/                 単体テストのカバレッジファイルを配置する            
    ┣ Makefile
    ┣ go.mod
    ┗ go.sum
```

# 実行手順

## サーバー起動

下記makeコマンドで、8080ポートにAPIサーバーが起動する。

```shell
make serve
```

## リクエスト例

サーバー起動しながらHTTPリクエストすると動作確認ができる。

### 取得APIリクエスト

```shell
curl -i localhost:8080/todos/
```

### 新規登録APIリクエスト

```shell
curl -i -X POST -H "Content-Type: application/json" -d '{"title":"test", "content":"テストです。"}' localhost:8080/todos/
```

### 更新APIリクエスト

```shell
curl -i -X PUT -H "Content-Type: application/json" -d '{"title":"test", "content":"変更テスト"}' localhost:8080/todos/4
```

### 削除APIリクエスト

```shell
curl -i -X DELETE localhost:8080/todos/4
```
