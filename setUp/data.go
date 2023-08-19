package main

type Ingredient struct {
    name string
    description string
    vol int
}

type Cocktail struct {
    name string
    description string
    vol int
    ingredients []Ingredient
}

func CocktailFactory(
	name string,
	description string,
	vol int,
	ingredients []Ingredient,
) *Cocktail {
	return &Cocktail{
		name: name,
		description: description,
		vol: vol,
		ingredients: ingredients,
	}
}

func IngredientFactory(
	name string,
    description string,
    vol int,
) *Ingredient {
	return &Ingredient{
		name: name,
		description: description,
		vol: vol,
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

var cocktailArr = []*Cocktail{
	CocktailFactory("ジンバッグ", "ジンベースのカクテル", 15, []Ingredient{
		*gin, *gigerAle,
	}),
	CocktailFactory("スクリュードライバー", "ウォッカベースのカクテル", 15, []Ingredient{
		*vodka, *orangeJuice,
	}),
	CocktailFactory("ラムバッグ", "ラムベースのカクテル", 15, []Ingredient{
		*ram, *gigerAle,
	}),
	CocktailFactory("ジンジャーハイ", "", 15, []Ingredient{
		*whiskey, *gigerAle,
	}),
	CocktailFactory("カシスオレンジ", "", 15, []Ingredient{
		*cassis, *orangeJuice,
	}),
	CocktailFactory("ピーチオレンジ", "", 15, []Ingredient{
		*fuzzyNavel, *orangeJuice,
	}),
}