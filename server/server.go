package server

import (
	"cocktail_api/common"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var pool *pgxpool.Pool


func StartServer() {
	fmt.Println("server.go start")
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
	// filter short, long, non_alcohol
	router.GET("/cocktails/*filter", responseCocktails)
	router.Run("localhost:9090")
	fmt.Println("end app")
}

func getAllIngredients() []common.Ingredient {
	rows, err := pool.Query(context.Background(),
	"SELECT ingredient_id, shortname, longname, description, vol, ic.ingredient_category_id, ic.name FROM ingredients " +
	"INNER JOIN ingredient_categories ic ON  ingredients.ingredient_category_id=ic.ingredient_category_id;")
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
		"SELECT cocktail_id, cocktails.name, description, cc.cocktail_category_id, cocktails.vol, ingredient_count, parent_cocktail_id FROM cocktails " +
		"INNER JOIN cocktail_categories cc ON cocktails.cocktail_category_id=cc.cocktail_category_id " +
		"WHERE cocktails.cocktail_id=$1",
		id,
		).Scan(&cocktail.CocktailId, &cocktail.Name, &cocktail.Description, &cocktail.CocktailCategoryId, &cocktail.Vol, &cocktail.IngredientCount, &cocktail.ParentCocktailId)
		fmt.Println("parentCocktail name ", cocktail.Name)
	pool.QueryRow(
		context.Background(),
		"SELECT name FROM cocktail_categories WHERE cocktail_category_id=$1",
		cocktail.CocktailCategoryId,
	).Scan(&cocktail.Category)
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
		).Scan(&cocktail.CocktailId, &cocktail.Name, &cocktail.Description, &cocktail.CocktailCategoryId, &cocktail.Category, &cocktail.Vol, &cocktail.IngredientCount, &cocktail.ParentCocktailId)

	if(cocktail.ParentCocktailId != 0) {
		fmt.Println("parentCocktailId: ", cocktail.ParentCocktailId)
		cocktail.ParentName = queryCocktailById(cocktail.ParentCocktailId).Name
	}
	return cocktail
}

func queryCocktailCategoryId(categoryName string) int {
	id := 0
	pool.QueryRow(
		context.Background(),
		"SELECT cocktail_category_id FROM cocktail_categories WHERE name=" + "'" + categoryName + "'",
	).Scan(&id)
	fmt.Println("categoryName: ", categoryName, " id: ", id)
	return id
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

func computeCraftableCocktails(availableIngredients []string, filter string) []*common.Cocktail {
	// 既に持っている材料のクエリを作る
	queryStr := getQueryString(availableIngredients)
	fmt.Println("queryStr: ", queryStr)

	// フィルタがある場合はフィルタをかける
	filterStr := ""
	categiryId := 0
	switch filter {
		case "short":
			categiryId = queryCocktailCategoryId("ショートカクテル")
			filterStr = " AND c.cocktail_category_id = " + strconv.Itoa(categiryId)
			break
		case "long":
			categiryId = queryCocktailCategoryId("ロングカクテル")
			filterStr = " AND c.cocktail_category_id = " + strconv.Itoa(categiryId)
			break
		case "non_alcohol":
			filterStr = " AND c.vol = 0"
			break
		default:
			break
	}
	fmt.Println("filterStr: ", filterStr)


	cocktailIngredientMap := map[string][]common.Ingredient{}
	cocktailIngredientCountMap := map[string]int{}
	rows, err := pool.Query(
		context.Background(),
		"SELECT c.cocktail_id, c.name, c.ingredient_count, ic.name, i.ingredient_id, i.shortname, i.longname, i.description, i.vol, i.ingredient_category_id FROM cocktail_ingredients " +
		"INNER JOIN ingredients i ON i.ingredient_id = cocktail_ingredients.ingredient_id " +
		"INNER JOIN cocktails c ON c.cocktail_id = cocktail_ingredients.cocktail_id " +
		"INNER JOIN ingredient_categories ic ON i.ingredient_category_id = ic.ingredient_category_id " +
		"WHERE i.longname IN " + queryStr + filterStr,
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to get craftableCocktail data %v\n", err)
		// os.Exit(1)
	}

	for rows.Next() {
		cocktailId := 0
		cocktailName := ""
		ingredientCount := 0
		ingredientCategory := ""
		ingredientId := 0
		ingredientShortName := ""
		ingredientLongName := ""
		ingredientDescription := ""
		ingredientVol := 0
		ingredientCategoryId := 0
		rows.Scan(&cocktailId, &cocktailName, &ingredientCount, &ingredientCategory, &ingredientId, &ingredientShortName, &ingredientLongName, &ingredientDescription, &ingredientVol, &ingredientCategoryId)
		cocktailIngredientCountMap[cocktailName + "-" + strconv.Itoa(cocktailId)] = ingredientCount
		cocktailIngredientMap[cocktailName + "-" + strconv.Itoa(cocktailId)] = append(
			cocktailIngredientMap[cocktailName + "-" + strconv.Itoa(cocktailId)],
			common.Ingredient{IngredientId: ingredientId,Category: ingredientCategory, ShortName: ingredientShortName, LongName: ingredientLongName, Description: ingredientDescription, Vol: ingredientVol, IngredientCategoryId: ingredientCategoryId},
		)
	}
	craftableCocktailArr := []*common.Cocktail{}
	for key, ingredientArr := range cocktailIngredientMap {
		ingredientCount := cocktailIngredientCountMap[key]
		if len(ingredientArr) == ingredientCount {
			sp := strings.Split(key, "-")
			id, err := strconv.Atoi(sp[1])
			if err != nil {
				
			}
			cocktail := queryCocktailById(id)
			cocktail.Ingredients = ingredientArr
			craftableCocktailArr = append(craftableCocktailArr, &cocktail)
		}
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
	filter := strings.Replace(ctx.Param("filter"), "/", "", -1)
	fmt.Println("filter: ", filter)
	query := ctx.Request.URL.Query()
	availableIngredients := query["ingredients[]"]
	fmt.Println("availableIngredients: ", availableIngredients)
	cocktails := computeCraftableCocktails(availableIngredients, filter)

	ctx.JSON(200, cocktails)
}
