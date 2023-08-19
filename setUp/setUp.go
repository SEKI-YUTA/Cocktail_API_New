package setup

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)


// 既に存在していた場合は０以外の値を返す
func checkExists(
	tableName string,
	colName string,
	name string,
	conn *pgx.Conn,
) int {
	fmt.Println("checkExists: ", tableName, name)
	var id int
	conn.QueryRow(
		context.Background(),
		"SELECT " + colName + " FROM " + tableName + " WHERE name=$1",
		name).Scan(&id)
	if(id == 0) {
		fmt.Println("not exists: ", name)
	} else {
		fmt.Println("exists: ", name)
	}
	return id
}

func insertToJoinTable(cocktail_id int, ingredient_id int, conn *pgx.Conn) {
	fmt.Println("中間テーブルへ挿入 cocktail_id: ", cocktail_id, " ingredient_id: ", ingredient_id)
			cocktail_ingredient_id := 0
			conn.QueryRow(
				context.Background(),
				"INSERT INTO cocktail_ingredients (cocktail_id, ingredient_id) VALUES ($1, $2) RETURNING cocktail_ingredient_id",
				cocktail_id, ingredient_id,
			).Scan(&cocktail_ingredient_id)
			fmt.Println("cocktail_ingredient_id: ", cocktail_ingredient_id)
}

func StartSetUp() {
	fmt.Println("setUp.go start")

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		fmt.Println("XXXXX")
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	for _, cocktail := range cocktailArr {
		// var cocktail_id int
		// var name string
		cocktail_id := 0
		name := ""
		fmt.Println("カクテル挿入 ", cocktail.Name, " ", cocktail.Description, " ", cocktail.Vol)
		if(checkExists(cocktail_table,"cocktail_id" ,cocktail.Name, conn)!=0) {
			fmt.Println("already exists: ", cocktail.Name)
			continue
		}
		rows := conn.QueryRow(
			context.Background(),
			"INSERT INTO cocktails (name, description, vol) VALUES ($1, $2, $3) RETURNING cocktail_id, name",
			cocktail.Name, cocktail.Description, cocktail.Vol,
		)
		rows.Scan(&cocktail_id, &name)
		fmt.Println("inserted cocktail id: ", cocktail_id, "name: ", name)

		ingredients := cocktail.Ingredients
		fmt.Println("ingredients: ", ingredients)
		for _, ingredient := range ingredients {
			ingredient_id := checkExists(ingredient_table, "ingredient_id", ingredient.Name, conn)
			if(ingredient_id!= 0) {
				fmt.Println("already exists: ", ingredient.Name)
				insertToJoinTable(cocktail_id, ingredient_id, conn)
				continue
			}
			conn.QueryRow(
				context.Background(),
				"INSERT INTO ingredients (name, description, vol) VALUES ($1, $2, $3) RETURNING ingredient_id, name",
				ingredient.Name, ingredient.Description, ingredient.Vol,
			).Scan(&ingredient_id, &name)
			fmt.Println("inserted ingredient id: ", ingredient_id, "name: ", name)
			// ingredient_id_arr = append(ingredient_id_arr, ingredient_id)

			insertToJoinTable(cocktail_id, ingredient_id, conn)
		}

		fmt.Println("\n")
	}
}