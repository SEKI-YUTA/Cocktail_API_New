package common

type Ingredient struct {
    IngredientId int `json:"ingredient_id"`
    ShortName string `json:"short_name"`
    LongName string `json:"long_name"`
    Description string `json:"description"`
    Vol int `json:"vol"`
    IngredientCategoryId int `json:"ingredient_category_id"`
    Category string `json:"category"`
}

type Cocktail struct {
    CocktailId int `json:"cocktail_id"`
    Name string `json:"name"`
    Description string `json:"description"`
    Vol int `json:"vol"`
    ParentName string `json:"parent_name"`
    ParentCocktailId int `json:"parent_cocktail_id"`
	IngredientCount int `json:"ingredient_count"`
    Category string `json:"category"`
    CocktailCategoryId int `json:"cocktail_category_id"`
    Ingredients []Ingredient `json:"ingredients"`
}