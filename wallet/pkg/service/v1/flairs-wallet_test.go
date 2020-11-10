package v1

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	v1internals "wallet/internals/v1"
	"wallet/libs/setup"
	v1 "wallet/pkg/api/v1"
	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

var testDb *gorm.DB

func TestMain(m *testing.M) {

	initDB()
	code := m.Run()
	//clearDB()
	os.Exit(code)
	//testDb.Exec(setup.DropDB)

}
func initDB() {
	var err error
	dsn := fmt.Sprintf("root:password@tcp(127.0.0.1)/?charset=utf8&parseTime=True&loc=Local")
	testDb, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Panicf("flairServiceServer.AddNewUser() error = %v,", err)
	}
	testDb.Exec(setup.DropDB)
	testDb.Exec(setup.CreateDatabase)
	testDb.Exec(setup.UseAlphaWallet)
	testDb.Exec(setup.CreateWalletTable)
	testDb.Exec(setup.SQLMode)
}

func TestCreateWallet_ok(t *testing.T)  {
	clearWalletTable()
	ctx := context.Background()
	sqlLayer := v1internals.NewMysqlLayer(testDb)
	s := NewflairsWalletServer(sqlLayer)

	rq := &v1.NewWalletRequest{
		AccountBal: 0.00,
		Currency: "1",
		LedgerBal: 0.00,
		Memo: "This is a test wallet",
		Name: "Test wallet",
		WalletType: 101,
	}
	got, err := s.AddNewWallet(ctx, rq)
	if err != nil {
		t.Errorf("flairWalletServer.CreateWallet_ok() error = %v, wantErr %v", err, "f")
		return
	}

	var w v1.Wallet
	testDb.Last(&w)

	if w.ID != got.ID{
		t.Errorf("flairWalletServer.CreateWallet_ok() = %v, want %v", got.ID, w.ID)

	}

}


func clearWalletTable() {
	testDb.Exec(setup.ClearWalletTable)
}