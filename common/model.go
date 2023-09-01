package common

type Ingredient struct {
    // 材料のプライマリーキー
    IngredientId int `json:"ingredient_id"`

    // 材料の大まかな名前 ex: ウィスキー
    ShortName string `json:"short_name"`

    // 材料の正式名称 ex: バーボンウィスキー
    LongName string `json:"long_name"`

    // 材料の説明
    Description string `json:"description"`

    // 度数
    Vol int `json:"vol"`

    // 材料のカテゴリーID
    IngredientCategoryId int `json:"ingredient_category_id"`

    // 材料のカテゴリー名 スピリッツ、リキュール、ソフトドリンク、ウィスキー、ブランデー、焼酎、日本酒、飾り
    Category string `json:"category"`
}

type Cocktail struct {
    // カクテルのプライマリーキー
    CocktailId int `json:"cocktail_id"`

    // カクテルの名前 ex: ホワイトレディ
    Name string `json:"name"`

    // カクテルの説明
    Description string `json:"description"`

    // カクテルの度数
    Vol int `json:"vol"`

    // 親カクテルの名前 ex: サイドカー
    ParentName string `json:"parent_name"`

    // 親カクテルのID
    ParentCocktailId int `json:"parent_cocktail_id"`

    // 材料の数（/cocktailエンドポイントで作れるカクテルを検索するときに使用するので必要）
	IngredientCount int `json:"ingredient_count"`

    // カクテルのカテゴリー名 ショートカクテル、ロングカクテル、ノンアルコール
    Category string `json:"category"`

    // カクテルのカテゴリーID
    CocktailCategoryId int `json:"cocktail_category_id"`

    // カクテルに使用する材料の配列
    Ingredients []Ingredient `json:"ingredients"`
}