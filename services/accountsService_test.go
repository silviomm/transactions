package services

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"pismo-challenge/models/account"
	"testing"
)

func TestAccountExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	InitMocks(ctrl)
	defer ctrl.Finish()

	mockAccountsRepository.EXPECT().GetAccount(gomock.Any()).Return(&account.Account{})
	exists := Accounts.AccountExists(1)
	assert.True(t, exists)
}

func TestAccountDoesNotExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	InitMocks(ctrl)
	defer ctrl.Finish()
	
	mockAccountsRepository.EXPECT().GetAccount(gomock.Any()).Return(nil)
	exists := Accounts.AccountExists(1)
	assert.False(t, exists)
}
