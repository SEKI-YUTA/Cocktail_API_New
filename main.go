package main

import (
	"fmt"
	"cocktail_api/server"
	// docker-compose upの時にデータを入れるように変更したのでセットアップの処理を走らせる必要がなくなった
	// "cocktail_api/setup"
)

func main() {
	fmt.Println("main.go start")
	// docker-compose upの時にデータを入れるように変更したのでセットアップの処理を走らせる必要がなくなった
	// setup.StartSetUp()
	server.StartServer()
}