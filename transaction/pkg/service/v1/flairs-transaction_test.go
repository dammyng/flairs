package v1

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	v1internals "transaction/internals/v1"
	"transaction/libs/setup"
	v1 "transaction/pkg/api/v1"

	_ "github.com/go-sql-driver/mysql"
	"github.com/streadway/amqp"

	"github.com/jinzhu/gorm"
	e_amqp "shared/events/amqp"

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

func TestAddnewTransaction_ok(t *testing.T) {
	clearTransactionTable()
	ctx := context.Background()
	sqlLayer := v1internals.NewMysqlLayer(testDb)
	s := NewflairsTransactionServer(sqlLayer, testEmitter)

	rq := &v1.NewTransactionReq{
		T_ID: "1695241",
	}
	os.Setenv("FlutterSecret","FLWSECK_TEST-be6475503d295c1be0b10ee8e971671f-X")
	got, err := s.AddnewTransaction(ctx, rq)
	if err != nil {
		t.Errorf("flairWalletServer. AddnewTransaction_ok() error = %v, wantErr %v", err, "f")
		return
	}

	var w v1.Transaction
	testDb.Last(&w)

	if w.ID != got.ID {
		t.Errorf("flairWalletServer.AddnewTransaction_ok() = %v, want %v", got.ID, w.ID)

	}
}

func initDB() {
	var err error
	dsn := fmt.Sprintf("root:password@tcp(127.0.0.1)/?charset=utf8&parseTime=True&loc=Local")
	testDb, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Panicf("flairTransaction DB setup() error = %v,", err)
	}
	testDb.Exec(setup.DropDB)
	testDb.Exec(setup.CreateDatabase)
	testDb.Exec(setup.UseAlphaWallet)
	testDb.Exec(setup.CreateWalletTable)
	testDb.Exec(setup.SQLMode)
}


func clearTransactionTable() {
	testDb.Exec(setup.ClearWalletTable)
}
