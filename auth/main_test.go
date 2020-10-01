package main

import (
	"auth/libs/persistence"
	"auth/libs/setup"
	"context"
	"errors"
	"os"
	"shared/models/appuser"
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

func TestAddNewUserGRPC(t *testing.T) {

	correctUser = appuser.User{}

	tests := map[string]struct {
		input appuser.UserArg
		want  appuser.User
		err   error
	}{
		"correct":                {input: appuser.UserArg{UserPayload: &correctUser}, want: correctUser},
		"invalid_user":           {input: appuser.UserArg{UserPayload: &userWithoutEmail}, err: errors.New("")},
		"email_should_be_unique": {input: appuser.UserArg{UserPayload: &notUniqueEmail}, err: errors.New("")},
	}

	for name, tc := range tests {
		cxt := context.Background()
		t.Run(name, func(t *testing.T) {
			got, err := a.authServer.AddUser(cxt, &appuser.UserArg{UserPayload: &appuser.User{}})
			if err != nil && err != tc.err {
				t.Fatalf("Test could not create a user | expected: %v, got: %v", tc.err, err)
			}
			if got != nil {
				t.Fatalf("Test failed after creating a user | expected: %v, got: %v", tc.want, got)
			}

		})
	}
}

func TestAddNewUserHTTP(t *testing.T) {
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
