package main

import (
	"marine-ar-backend/routes"
)

func main() {
	router := routes.SetupRouter();
	router.Run(":8080")
}
