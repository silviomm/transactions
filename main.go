package main

import (
	"fmt"
	"log"
	"os"
	"pismo-challenge/database"
	"strconv"
)

func main() {
	database.Connect()
	gin := HandleRequests()
	port, _ := strconv.Atoi(os.Getenv("SERVER_PORT"))
	err := gin.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Panic("Error starting GIN server", err)
	}
}
