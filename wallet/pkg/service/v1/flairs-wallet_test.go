package v1

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"
	v1internals "wallet/internals/v1"
	"wallet/libs/setup"
	v1 "wallet/pkg/api/v1"

	"google.golang.org/grpc/metadata"

	"github.com/dgrijalva/jwt-go"
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

func TestCreateWallet_ok(t *testing.T) {
	clearWalletTable()
	ctx := context.Background()
	sqlLayer := v1internals.NewMysqlLayer(testDb)
	s := NewflairsWalletServer(sqlLayer)

	rq := &v1.NewWalletRequest{
		AccountBal: 0.00,
		Currency:   "1",
		LedgerBal:  0.00,
		Memo:       "This is a test wallet",
		Name:       "Test wallet",
		WalletType: 101,
	}
	got, err := s.AddNewWallet(ctx, rq)
	if err != nil {
		t.Errorf("flairWalletServer.CreateWallet_ok() error = %v, wantErr %v", err, "f")
		return
	}

	var w v1.Wallet
	testDb.Last(&w)

	if w.ID != got.ID {
		t.Errorf("flairWalletServer.CreateWallet_ok() = %v, want %v", got.ID, w.ID)

	}

}

func TestGetMyWallet_ok(t *testing.T) {
	clearWalletTable()
	ctx := context.Background()
	sqlLayer := v1internals.NewMysqlLayer(testDb)
	s := NewflairsWalletServer(sqlLayer)

	expirationTime := time.Now().Add(24 * 60 * time.Minute)

	claims := &Claims{
		UserID: "usered",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secrek_key"))

	rq := &v1.NewWalletRequest{
		AccountBal: 0.00,
		Currency:   "1",
		LedgerBal:  0.00,
		Memo:       "This is a test wallet",
		Name:       "Test wallet",
		WalletType: 101,
		UserId:     "usered",
	}

	md := metadata.Pairs("authorization", tokenString)
	ctx = metadata.NewIncomingContext(ctx, md)

	_, err = s.AddNewWallet(ctx, rq)
	if err != nil {
		t.Errorf("flairWalletServer.TestGetMyWallet_ok() failed because user could not be created with error  %v", err)
		return
	}

	got, err := s.GetMyWallets(ctx, &v1.GetMyWalletsRequest{UserId: "usered"})
	if err != nil {
		t.Errorf("flairWalletServer.TestGetMyWallet_ok() failed because user could not Get user got wallets returned error   %v", err)
		return
	}
	if got.Wallets[0].UserId != "usered" {
		t.Errorf("flairWalletServer.TestGetMyWallet_ok() error = %v, wantErr %v", got.Wallets[0].UserId, "usered")
	}
}

func TestTransact_ok(t *testing.T) {
	//clearWalletTable()
	ctx := context.Background()
	sqlLayer := v1internals.NewMysqlLayer(testDb)
	s := NewflairsWalletServer(sqlLayer)

	// createThe wallet
	testDb.Save(v1.Wallet{
		AccountBal: 0.00,}
	)
	// create perform transact req

	// Perform Transact

	// Test response
}

func clearWalletTable() {
	testDb.Exec(setup.ClearWalletTable)
}
