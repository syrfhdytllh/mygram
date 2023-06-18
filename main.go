package main

import (
	"mygram/database"
	router "mygram/routers"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	port := "8080"
	r.Run(":" + port)
}
