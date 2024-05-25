package main

import (
	"fmt"
	"interfacesexplained/mysqldb"
	"log"
)

type dbcontract interface {
	Close()
	InsertUser(userName string) error
	SelectSingleUser(userName string) (string, error)
}

type Application struct {
	db dbcontract
}

func NewApplication(db dbcontract) *Application {
	return &Application{db: db}
}

func (app Application) Run() {
	userName := "user1"

	err := app.db.InsertUser(userName)
	if err != nil {
		log.Println(err)
	}

	user, err := app.db.SelectSingleUser(userName)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(user)
}

func main() {
	dbUser := "user"
	dbPassword := "password"
	dbHost := "host"
	dbPort := "port"
	dbName := "name"

	db, err := mysqldb.New(dbUser, dbPassword, dbHost, dbPort, dbName)
	if err != nil {
		log.Fatalf("failed to initialize db connection: %v", err)
	}
	defer db.Close()

	app := NewApplication(db)

	app.Run()
}
