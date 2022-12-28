package entities

import (
	"github.com/airstrik/airstrik/src/constant"
	"github.com/airstrik/gobase/pkg/schema/system"
)

type User struct {
	system.BaseModel
	Id             string `gorm:"primary_key;AUTO_INCREMENT"`
	Name           string `gorm:"size:50"`
	FirstName      string `gorm:"size:25"`
	LastName       string `gorm:"size:25"`
	NickName       string `gorm:"size:50"`
	Password       string `gorm:"size:250"`
	ProfilePhoto   string `gorm:"size:50"`
	Email          string `gorm:"size:50"`
	Mobile         string `gorm:"size:20"`
	MobileVerified bool
	EmailVerified  bool
	Status         constant.UserStatus
}

type Group struct {
	system.BaseModel
	Id           string `gorm:"primary_key"`
	Name         string
	ParentGroup  string
	Password     string
	ProfilePhoto string
	Email        string
	Owner        string
}
