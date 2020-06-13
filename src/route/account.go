package route

import (
	"encoding/json"
	"github.com/airstrik/airstrik/src/service"
	"github.com/airstrik/gobase/pkg/config"
	"github.com/airstrik/gobase/pkg/utils"
	"log"
	"net/http"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context().Value("AccountCtx").(*config.Context)
	log.Print("Log from the Get Account")

	// Extract the Payload from the body
	email, err := r.Body.Read([]byte("Email"))
	_email := string(email)
	utils.HandleError(err)
	_AccountId := service.GenerateAccountId()
	// Create Tenant DB and Run all the Index for that Database
	ctx.BaseDB.CreateDB(_AccountId)
	// TODO: Run all the Index for the New Tenant Database

	username, domain := utils.EmailDomainExtractor(_email)
	account := service.CreateAccount(_AccountId, username, _email, domain)
	ctx.BaseDB.Create(account)
	// Create new User subscription for that Account
	//_ = service.CreateNewUser(_email)
	_ = json.NewEncoder(w).Encode(utils.GenerateSuccessResponse())
}

// Get the Account from the database
func GetAccount(w http.ResponseWriter, r *http.Request) {
	log.Print("Log from the Get Account")

}
