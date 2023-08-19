package main

import (
	"fmt"
	"cocktail_api/server"
	"cocktail_api/setup"
)

func main() {
	fmt.Println("main.go start")
	setup.StartSetUp()
	server.StartServer()
}