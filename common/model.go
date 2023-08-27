package common

type Ingredient struct {
    ShortName string
    LongName string
    Category string
    Description string
    Vol int
}

type Cocktail struct {
    Name string
    Category string
    Description string
    Vol int
    ParentName string
	IngredientCount int
    Ingredients []Ingredient
}