package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"
	v1internals "transaction/internals/v1"
	"transaction/libs/setup"
	v1 "transaction/pkg/api/v1"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/streadway/amqp"
	"google.golang.org/grpc/metadata"

	e_amqp "shared/events/amqp"

	"github.com/jinzhu/gorm"
)

var testDb *gorm.DB
var testEmitter e_amqp.EventEmitter

func TestMain(m *testing.M) {
	initDB()
	initAMQP()
	code := m.Run()
	os.Exit(code)
}

func initAMQP() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("could not establish amqp connection :" + err.Error())
	}

	testEmitter, err = e_amqp.NewAMQPEventEmitter(conn, "auth")
	if err != nil {
		log.Fatal("could not establish amqp connection :" + err.Error())
	}
}


func TestAddnewTransaction_ok_case0(t *testing.T) {
	clearTransactionTable()
	ctx := context.Background()
	sqlLayer := v1internals.NewMysqlLayer(testDb)
	s := NewflairsTransactionServer(sqlLayer, testEmitter)


	os.Setenv("FlutterSecret", "FLWSECK_TEST-be6475503d295c1be0b10ee8e971671f-X")
}

func TestAddnewTransaction_ok_case1(t *testing.T) {
	clearTransactionTable()
	//ctx := context.Background()
	//sqlLayer := v1internals.NewMysqlLayer(testDb)
	//s := NewflairsTransactionServer(sqlLayer, testEmitter)

	
	reqURL, _ := url.Parse("http://localhost:9000/v1/wallet")

	// create request body
	bodyContent := `{"userId":"usered", "accountBal":"0.54"}`

	reqBody := ioutil.NopCloser(strings.NewReader(bodyContent))

	req := &http.Request{
		Method: "POST",
		URL:    reqURL,
		Header: map[string][]string{
			"Content-Type":  {"application/json; charset=UTF-8"},
			"Authorization": {""},
		},
		Body: reqBody,
	}

	_, err := HttpReq(req)
	if err != nil {
		t.Errorf("flairWalletServer. AddnewTransaction_ok cse 1 () could not create test wallet  error = %v, wantErr %v", err, "f")
		return
	}
}

func initDB() {
	var err error
	dsn := fmt.Sprintf("root:password@tcp(127.0.0.1)/?charset=utf8&parseTime=True&loc=Local")
	testDb, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Panicf("flairTransaction DB setup() error = %v,", err)
	}
	//testDb.Exec(setup.DropDB)
	testDb.Exec(setup.CreateDatabase)
	testDb.Exec(setup.UseAlphaTransaction)
	testDb.Exec(setup.CreateTransactionTable)
	testDb.Exec(setup.SQLMode)
}

func clearTransactionTable() {
	testDb.Exec(setup.ClearTransactionTable)
}
