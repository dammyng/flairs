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

func clearDB() {
	dbHandler := persistence.NewMysqlLayer(os.Getenv("DBConfig"))
	dbHandler.Session.Exec(setup.DropDB)
}
