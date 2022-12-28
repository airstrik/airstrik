package entities

import (
	"github.com/airstrik/gobase/pkg/schema/system"
)

type Account struct {
	system.BaseModel
	Id         string `gorm:"primary_key"`
	DomainName string
	Endpoint   string
	Name       string
	Logo       string
	Email      string
	Owner      string
}
