package main

import (
	"log"
	"pismo-challenge/database"
	"pismo-challenge/services"
)

func main() {
	database.Connect()
	services.InitServices()
	gin := HandleRequests()
	err := gin.Run(":8080")
	if err != nil {
		log.Panic("Error starting GIN server", err)
	}
}
