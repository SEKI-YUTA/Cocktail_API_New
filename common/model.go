package common

type Ingredient struct {
    IngredientId int
    ShortName string
    LongName string
    Description string
    Vol int
    IngredientCategoryId int
    Category string
}

type Cocktail struct {
    CocktailId int
    Name string
    Description string
    Vol int
    ParentName string
    ParentCocktailId int
	IngredientCount int
    Category string
    CocktailCategoryId int
    Ingredients []Ingredient
}