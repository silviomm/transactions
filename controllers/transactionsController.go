package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pismo-challenge/database"
	"pismo-challenge/models/transaction"
	"pismo-challenge/services"
)

func PostTransaction(c *gin.Context) {
	var dto transaction.CreateTransactionDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.Transactions.ValidateTransactionDto(dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t, err := services.Transactions.CreateTransaction(&dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if services.Transactions.ShouldDischarge(t.OperationType) {
		services.Transactions.Discharge(t)
	}

	c.JSON(http.StatusOK, t.Id)
}

// todo: should be paginated
func GetTransactions(c *gin.Context) {
	var dto transaction.GetTransactionsDto
	if err := c.ShouldBindQuery(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var ts []transaction.Transaction
	query := database.DB

	if dto.AccountId != 0 {
		query = query.Where("transactions.\"AccountId\"=?", dto.AccountId)
	}
	if dto.OperationTypeId != 0 {
		query = query.Where("transactions.\"OperationType\"=?", dto.OperationTypeId)
	}
	if !dto.Before.IsZero() {
		query = query.Where("transactions.\"EventDate\" < ?", dto.Before)
	}
	if !dto.After.IsZero() {
		query = query.Where("transactions.\"EventDate\" > ?", dto.After)
	}

	query.Find(&ts)
	c.JSON(http.StatusOK, ts)
}
