package common

type Ingredient struct {
    Name string
    Description string
    Vol int
}

type Cocktail struct {
    Name string
    Description string
    Vol int
	IngredientCount int
    Ingredients []Ingredient
}