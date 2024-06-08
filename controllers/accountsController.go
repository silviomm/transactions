package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pismo-challenge/database"
	"pismo-challenge/models/account"
	"pismo-challenge/services"
)

func GetAccount(c *gin.Context) {
	var a account.Account
	id := c.Params.ByName("accountId")
	database.DB.First(&a, id)

	if a.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"result": "account not found"})
		return
	}

	dto := account.GetAccountDto{
		Id:             a.Id,
		DocumentNumber: a.DocumentNumber,
	}
	c.JSON(http.StatusOK, dto)
}

func PostAccount(c *gin.Context) {
	var dto account.CreateAccountDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if services.ExistsByDocumentNumber(dto.DocumentNumber) {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Document number '%s' already have an account", dto.DocumentNumber)})
		return
	}
	a := account.Account{DocumentNumber: dto.DocumentNumber}
	database.DB.Create(&a)
	c.JSON(http.StatusOK, gin.H{"account_id": a.Id})
}
