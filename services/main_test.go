package services

import (
	"github.com/golang/mock/gomock"
	"os"
	"pismo-challenge/database/repositories"
	"pismo-challenge/mocks"
	"testing"
)

var mockCtrl *gomock.Controller
var mockAccountsRepository *mocks.MockAccountsRepository
var mockTransactionsRepository *mocks.MockTransactionsRepository

func TestMain(m *testing.M) {
	mockCtrl = gomock.NewController(nil)
	defer mockCtrl.Finish()

	InitMocks(mockCtrl)
	InitServices()

	code := m.Run()
	os.Exit(code)
}

func InitMocks(ctrl *gomock.Controller) {
	mockAccountsRepository = mocks.NewMockAccountsRepository(ctrl)
	repositories.Accounts = mockAccountsRepository
	mockTransactionsRepository = mocks.NewMockTransactionsRepository(ctrl)
	repositories.Transactions = mockTransactionsRepository
}
