package route

import (
	"log"
	"net/http"
)

// CreateUser : Create new User for the This Account
func CreateUser(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context().Value("AccountCtx").(*config.Context)
	log.Print("Create User in Account")
	//user := r.Context().Value("Body").(interface{})
	//log.Print("Creating the User in db", user)
	//service.CreateNewUser("")

}

// GetUser: Get the User Subscription for this Account
func GetUser(w http.ResponseWriter, r *http.Request) {

}

// DeleteUser: Delete the User Subscription for this Account
func DeleteUser(w http.ResponseWriter, r *http.Request) {

}

// UpdateUser: Update the User Subscription for this Account
func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

// GetUserList : Get all active User in the Account from the process
func GetUserList(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context().Value("AccountCtx").(*config.Context)
	//
}
