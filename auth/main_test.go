package main

import (
	"auth/grpc/authclient"
	"auth/libs/persistence"
	"auth/libs/setup"
	"context"
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

//Fetch user that do not exist
//
func TestCreateNewUser(t *testing.T) {
	clearUsersTable()
	cxt := context.Background()
	got, err := a.authServer.AddUser(cxt, &correctAppUser)
	if err != nil {
		t.Errorf("Test failed with an error '%v'", err.Error())
	}
	if got.ID != correctUser.ID {
		t.Errorf("Expected user details to be '%v'. Got '%v'", correctUser.ID, got.ID)
	}
	if got.Email != correctUser.Email {
		t.Errorf("Expected user details to be '%v'. Got '%v'", correctUser.Email, got.Email)
	}
	if got.DOB != correctUser.DOB {
		t.Errorf("Expected user details to be '%v'. Got '%v'", correctUser.DOB, got.DOB)
	}
}

func TestGetAUser(t *testing.T) {
	clearUsersTable()
	addDefaultUser()

	defaltUser := appuser.User{
		Email: "someone@flairs.com",
		ID:    "a65a388b-9c94-46f8-a99a-90c4807ce83b",
	}
	defaultAppUser := appuser.UserArg{UserPayload: &defaltUser}

	cxt := context.Background()
	got, err := a.authServer.GetUser(cxt, &defaultAppUser)

	if err != nil {
		t.Errorf("Test failed with an error '%v'", err.Error())
	}
	if got.ID != defaltUser.ID {
		t.Errorf("Expected user details to be '%v'. Got '%v'", defaultAppUser.UserPayload.ID, got.ID)
	}
	if got.Email != defaltUser.Email {
		t.Errorf("Expected user details to be '%v'. Got '%v'", defaultAppUser.UserPayload.Email, got.Email)
	}
}

func TestUpdateAUser(t *testing.T) {
	clearUsersTable()
	addDefaultUser()
	defaultUser := getDefaultUser()
	oldEmail := defaultUser.Email

	updateDefaltUser := appuser.User{
		Email: "some@flairs.com",
		ID:    "a65a388b-9c94-46f8-a99a-90c4807ce83b",
	}

	updater := appuser.UpdateArg{
		OldUser: &defaultUser,
		NewObj:  &updateDefaltUser,
	}

	err := a.DbHandler.UpdateUser(&updater)

	if err != nil {
		t.Errorf("Test failed with an error '%v'", err.Error())
	}

	if defaultUser.Email == oldEmail {
		t.Errorf("New Email is expected to be different. expected '%v' as new email. but it still equals '%v'", defaultUser.Email, oldEmail)
	}
}

func TestAdduserGRPClient(t *testing.T) {
	clearUsersTable()
	result, err := authclient.AddNewUser(&correctAppUser)
	if err != nil {
		t.Errorf("Test failed with an error '%v'", err.Error())
	}
	if result.ID != correctUser.ID {
		t.Errorf("Test failed expected ID %v got '%v'", correctUser.ID, result.ID)
	}
}

func TestGetuserGRPClient(t *testing.T) {
	clearUsersTable()
	addDefaultUser()
	defaltUser := appuser.User{
		Email: "someone@flairs.com",
		ID:    "a65a388b-9c94-46f8-a99a-90c4807ce83b",
	}
	defaultAppUser := appuser.UserArg{UserPayload: &defaltUser}

	result, err := authclient.GetAUser(&defaultAppUser)
	if err != nil {
		t.Errorf("Test failed with an error '%v'", err.Error())
	}
	if result.ID != defaltUser.ID {
		t.Errorf("Test failed expected ID %v got '%v'", defaltUser.ID, result.ID)
	}
}

func TestUpdateuserGRPClient(t *testing.T) {
	clearUsersTable()
	addDefaultUser()
	defaltUser := appuser.User{
		Email: "somes@flairs.com",
		ID:    "a65a388b-9c94-46f8-a99a-90c4807ce83b",
	}
	olduser := getDefaultUser()

	updater := appuser.UpdateArg{
		OldUser: &olduser,
		NewObj:  &defaltUser,
	}

	_, err := authclient.UpdateUser(&updater)
	if err != nil {
		t.Errorf("Test failed with an error '%v'", err.Error())
	}
	newuser := getDefaultUser()

	if olduser.Email == newuser.Email {
		t.Errorf("Email is expected to be different  but got the same %v | %v after update", newuser.Email, olduser.Email)
	}
}

func addDefaultUser() {
	dbHandler := persistence.NewMysqlLayer(os.Getenv("DBConnString"))
	dbHandler.Session.Exec(setup.InsertDemoUser)
}
func getDefaultUser() appuser.User {
	var users []appuser.User
	dbHandler := persistence.NewMysqlLayer(os.Getenv("DBConnString"))
	dbHandler.Session.Raw(setup.SelectDefaultUser).Scan(&users)
	return users[0]
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

func clearUsersTable() {
	dbHandler := persistence.NewMysqlLayer(os.Getenv("DBConnString"))
	dbHandler.Session.Exec(setup.ClearUserTable)
}
