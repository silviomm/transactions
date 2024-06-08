package main

import (
	"log"
	"pismo-challenge/database"
)

func main() {
	database.Connect()
	gin := HandleRequests()
	err := gin.Run(":8080")
	if err != nil {
		log.Panic("Error starting GIN server", err)
	}
}
