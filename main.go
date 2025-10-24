package main

import (
	"multi-tenant/router"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	router := router.NewRouter()

	router.Run(":3000")
}
