package server

import (
	"cocktail_api/common"
	"cocktail_api/setup"
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

func responseAllIngredients(ctx *gin.Context) {
	ingredients := getAllIngredients()
	ctx.JSON(200, ingredients)
}

func responseCocktails(ctx *gin.Context) {
	cocktails := setup.CocktailArr
	ctx.JSON(200, cocktails)
}
