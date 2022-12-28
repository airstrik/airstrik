package route

import (
	"encoding/json"
	"github.com/airstrik/airstrik/src/service"
	"github.com/airstrik/gobase/pkg/config"
	error2 "github.com/airstrik/gobase/pkg/server/error"
	"github.com/airstrik/gobase/pkg/utils"
	"net/http"
)

type SignUp struct {
	MobileNo  string
	EmailId   string
	FirstName string
	LastName  string
}

func SignUpMobile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context().Value("AccountCtx").(*config.Context)
	var signUp SignUp
	err := json.NewDecoder(r.Body).Decode(&signUp)
	utils.HandleCustomError(err, error2.JsonParseFailure)
	user := service.CreateUserByMobileNo(ctx, signUp.MobileNo, signUp.FirstName, signUp.LastName)
	json.NewEncoder(w).Encode(user)
}

func SignUpWithEmail(w http.ResponseWriter, r *http.Request) {
	//Check is this Mobile No or EMail based Sign in

}

func SignUserNamePassword(w http.ResponseWriter, r *http.Request) {

}

func SignGoogle(w http.ResponseWriter, r *http.Request) {

}

func CallbackGoogle(w http.ResponseWriter, r *http.Request) {

}
