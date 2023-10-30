## 概要
* postgreSQL_CocktailDB_many_to_manyフォルダ内のものはDockerを使用してデータベースを立ち上げるためのもの
* setUpフォルダ内のものはデータベースにデータを挿入するためのもの
  * プログラムは`go run setUp`で実行できる
* serverフォルダ内のものはAPIサーバーを立ち上げるためのもの

## 実行する順番
1. PostgreSQL_CocktailDB_many_to_manyフォルダ内のものを実行`docker-compose up`してデータベースを立ち上げる
2. setUpフォルダ内のものを実行`go run setUp`してデータベースにデータを挿入する
3. serverフォルダ内のものを実行`go run main.go`してAPIサーバーを立ち上げる

## プログラムごとの役割
### main.go
- データベースの接続を行いデータを挿入してからAPIサーバーを立ち上げる
### server/server.go
- APIサーバー機能

  エンドポイント
    - cocktails
      - クエリに使用できる素材を渡すことによって作れるカクテルがレスポンスで帰ってくる
      - 持っている素材がジンジャエールとラムとカシスとオレンジジュースの場合のクエリの例↓
      - /cocktails?ingredients[]=ジンジャーエール&ingredients[]=ラム&ingredients[]=カシス&ingredients[]=オレンジジュース
    - ingredients
      - 全ての材料のリストが返ってくる（持っていない材料を含む）
  > **Note**
  > このAPIは現在持っている材料のリストはクライアント側で保持することを想定

### setup/setUp.go
- DBにカクテルやカクテルの素材のデータを入れる処理を書いている
### setup/data.go
- カクテルやカクテルの素材のデータを定義している


