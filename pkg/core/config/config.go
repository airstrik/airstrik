package config

import (
	"encoding/json"
	"github.com/itsparser/airstrike/pkg/core/server"
	"github.com/itsparser/airstrike/pkg/core/utils/helper"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var envAmazonSecret = os.Getenv("_AMAZON_SECRET_")
var envAmazonClientId = os.Getenv("_AMAZON_CLIENT_ID_")
var envAmazonClientSecret = os.Getenv("_AMAZON_CLIENT_SECRET_")
var env = os.Getenv("_ENV_")
var Config = initConfig()

func initConfig() *Configuration {
	conf := &Configuration{fileLoadStatus: false}
	loadDBConfig(conf)
	return conf
}

//Database configuration for this application
func loadDBConfig(conf *Configuration) {
	conf.Session = server.NewDB("Postgres", "postgres", "localhost", "5432", "user",
		"pass", "db")
	db := conf.Session.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
}

type Configuration struct {
	fileLoadStatus bool
	Session        *gorm.DB
	config         map[string]interface{}
}

func (c *Configuration) Get(name string) interface{} {
	if c.fileLoadStatus == false {
		conf := loadConfiguration()
		c.config = conf.(map[string]interface{})
	}
	result := c.config[name]
	return result
}

func loadConfiguration() interface{} {
	var config interface{}
	env := os.Getenv("_ENV_")
	if env == "" {
		env = "Dev"
	}
	file, err := ioutil.ReadFile("config." + strings.ToLower(env) + ".json")
	if err != nil {
		helper.SuppressError(err)
	} else {
		if err = json.Unmarshal(file, &config); err != nil {
			helper.SuppressError(err, "Unmarshal the file failed from the config json")
		}
	}
	return config
}
