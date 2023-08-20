package server

import (
	"cocktail_api/common"
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var pool *pgxpool.Pool


func StartServer() {
	fmt.Println("server.go start")

	// fmt.Println("user name: " + first_user)
	connConfig, err := pgx.ParseConfig(DB_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse db config")
		os.Exit(1)
	}

	poolConfig, err := pgxpool.ParseConfig("")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse pool config")
		os.Exit(1)
	}
	poolConfig.ConnConfig = connConfig

	pool, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to connect to db")
		os.Exit(1)
	}

	defer pool.Close()

	fmt.Println("start app")
	router := gin.Default()
	router.GET("/ingredients", responseAllIngredients)
	router.GET("/cocktails", responseCocktails)
	router.Run("localhost:9090")
	fmt.Println("end app")
}

func getAllIngredients() []common.Ingredient {
	rows, err := pool.Query(context.Background(),"select name, description, vol from ingredients;")
	// rows, err := conn.Query(context.Background(),"select * from user_list;")
	// err = conn.QueryRow(context.Background(), "select name from user_list where id = 1;").Scan(&first_user)
	// defer rows.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to get user name %v\n", err)
		os.Exit(1)
	}
	ingredients := []common.Ingredient{}
	for rows.Next() {
		var i common.Ingredient
		err := rows.Scan(&i.Name, &i.Description, &i.Vol)
		if err != nil {
			fmt.Println("failed to scan data")
		}
		ingredients = append(ingredients, i)
	}
	return ingredients
}

func queryCocktail(cocktailName string) common.Cocktail {
	cocktail := common.Cocktail{}
	pool.QueryRow(context.Background(), "select name, description, vol from cocktails where name = $1;", cocktailName).
	Scan(&cocktail.Name, &cocktail.Description, &cocktail.Vol)

	return cocktail
}

func getQueryString(availableIngredients []string) string {
	arrLenght := len(availableIngredients)
	queryStr := ""
	for i, ingredient := range availableIngredients {
		queryStr += "'" + ingredient + "'"
		if i < arrLenght-1 {
			queryStr += ", "
		}
	}
	return "(" + queryStr + ")"
}

func computeCraftableCocktails(availableIngredients []string) []*common.Cocktail {
	queryStr := getQueryString(availableIngredients)

	cocktailIngredientMap := map[string][]common.Ingredient{}
	cocktailIngredientCountMap := map[string]int{}
	rows, err := pool.Query(
		context.Background(),
		"select c.name, c.ingredient_count, i.name, i.description, i.vol from cocktail_ingredients INNER JOIN ingredients i ON i.ingredient_id = cocktail_ingredients.ingredient_id INNER JOIN cocktails c ON c.cocktail_id = cocktail_ingredients.cocktail_id WHERE i.name IN " + queryStr,
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to get craftableCocktail data %v\n", err)
		// os.Exit(1)
	}

	for rows.Next() {
		cocktailName := ""
		ingredientCount := 0
		ingredientName := ""
		ingredientDescription := ""
		ingredientVol := 0
		rows.Scan(&cocktailName, &ingredientCount, &ingredientName, &ingredientDescription, &ingredientVol)
		cocktailIngredientCountMap[cocktailName] = ingredientCount
		cocktailIngredientMap[cocktailName] = append(
			cocktailIngredientMap[cocktailName],
			common.Ingredient{Name: ingredientName, Description: ingredientDescription, Vol: ingredientVol},
		)
	}

	craftableCocktailArr := []*common.Cocktail{}
	for key, ingredientArr := range cocktailIngredientMap {
		ingredientCount := cocktailIngredientCountMap[key]
		// fmt.Println("key: ", key, " ingredientCount: ", ingredientCount, " ingredientArr: ", ingredientArr)
		if len(ingredientArr) == ingredientCount {
			fmt.Println(key + "を作れるよ！")
			cocktail := queryCocktail(key)
			cocktail.Ingredients = ingredientArr
			craftableCocktailArr = append(craftableCocktailArr, &cocktail)
		}
	}

	for _, cocktail := range craftableCocktailArr {
		fmt.Println(cocktail)
	}

	/*
	select c.name, c.vol, c.ingredient_count, i.name from cocktail_ingredients
	INNER JOIN ingredients i ON i.ingredient_id = cocktail_ingredients.ingredient_id
	INNER JOIN cocktails c ON c.cocktail_id = cocktail_ingredients.cocktail_id
	WHERE i.name IN ('ジン', 'トニックウォーター')
	上記のクエリで現在持っている材料が使われているカクテルを取得できる
	そして、カクテル名をキーに持つmapを作成してvalに材料をほりこむ
	材料の配列の長さがカクテルの材料数と一致したら、そのカクテルは作れる
	key: カクテル名 val []材料名
	key: カクテル名 val 材料数
	*/

	return craftableCocktailArr
}

func responseAllIngredients(ctx *gin.Context) {
	ingredients := getAllIngredients()
	ctx.JSON(200, ingredients)
}

func responseCocktails(ctx *gin.Context) {
	query := ctx.Request.URL.Query()
	availableIngredients := query["ingredients[]"]
	fmt.Println("availableIngredients: ", availableIngredients)
	fmt.Printf("type: %T\n ", availableIngredients)

	cocktails := computeCraftableCocktails(availableIngredients)

	ctx.JSON(200, cocktails)
}
