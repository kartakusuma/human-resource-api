package main

import "human-resource-api/routes"

func main() {
	router := routes.RouterSetup()

	router.Run("8080")
}
