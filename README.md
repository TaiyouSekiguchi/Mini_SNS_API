# Mini_SNS_API

このプロジェクトでは、Go言語を使ったSNSのWebAPIを作成しました。

## 使用技術

- Go言語
- Echoフレームワーク
- MySQL
- Docker

## エンドポイント

### GET /get_friend_list

クエリパラメータでユーザーのIDを受け取り、そのユーザーの友達の一覧を返す。

### GET /get_friend_of_friend_list

クエリパラメータでユーザーのIDを受け取り、そのユーザーの友達の友達一覧を返す。
ブロックユーザーは除く。

### GET /get_friend_of_friend_list_paging

クエリパラメータでユーザーのIDと1ページあたり表示件数、ページ番号を受け取り、そのユーザーの友達の一覧を返す。
ブロックユーザーは除く。

## 意識したところ

- ディレクトリ構成を役割ごとに分割してわかりやすくしました。
- Airを使ってホットリロードを導入しました。
- コントローラーとモデルを用意して、役割を分けました。
- Go doc でコメントが確認できるようにしました。
- WebAPIの返り値において、わかりやすい構成を心がけました。
