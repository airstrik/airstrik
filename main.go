package main

import (
	"encoding/json"
	"fmt"
	"github.com/itsparser/airstrike/pkg/core/model/engine"
	engine2 "github.com/itsparser/airstrike/pkg/engine"
	"github.com/tebeka/selenium"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//testinmg

	// Start a Selenium WebDriver server instance (if one is not already
	// running).
	const (
		// These paths will be different on your system.
		seleniumPath    = "vendor/selenium-server.jar"
		geckoDriverPath = "vendor/geckodriver"
		port            = 8080
	)
	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
		selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.Output(os.Stderr),            // Output debug information to STDERR.
	}
	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()
	//engine := server.ApplicationServer("Blocker", 9000)
	//user := &model.User{
	//	Name:  "Vasanth",
	//	Email: "",
	//	Group: []model.Group{
	//		model.Group{Name: "Group"},
	//	},
	//}
	//config.Config.Session.Save(user)
	//config.Config.Session.CreateTable(&model.User{})
	//group := &model.UserGroup{Name: "vasanth"}
	//config.Config.Session.Create(&group)


	//grp := engine.Server.Group("/api/1/")
	//grp.GET("/", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"v1-auth": true,
	//	})
	//})
	//engine.Run()
	jsonFile, err := os.Open("data/sample.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var suit engine.Suit
	err = json.Unmarshal(byteValue, &suit)
	if err!=nil {
		log.Fatal(err)
	}
	engine2.ProcessSuit(suit)
	fmt.Print(suit)
}
