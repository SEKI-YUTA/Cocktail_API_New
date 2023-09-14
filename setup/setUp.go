package setup

import (
	"cocktail_api/common"
	"context"
	"fmt"
	"os"
	"reflect"
	"sort"
	// v5だとpoolができない
	"github.com/jackc/pgx/v4"
)

/*
この関数は同じ名前があれば既に存在していると判断して０以外の値を返す
既に存在していた場合は０以外の値を返す(id)
*/
func checkExists(
	tableName string,
	idColName string,
	nameColName string,
	name string,
	conn *pgx.Conn,
) int {
	fmt.Println("checkExists: ", tableName, name)
	var id int
	conn.QueryRow(
		context.Background(),
		"SELECT " + idColName + " FROM " + tableName + " WHERE " + nameColName + "=$1",
		name).Scan(&id)
	if(id == 0) {
		fmt.Println("not exists: ", name)
	} else {
		fmt.Println("exists: ", name)
	}
	return id
}

/*
この関数はカクテルの存在確認特化の関数で、カクテルの名前と素材が同じであれば既に存在していると判断して０以外の値を返す
*/
func checkExistsCocktail(
	cocktail *common.Cocktail,
	conn *pgx.Conn,
) int {
	fmt.Println("checkExistsCocktail: ", cocktail.Name)
	var idArr = []int{}
	rows1, _ := conn.Query(
		context.Background(),
		"SELECT cocktail_id FROM cocktails WHERE name=$1",
		cocktail.Name,
	)
	for rows1.Next() {
		tmp := 0
		rows1.Scan(&tmp)
		if(tmp != 0) {
			idArr = append(idArr, tmp)
		}
	}
	if(len(idArr) == 0) {
		// この時点でidArrの長さが0であれば、名前で該当するカクテルがない＝存在していないと判断して０を返す
		return 0
	}

	for _, id := range idArr {
		rows, _ := conn.Query(
			context.Background(),
			"SELECT ingredients.longname FROM cocktail_ingredients " +
			"INNER JOIN ingredients ON cocktail_ingredients.ingredient_id=ingredients.ingredient_id " +
			"WHERE cocktail_ingredients.cocktail_id=$1",
			id,
		)
		ingredientNames := []string{}
		for rows.Next() {
			var ingredientName string
			rows.Scan(&ingredientName)
			ingredientNames = append(ingredientNames, ingredientName)
		}
		fmt.Println("ingredientNames: ", ingredientNames)
		if(len(ingredientNames) != len(cocktail.Ingredients)) {
			// チェックを要求されたカクテルとデータベースを検索で見つかったカクテルと材料の数が違うので違うカクテルと判断する
			// この部分でcontinenuされるのはカクテルの名前は同じだが材料の数が違う場合（そんなものがあるかは知らん）
			continue
		}
		cocktailIngredientNames := []string{}
		for _, ingredient := range cocktail.Ingredients {
			cocktailIngredientNames = append(cocktailIngredientNames, ingredient.LongName)
		}
		fmt.Println("cocktailIngredientNames: ", cocktailIngredientNames)
		sort.Strings(ingredientNames)
		sort.Strings(cocktailIngredientNames)
		// この時点で材料の数が同じであることは保証されているのであとは材料の名前が全部同じかをチェックする
		if(reflect.DeepEqual(ingredientNames, cocktailIngredientNames)) {
			// 素材の名前が全て同じなので存在していると判断してidを返す
			return id
		} else {
			continue
		}
	}
	return 0
}
// rows.Scanとかで値を取得しない時はExecを使う
// QueryとかQueryRowを使用した際にScanを使用しないとconn busyというエラーがでる

func setUpCocktailCategoriesTable(conn *pgx.Conn) {
	arr := []string{
		cocktail_category_short,
		cocktail_category_long,
		cocktail_category_non_alcohol,
	}
	for _, name := range arr {
		if(checkExists(cocktail_categories_table, "cocktail_category_id","name", name, conn)!=0) {
			fmt.Println("already exists: ", name)
			continue
		}
		_, err := conn.Exec(
			context.Background(),
			"INSERT INTO cocktail_categories (name) VALUES ($1)",
			name,
		)
		if(err != nil) {
			fmt.Println("failed to insert: ", name)
			fmt.Println(err)
		}
	}
}

