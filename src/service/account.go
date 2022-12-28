package service

import (
	"github.com/airstrik/airstrik/src/entities"
	"github.com/airstrik/gobase/pkg/utils"
)

func GenerateAccountId() string {
	return utils.ShortUUID("ac")
}

func CreateAccount(accountId string, name string, email string, logo string) *entities.Account {
	account := &entities.Account{
		Id:         accountId,
		DomainName: accountId,
		Endpoint:   accountId,
		Name:       name,
		Email:      email,
	}

	return account
}
