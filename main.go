package main

import (
	"github.com/airstrik/airstrik/src/route"
	"github.com/airstrik/gobase/pkg/server"
	"os"
)

var (
	Env string
	Srv *server.Server
)

func init() {
	Env = os.Getenv("ENV")
	Srv = server.GetServer("Airstrik", 9001)
	//Srv
}

func main() {
	// Register the Api for the Service from router
	Srv.RegisterRoute("/admin/{VersionId}", route.AdminRoute)
	Srv.RegisterRoute("/signin/{VersionId}", route.SignInRoute)
	Srv.RegisterRoute("/user/{VersionId}", route.UserRoute)
	// Public URL used in the Application
	Srv.RegisterRoute("/account", route.PublicRoute)

	Srv.Start()
}
