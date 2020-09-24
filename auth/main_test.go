package main_test

import (
	"flairs/auth/libs/config"
	"flairs/auth/libs/persistence"
	"flairs/auth/rest"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

var serviceHandler rest.ServiceHandler
var dbHandler *persistence.MysqlLayer

func TestMain(m *testing.M)  {
	initDB()
	//initUserTable()
	//setUpRouter()
	code:= m.Run()
	//clearDB()
	os.Exit(code)
}

func initDB() {
	dbHandler = persistence.NewMysqlLayer(config.DBConfig)
	dbHandler.Session.Exec(config.CreateDatabase)
	dbHandler.Session.Exec(config.UseAlphaPlus)
	dbHandler.Session.Exec(config.CreateUserTable)
}

func setUpRouter() *mux.Router {
	return rest.ServerRoute(dbHandler)
}

func clearDB()  {
	dbHandler.Session.Exec(config.DropDB)
}


func TestEmptyTable(t *testing.T)  {
	clearTable()
	req, _ := http.NewRequest("GET", "/auth/users", nil)
	response := makeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	if body := response.Body.String(); body !="[]"{
		t.Errorf("Expected an empty array []. Got: %s ", body)
	}
}

func clearTable() {
    dbHandler.Session.Exec("DELETE FROM users")
    dbHandler.Session.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1")
}

func makeRequest(req * http.Request) *httptest.ResponseRecorder {
	rr:= httptest.NewRecorder()
	setUpRouter().ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }
}