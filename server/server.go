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
	/*
	select * from ingredients 
	INNER JOIN ingredient_categories ic ON  ingredients.ingredient_category_id=ic.ingredient_category_id
	*/
	rows, err := pool.Query(context.Background(),
	"select ingredient_id, shortname, longname, description, vol, ic.ingredient_category_id, ic.name from ingredients INNER JOIN ingredient_categories ic ON  ingredients.ingredient_category_id=ic.ingredient_category_id;")
	// rows, err := conn.Query(context.Background(),"select * from user_list;")
	// err = conn.QueryRow(context.Background(), "select name from user_list where id = 1;").Scan(&first_user)
	// defer rows.Close()
	if err != nil {
		os.Exit(1)
	}
	ingredients := []common.Ingredient{}
	for rows.Next() {
		var i common.Ingredient
		err := rows.Scan(&i.IngredientId,&i.ShortName, &i.LongName, &i.Description, &i.Vol, &i.IngredientCategoryId, &i.Category)
		if err != nil {
			fmt.Println("failed to scan data")
		}
		ingredients = append(ingredients, i)
	}
	return ingredients
}

func queryCocktailById(id int) common.Cocktail {
	cocktail := common.Cocktail{}
	pool.QueryRow(
		context.Background(),
		"SELECT cocktail_id, cocktails.name, description, cc.cocktail_category_id, cocktails.vol, ingredient_count, parent_cocktail_id from cocktails " +
		"INNER JOIN cocktail_categories cc ON cocktails.cocktail_category_id=cc.cocktail_category_id WHERE cocktails.cocktail_id=$1",
		id,
		).Scan(&cocktail.CocktailId, &cocktail.Name, &cocktail.Description, &cocktail.CocktailCategoryId, &cocktail.Vol, &cocktail.IngredientCount, &cocktail.ParentCocktailId)
		fmt.Println("parentCocktail name ", cocktail.Name)
	return cocktail
}


func queryCocktail(cocktailName string) common.Cocktail {
	cocktail := common.Cocktail{}
	fmt.Println("cocktailName: ", cocktailName)
	pool.QueryRow(
		context.Background(),
		"SELECT cocktail_id, cocktails.name, description, cc.cocktail_category_id, cc.name, cocktails.vol, ingredient_count, parent_cocktail_id from cocktails " +
		"INNER JOIN cocktail_categories cc ON cocktails.cocktail_category_id=cc.cocktail_category_id " +
		"WHERE cocktails.name=" + "'" + cocktailName + "';",
		// "select cocktails.name, description, vol from cocktails where name = $1;",
		// ).Scan(&cocktail.Name, &cocktail.Description, &cocktail.Vol)
		).Scan(&cocktail.CocktailId, &cocktail.Name, &cocktail.Description, &cocktail.CocktailCategoryId, &cocktail.Category, &cocktail.Vol, &cocktail.IngredientCount, &cocktail.ParentCocktailId)

	if(cocktail.ParentCocktailId != 0) {
		fmt.Println("parentCocktailId: ", cocktail.ParentCocktailId)
		cocktail.ParentName = queryCocktailById(cocktail.ParentCocktailId).Name
	}

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
	fmt.Println("queryStr: ", queryStr)

	cocktailIngredientMap := map[string][]common.Ingredient{}
	cocktailIngredientCountMap := map[string]int{}
	rows, err := pool.Query(
		context.Background(),
		"select c.name, c.ingredient_count, ic.name, i.ingredient_id, i.shortname, i.longname, i.description, i.vol, i.ingredient_category_id from cocktail_ingredients " +
		"INNER JOIN ingredients i ON i.ingredient_id = cocktail_ingredients.ingredient_id " +
		"INNER JOIN cocktails c ON c.cocktail_id = cocktail_ingredients.cocktail_id " +
		"INNER JOIN ingredient_categories ic ON i.ingredient_category_id = ic.ingredient_category_id " +
		"WHERE i.longname IN " + queryStr,
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to get craftableCocktail data %v\n", err)
		// os.Exit(1)
	}

	for rows.Next() {
		fmt.Println("rows.Next()")
		cocktailName := ""
		ingredientCount := 0
		ingredientCategory := ""
		ingredientId := 0
		ingredientShortName := ""
		ingredientLongName := ""
		ingredientDescription := ""
		ingredientVol := 0
		ingredientCategoryId := 0
		rows.Scan(&cocktailName, &ingredientCount, &ingredientCategory, &ingredientId, &ingredientShortName, &ingredientLongName, &ingredientDescription, &ingredientVol, &ingredientCategoryId)
		cocktailIngredientCountMap[cocktailName] = ingredientCount
		cocktailIngredientMap[cocktailName] = append(
			cocktailIngredientMap[cocktailName],
			common.Ingredient{IngredientId: ingredientId,Category: ingredientCategory, ShortName: ingredientShortName, LongName: ingredientLongName, Description: ingredientDescription, Vol: ingredientVol, IngredientCategoryId: ingredientCategoryId},
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

	/*
	SELECT cocktails.name, c2.name from cocktails
	INNER JOIN cocktails c2 ON c2.cocktail_id = cocktails.parent_id
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
