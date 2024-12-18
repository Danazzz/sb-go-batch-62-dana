package main

import (
	"formative-13/config"
	"formative-13/routers"
)

func main() {

	db := config.ConnectDB()
	defer db.Close()

	routers.StartServer().Run(":8080")
}