func setUpIngredientCategoriesTable(conn *pgx.Conn) {
	arr := []string{
		ingredient_category_spirit,
		ingredient_category_liqueur,
		ingredient_category_softdrink,
		ingredient_category_whisky,
		ingredient_category_brandy,
		ingredient_category_shothu,
		ingredient_category_japanese_style_alcohol,
	}
	for _, name := range arr {
		if(checkExists(ingredient_categories_table, "ingredient_category_id", "name", name, conn)!=0) {
			fmt.Println("already exists: ", name)
			continue
		}
		_, err := conn.Exec(
			context.Background(),
			"INSERT INTO ingredient_categories (name) VALUES ($1)",
			name,
		)
		if(err != nil) {
			fmt.Println("failed to insert: ", name)
			fmt.Println(err)
		}
	}
}

func getCocktailCategoryId(name string, conn *pgx.Conn) int {
	var id int
	conn.QueryRow(
		context.Background(),
		"SELECT cocktail_category_id FROM cocktail_categories WHERE name=$1",
		name).Scan(&id)
	return id
}

func getIngredientCategoryId(name string, conn *pgx.Conn) int {
	var id int
	conn.QueryRow(
		context.Background(),
		"SELECT ingredient_category_id FROM ingredient_categories WHERE name=$1",
		name).Scan(&id)
	return id
}

func insertCocktailParentId(conn *pgx.Conn) {
	for _, cocktail := range CocktailArr {
		parentName := cocktail.ParentName
		if(parentName == "") {
			// 親がないので何もしない
		} else {
			parentId := 0
			conn.QueryRow(
				context.Background(),
				"SELECT cocktail_id FROM cocktails WHERE name=$1",
				parentName,
			).Scan(&parentId)
			conn.Exec(
				context.Background(),
				"UPDATE cocktails SET parent_cocktail_id=$1 WHERE name=$2",
				parentId, cocktail.Name,
			)
			
		}
	}
}

func insertToCocktailIngredientsTable(cocktail_id int, ingredient_id int, conn *pgx.Conn) {
	fmt.Println("中間テーブルへ挿入 cocktail_id: ", cocktail_id, " ingredient_id: ", ingredient_id)
	cocktail_ingredient_id := 0
	conn.QueryRow(
		context.Background(),
		"INSERT INTO cocktail_ingredients (cocktail_id, ingredient_id) VALUES ($1, $2) RETURNING cocktail_ingredient_id",
		cocktail_id, ingredient_id,
	).Scan(&cocktail_ingredient_id)
}

func StartSetUp() {
	fmt.Println("setUp.go start")

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	// 先にカテゴリテーブルのデータを挿入
	setUpCocktailCategoriesTable(conn)
	setUpIngredientCategoriesTable(conn)

	for _, cocktail := range CocktailArr {
		cocktail_id := 0
		cocktail_category_id := 0
		name := ""
		if(checkExistsCocktail(cocktail,conn)!=0) {
			fmt.Println("already exists: ", cocktail.Name)
			continue
		}
		cocktail_category_id = getCocktailCategoryId(cocktail.Category, conn)
		fmt.Println("category_id: ", cocktail_category_id)
		conn.QueryRow(
			context.Background(),
			"INSERT INTO cocktails (name, description, cocktail_category_id, vol, ingredient_count) VALUES ($1, $2, $3, $4, $5) RETURNING cocktail_id, name",
			cocktail.Name, cocktail.Description, cocktail_category_id, cocktail.Vol, cocktail.IngredientCount,
		).Scan(&cocktail_id, &name)
		fmt.Println("inserted cocktail id: ", cocktail_id, "name: ", name)

		ingredients := cocktail.Ingredients
		fmt.Println("ingredients: ", ingredients)
		for _, ingredient := range ingredients {
			ingredient_id := checkExists(ingredient_table, "ingredient_id", "longname", ingredient.LongName, conn)
			longname := ""
			if(ingredient_id!= 0) {
				fmt.Println("already exists: ", ingredient.LongName)
				insertToCocktailIngredientsTable(cocktail_id, ingredient_id, conn)
				continue
			}
			ingredient_category_id := getIngredientCategoryId(ingredient.Category, conn)
			conn.QueryRow(
				context.Background(),
				"INSERT INTO ingredients (shortname, longname, description, vol, ingredient_category_id) VALUES ($1, $2, $3, $4, $5) RETURNING ingredient_id, longname",
				ingredient.ShortName, ingredient.LongName, ingredient.Description, ingredient.Vol, ingredient_category_id,
			).Scan(&ingredient_id, &longname)
			fmt.Println("inserted ingredient id: ", ingredient_id, "longname: ", longname)
			// ingredient_id_arr = append(ingredient_id_arr, ingredient_id)

			insertToCocktailIngredientsTable(cocktail_id, ingredient_id, conn)
		}

	}

	// 挿入するときにこの処理をしようとすると親になるカクテルが絶対先に処理しなくてはいけないようになるので数が増えると大変
	// 親のカクテルのIDを挿入する処理
	insertCocktailParentId(conn)
}