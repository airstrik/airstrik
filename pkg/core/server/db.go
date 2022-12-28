package server

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // This specifies the Dialect to get declare
	"log"
)

type DB struct {
	Name       string
	dialect    string
	host       string
	port       string
	user       string
	password   string
	dbName     string
	Connection *gorm.DB
}

//Create New Database connection for the Airstrike Application to Connect
//Store and Process the data
func NewDB(name string, dialect string, host string, port string, user string, password string, dbName string) *gorm.DB {
	var con *gorm.DB
	var err error
	conUrl := "host=" + host + " port=" + port + " user=" + user + " dbname=" + dbName + " password=" + password + " sslmode=disable"
	con, err = gorm.Open(dialect, conUrl)
	if err != nil {
		log.Print("Error from the Database creation")
		log.Fatal(err)
	}
	return con
}

func (database *DB) Postgres() {

	//database.Connection, _ = gorm.Open("postgres",
	//	"postgres://user:pass@localhost:5432/blocker")
	//"host=localhost port=5432 user=user dbname=blocker password=pass")
	//defer database.Connection.Close()
}
