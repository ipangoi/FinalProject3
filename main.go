package main

import (
	"finalProject3/database"
	"finalProject3/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
