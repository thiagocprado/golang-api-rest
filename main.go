package main

import (
	"api-rest-golang/database"
	"api-rest-golang/routes"
	"fmt"
)

func main() {
	database.ConnectDB()
	fmt.Println("Starting server on port 8080...")
	routes.HandleRequest()
}
