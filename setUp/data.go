package setup

import (
	"cocktail_api/common"
)

func CocktailFactory(
	name string,
	description string,
	vol int,
	ingredientCount int,
	ingredients []common.Ingredient,
) *common.Cocktail {
	return &common.Cocktail{
		Name: name,
		Description: description,
		Vol: vol,
		IngredientCount: ingredientCount,
		Ingredients: ingredients,
	}
}

func IngredientFactory(
	name string,
    description string,
    vol int,
) *common.Ingredient {
	return &common.Ingredient{
		Name: name,
		Description: description,
		Vol: vol,
	}
}

const dbURL string = "postgres://root:root@localhost:5432/cocktail_db"
const cocktail_table string = "cocktails"
const ingredient_table string = "ingredients"

var Gin = IngredientFactory("ジン", "", 40)
var Vodka = IngredientFactory("ウォッカ", "",  40)
var Ram = IngredientFactory("ラム", "", 40)
var Whiskey = IngredientFactory("ウィスキー", "", 40)
var Cassis = IngredientFactory("カシス", "", 40)
var fuzzyNavel = IngredientFactory("ファジーネーブル", "", 40)
// plumWine := IngredientFactory("梅酒", "", true, 40)
// brandy := IngredientFactory("ブランデー", "", true, 40)
var GigerAle = IngredientFactory("ジンジャーエール", "ジンジャーの味がする炭酸飲料", 0)
var OrangeJuice = IngredientFactory("オレンジジュース", "", 0)
// limeJuice := IngredientFactory("ライムジュース", "ジンジャーの味がする炭酸飲料", false, 0)

var CocktailArr = []*common.Cocktail{
	CocktailFactory("ジンバッグ", "ジンベースのカクテル", 15, 2, []common.Ingredient{
		*Gin, *GigerAle,
	}),
	CocktailFactory("スクリュードライバー", "ウォッカベースのカクテル", 15, 2, []common.Ingredient{
		*Vodka, *OrangeJuice,
	}),
	CocktailFactory("ラムバッグ", "ラムベースのカクテル", 15, 2, []common.Ingredient{
		*Ram, *GigerAle,
	}),
	CocktailFactory("ジンジャーハイ", "", 15, 2, []common.Ingredient{
		*Whiskey, *GigerAle,
	}),
	CocktailFactory("カシスオレンジ", "", 15, 2, []common.Ingredient{
		*Cassis, *OrangeJuice,
	}),
	CocktailFactory("ピーチオレンジ", "", 15, 2, []common.Ingredient{
		*fuzzyNavel, *OrangeJuice,
	}),
}