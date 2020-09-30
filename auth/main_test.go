package main

import (
	"auth/libs/persistence"
	"auth/libs/setup"
	"os"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a.InitDB()
	a.StartGRPC()
	a.InitHandler()
	code := m.Run()
	clearDB()
	os.Exit(code)
}

func TestAddNewUserGRPC()  {
	// Create golang user object
	// Call Auth server GRPC Method passing it as parameter
	// Test for all response possibility
}

func TestAddNewUserHTTP()  {
	// Create json byte of user
	// Pass as http body to route
	// Marshal to app user golang object
	// Pass as parameter to handler
	// Test for all response possibility
}

func clearDB() {
	dbHandler := persistence.NewMysqlLayer(os.Getenv("DBConfig"))
	dbHandler.Session.Exec(setup.DropDB)
}
