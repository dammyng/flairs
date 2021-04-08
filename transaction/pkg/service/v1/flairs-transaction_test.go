package v1

import (
	"fmt"
	"log"
	"os"
	"testing"
	"transaction/libs/setup"
	v1 "transaction/pkg/api/v1"

	_ "github.com/go-sql-driver/mysql"
	"github.com/streadway/amqp"

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


func TestAddnewTransaction(t *testing.T) {
	clearTransactionTable()
	testDb.AutoMigrate(&v1.Transaction{})

	os.Setenv("FlutterSecret", "FLWSECK_TEST-be6475503d295c1be0b10ee8e971671f-X")
}

func TestAddnewTransaction_ok_case1(t *testing.T) {
	clearTransactionTable()
	
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
