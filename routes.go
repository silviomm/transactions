package main

import (
	"github.com/gin-gonic/gin"
	"pismo-challenge/controllers"
)

func HandleRequests() *gin.Engine {
	r := gin.Default()

	r.GET("/accounts/:accountId", controllers.GetAccount)
	r.POST("/accounts", controllers.PostAccount)
	r.POST("/transactions", controllers.PostTransaction)
	r.GET("/transactions", controllers.GetTransactions)

	return r
}
