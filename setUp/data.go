package setup

import (
	"cocktail_api/common"
)

func CocktailFactory(
	name string,
	description string,
	vol int,
	ingredients []common.Ingredient,
) *common.Cocktail {
	return &common.Cocktail{
		Name: name,
		Description: description,
		Vol: vol,
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

var gin = IngredientFactory("ジン", "", 40)
var vodka = IngredientFactory("ウォッカ", "",  40)
var ram = IngredientFactory("ラム", "", 40)
var whiskey = IngredientFactory("ウィスキー", "", 40)
var cassis = IngredientFactory("カシス", "", 40)
var fuzzyNavel = IngredientFactory("ファジーネーブル", "", 40)
// plumWine := IngredientFactory("梅酒", "", true, 40)
// brandy := IngredientFactory("ブランデー", "", true, 40)
var gigerAle = IngredientFactory("ジンジャーエール", "ジンジャーの味がする炭酸飲料", 0)
var orangeJuice = IngredientFactory("オレンジジュース", "", 0)
// limeJuice := IngredientFactory("ライムジュース", "ジンジャーの味がする炭酸飲料", false, 0)

var cocktailArr = []*common.Cocktail{
	CocktailFactory("ジンバッグ", "ジンベースのカクテル", 15, []common.Ingredient{
		*gin, *gigerAle,
	}),
	CocktailFactory("スクリュードライバー", "ウォッカベースのカクテル", 15, []common.Ingredient{
		*vodka, *orangeJuice,
	}),
	CocktailFactory("ラムバッグ", "ラムベースのカクテル", 15, []common.Ingredient{
		*ram, *gigerAle,
	}),
	CocktailFactory("ジンジャーハイ", "", 15, []common.Ingredient{
		*whiskey, *gigerAle,
	}),
	CocktailFactory("カシスオレンジ", "", 15, []common.Ingredient{
		*cassis, *orangeJuice,
	}),
	CocktailFactory("ピーチオレンジ", "", 15, []common.Ingredient{
		*fuzzyNavel, *orangeJuice,
	}),
}