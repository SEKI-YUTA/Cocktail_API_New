package setup

import (
	"cocktail_api/common"
)

const cocktail_category_short = "ショートカクテル"
const cocktail_category_long = "ロングカクテル"
const cocktail_category_non_alcohol = "ノンアルコール"

// 国によって扱いが違うからどう対応しようか
const ingredient_category_spirit = "スピリッツ"
const ingredient_category_liqueur = "リキュール"
const ingredient_category_softdrink = "ソフトドリンク"
const ingredient_category_whisky = "ウィスキー"
const ingredient_category_brandy = "ブランデー"
const ingredient_category_shothu = "焼酎"
const ingredient_category_japanese_style_alcohol = "日本酒"
const ingredient_category_decoration = "飾り"

func CocktailFactory(
	name string,
	category string,
	description string,
	vol int,
	parentName string,
	ingredientCount int,
	ingredients []common.Ingredient,
) *common.Cocktail {
	return &common.Cocktail{
		Name: name,
		Category: category,
		Description: description,
		Vol: vol,
		ParentName: parentName,
		IngredientCount: ingredientCount,
		Ingredients: ingredients,
	}
}

func IngredientFactory(
	shortName string,
	longName string,
    description string,
	category string,
    vol int,
) *common.Ingredient {
	return &common.Ingredient{
		ShortName: shortName,
		LongName: longName,
		Category: category,
		Description: description,
		Vol: vol,
	}
}

const dbURL string = "postgres://root:root@localhost:5432/cocktail_db"
const cocktail_table string = "cocktails"
const cocktail_categories_table string = "cocktail_categories"
const ingredient_categories_table string = "ingredient_categories"
const ingredient_table string = "ingredients"

var WhiteCuracao = IngredientFactory("キュラソー", "ホワイトキュラソー", "", ingredient_category_spirit,40)

var Gin = IngredientFactory("ジン", "ドライ・ジン", "", ingredient_category_spirit, 40)
var Vodka = IngredientFactory("ウォッカ", "ウォッカ", "", ingredient_category_spirit, 40)
var Ram = IngredientFactory("ラム", "ホワイトラム", "", ingredient_category_spirit, 40)
// var Whiskey = IngredientFactory("ウィスキー", "", 40)
// var Curacao = IngredientFactory("キュラソー", "", 40)
// var Tequila = IngredientFactory("テキーラ", "", 40)
// var Vermouth = IngredientFactory("ベルモット", "", 15)
// var Cassis = IngredientFactory("カシス", "", 15)
// var fuzzyNavel = IngredientFactory("ファジーネーブル", "", 15)
// var PlumWine = IngredientFactory("梅酒", "", 40)
var Brandy = IngredientFactory("ブランデー", "コニャック", "", ingredient_category_brandy, 40)
// var Beer = IngredientFactory("ビール", "", 0)
var GigerAle = IngredientFactory("ジンジャーエール", "ジンジャーエール", "ジンジャーの味がする炭酸飲料", ingredient_category_softdrink, 0)
// var OrangeJuice = IngredientFactory("オレンジジュース", "", 0)
var LemonJuice = IngredientFactory("レモンジュース", "レモンジュース", "", ingredient_category_softdrink, 0)
// var LimeJuice = IngredientFactory("ライムジュース", "", 0)
// var TonicWater = IngredientFactory("トニックウォーター", "", 0)

var CocktailArr = []*common.Cocktail{
	CocktailFactory("ジンバッグ", cocktail_category_long , "ジンベースのカクテル", 15, "", 2, []common.Ingredient{
		*Gin, *GigerAle,
	}),
	CocktailFactory("サイドカー", cocktail_category_short , "", 20, "", 3, []common.Ingredient{
		*Brandy, *WhiteCuracao, *LemonJuice,
	}),
	CocktailFactory("ホワイトレディー", cocktail_category_short , "ジンベースのカクテル", 20, "サイドカー", 3, []common.Ingredient{
		*Gin, *WhiteCuracao, *LemonJuice,
	}),
	// CocktailFactory("スクリュードライバー", "ウォッカベースのカクテル", 15, 2, []common.Ingredient{
	// 	*Vodka, *OrangeJuice,
	// }),
	// CocktailFactory("ラムバッグ", "ラムベースのカクテル", 15, 2, []common.Ingredient{
	// 	*Ram, *GigerAle,
	// }),
	// CocktailFactory("ジンジャーハイ", "", 15, 2, []common.Ingredient{
	// 	*Whiskey, *GigerAle,
	// }),
	// CocktailFactory("カシスオレンジ", "", 15, 2, []common.Ingredient{
	// 	*Cassis, *OrangeJuice,
	// }),
	// CocktailFactory("ピーチオレンジ", "", 15, 2, []common.Ingredient{
	// 	*fuzzyNavel, *OrangeJuice,
	// }),
	// CocktailFactory("モスコミュール", "", 15, 2, []common.Ingredient{
	// 	*Vodka, *GigerAle, *LimeJuice,
	// }),
	// CocktailFactory("梅酒ビア", "", 15, 2, []common.Ingredient{
	// 	*Beer, *PlumWine,
	// }),
	// CocktailFactory("ブランデートニック", "", 15, 2, []common.Ingredient{
	// 	*Brandy, *TonicWater,
	// }),
}