package service

import (
	AdminError "github.com/airstrik/airstrik/pkg/error"
	utils2 "github.com/airstrik/airstrik/pkg/utils"
	"github.com/airstrik/airstrik/src/constant"
	"github.com/airstrik/airstrik/src/entities"
	"github.com/airstrik/gobase/pkg/config"
	"log"
)

func CreateUserByMobileNo(ctx *config.Context, mobile string, firstName string, lastName string) *entities.User {
	if !utils2.IsValidateMobileNo(mobile) {
		panic(AdminError.InvalidMobileNo())
	}
	// Send OTP and st
	ctx.DB.DB.CreateTable(&entities.User{})

	user := &entities.User{
		MobileVerified: false,
		EmailVerified:  false,
		Mobile:         mobile,
		FirstName:      firstName,
		LastName:       lastName,
		Status:         constant.UserInvited,
	}
	ctx.DB.Create(user)
	//ctx.DB.DB.Model(user).Update("CreatedAt", time.Now())
	log.Print(user)
	return user
}
