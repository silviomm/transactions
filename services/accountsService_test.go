package services

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"pismo-challenge/database/repositories"
	"pismo-challenge/mocks"
	"pismo-challenge/models/account"
	"testing"
)

func TestAccountExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAccountsRepository(ctrl)
	repositories.Accounts = mockRepo

	mockRepo.EXPECT().GetAccount(gomock.Any()).Return(&account.Account{})
	exists := AccountExists(1)
	assert.True(t, exists)
}

func TestAccountDoesNotExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAccountsRepository(ctrl)
	repositories.Accounts = mockRepo

	mockRepo.EXPECT().GetAccount(gomock.Any()).Return(nil)
	exists := AccountExists(1)
	assert.False(t, exists)
}
