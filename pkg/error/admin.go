package AdminError

import error2 "github.com/airstrik/gobase/pkg/server/error"

var err error2.BaseError

func InvalidMobileNo() error2.BaseError {
	err = &error2.Response{
		ErrorMessage: "Mobile Number provided is Invalid",
		LangMessage:  "Mobile Number provided is Invalid",
		ErrorCode:    "ERR001001",
	}
	return err
}
