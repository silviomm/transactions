package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pismo-challenge/database"
	"pismo-challenge/database/repositories"
	"pismo-challenge/models/account"
	"strconv"
)

func GetAccount(c *gin.Context) {
	var a *account.Account
	id, _ := strconv.Atoi(c.Params.ByName("accountId"))
	a = repositories.Accounts.GetAccount(id)

	if a == nil {
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

	// This business rule isn't specified in the pdf, so I commented it. But can be done like this:
	//if services.ExistsByDocumentNumber(dto.DocumentNumber) {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Document number '%s' already have an account", dto.DocumentNumber)})
	//	return
	//}

	a := account.Account{DocumentNumber: dto.DocumentNumber}
	database.DB.Create(&a)
	c.JSON(http.StatusOK, gin.H{"account_id": a.Id})
}
