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
### server/server.go（未実装)
- APIサーバー機能

  エンドポイント
    - cocktails
      - クエリに使用できる素材を渡すことによって作れるカクテルがレスポンスで帰ってくる
    - ingredients
      - 全ての材料のリストが返ってくる（持っていない材料を含む）
  > **Note**
  > このAPIは現在持っている材料のリストはクライアント側で保持することを想定
###　setup/setUp.go
- DBにカクテルやカクテルの素材のデータを入れる処理を書いている
### setup/data.go
- カクテルやカクテルの素材のデータを定義している


### データベースへデータを入れる大まかな流れ
カクテルの構造体を定義して、カクテル型の配列を用意する。
まず、cocktailテーブルに既に同じ名前のデータがあるかを確認してなければ挿入（何かしらの方法で挿入したデータのidを取得しておく）。もし既に存在していた場合はこの時点で処理が終了する。
次に、素材の配列を回して、ingredientsテーブルに同じ名前のデータがあるかを確認してなければ挿入（何かしらの方法で挿入したデータのidを取得しておいて配列に保持しておく）。もし既に存在していた場合は現在の素材の処理は終了（continue)。
以上の処理をカクテルの配列で回す。

### カクテルの構造体と素材の構造体（簡単なメモ）
``` go
type Cocktial struct {
    name string
    description string
    vol int
    maxVol int
    minVol int
    ingredients []Ingredient
}

type Ingredient struct {
    name string
    description string
    isAlcohol bool
    vol int
    maxVol int
    minVol int
}

// これより下は大雑把なプログラム

cocktails :[]{
    // カクテルの構造体の配列
}

for _, item := range cocktails {
    cocktailName = item.name
    if(すでに同じカクテル名のデータが存在する) {
        continue
    }
    val insertedId = カクテルのデータを挿入する

    val ingredients = item.ingredients

    ingredientIds :[]{
        // 素材のidの配列(中間テーブルに挿入する際に使用する)
    }
    for _, ingredient := range ingredients {
        ingredientName = ingredient.name
        if(すでに同じ素材名のデータが存在する) {
            ingredientIds = append(ingredientIds, すでに存在している素材のid)
            continue
        }
        val insertedId = 素材のデータを挿入する
        append(ingredientIds, insertedId)
    }

    for _, ingredientId := range ingredientIds {
        中間テーブルにカクテルのidと素材のidを挿入する
        // insert into cocktail_ingredient (cocktail_id, ingredient_id) values (insertedId, ingredientId)
}

// 以下コンストラクタ関数
// functional option patternを使ったらデフォルト値を実現できるらしい
// func CocktailFactory(
//       name string,
//       description string,
//       vol int,
//       maxVol int,
//       minVol int,
//       ingredients []Ingredient,
//   ) *Cocktail {
//       return Cocktail{
//           name: name,
//           description: description,
//           vol: vol,
//           maxVol: maxVol,
//           minVol: minVol,
//           ingredients: ingredients,
//       }
//   }
```


<!-- ### プログラムを書く上でのメモ -->
