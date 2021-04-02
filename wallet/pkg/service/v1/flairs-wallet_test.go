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
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"

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

func TestCreateWallet(t *testing.T) {
	clearWalletTable()
	sqlLayer := v1internals.NewMysqlLayer(testDb)
	s := NewflairsWalletServer(sqlLayer)

	userId := "user1"
	wrongKey := "secret key two"
	md := metadata.Pairs("authorization", "")

	ctx := context.Background()
	emptyCtx := metadata.NewIncomingContext(ctx, md)

	wrongToken := "wrongtoken_obviously_wrong"
	md = metadata.Pairs("authorization", wrongToken)
	wrongCtx := metadata.NewIncomingContext(ctx, md)

	anotherUsersToken := createToken("user2", Key)
	md = metadata.Pairs("authorization", anotherUsersToken)
	anotherUserCtx := metadata.NewIncomingContext(ctx, md)

	wrongSignerToken := createToken("user2", wrongKey)
	md = metadata.Pairs("authorization", wrongSignerToken)
	wrongSignerCtx := metadata.NewIncomingContext(ctx, md)

	shouldWorkToken := createToken(userId, Key)
	md = metadata.Pairs("authorization", shouldWorkToken)
	shouldWorkCtx := metadata.NewIncomingContext(ctx, md)

	tests := map[string]struct {
		input          *v1.NewWalletRequest
		_context       context.Context
		expectedOutput *v1.AddWalletResponse
		expectedError  error
	}{
		"NoAuthHeader": {input: GenerateNewWallet(userId), _context: ctx, expectedOutput: nil, expectedError: NoAuthMetaDataError},
		"EmptyAuthHeader": {input: GenerateNewWallet(userId), _context: emptyCtx, expectedOutput: nil, expectedError: InvalidTokenError},
		"WrongToken": {input: GenerateNewWallet(userId), _context: wrongCtx, expectedOutput: nil, expectedError: WrongTokenStruct},
		"AuthorizationMisMatch": {input: GenerateNewWallet(userId), _context: anotherUserCtx, expectedOutput: nil, expectedError: InvalidTokenError},
		"WrongSigner": {input: GenerateNewWallet(userId), _context: wrongSignerCtx, expectedOutput: nil, expectedError: WrongTokenStruct},
		"ShouldWork": {input: GenerateNewWallet(userId), _context: shouldWorkCtx, expectedOutput: nil, expectedError: nil},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actualOutput, actualError := s.AddNewWallet(tc._context, tc.input)
			if actualError != nil{
				require.Equal(t, tc.expectedError.Error(), actualError.Error())
			}
			if actualOutput != nil{
				require.Len(t, actualOutput.ID , len(uuid.NewV4().String()))
			}
		})
	}
}


func TestGetWallet_ok(t *testing.T) {
	clearWalletTable()
	ctx := context.Background()
	sqlLayer := v1internals.NewMysqlLayer(testDb)
	s := NewflairsWalletServer(sqlLayer)
	testDb.Save(v1.Wallet{ID: "xxx-id"})
	rq := &v1.GetOneWalletReq{
		WalletId: "xxx-id",
	}
	_, err := s.GetWallet(ctx, rq)
	if err != nil {
		t.Errorf("flairWalletServer.CreateWallet_ok() error = %v, wantErr %v", err, "f")
		return
	}

	var w v1.Wallet
	testDb.Last(&w)

}

func TestGetMyWallet_ok(t *testing.T) {
	clearWalletTable()
	ctx := context.Background()
	sqlLayer := v1internals.NewMysqlLayer(testDb)
	s := NewflairsWalletServer(sqlLayer)

	userID := "user1"
	tokenString := createToken(userID, "")

	rq := &v1.NewWalletRequest{
		Currency: v1.Currency_NGR,
		Memo:     "This is a test wallet",
		Name:     "Test wallet",
		UserId:   "usered",
	}

	md := metadata.Pairs("authorization", tokenString)
	ctx = metadata.NewIncomingContext(ctx, md)

	_, err := s.AddNewWallet(ctx, rq)
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

func TestGetOneWallet_ok(t *testing.T) {
	clearWalletTable()
	ctx := context.Background()
	sqlLayer := v1internals.NewMysqlLayer(testDb)
	s := NewflairsWalletServer(sqlLayer)

	userID := "user1"
	tokenString := createToken(userID, "")

	rq := &v1.NewWalletRequest{
		Currency: v1.Currency_NGR,
		Memo:     "This is a test wallet",
		Name:     "Test wallet",
		Type:     101,
		UserId:   "usered",
	}

	md := metadata.Pairs("authorization", tokenString)
	ctx = metadata.NewIncomingContext(ctx, md)

	_, err := s.AddNewWallet(ctx, rq)
	if err != nil {
		t.Errorf("flairWalletServer.TestGetMyWallet_ok() failed because user could not be created with error  %v", err)
		return
	}

	_, err = s.GetWallet(ctx, &v1.GetOneWalletReq{WalletId: "usered"})
	if err != nil {
		t.Errorf("flairWalletServer.TestGetMyWallet_ok() failed because user could not Get user got wallets returned error   %v", err)
		return
	}
}

func GenerateNewWallet(userId string) *v1.NewWalletRequest {
	return &v1.NewWalletRequest{
		Currency: v1.Currency_NGR,
		Memo:     "This is a test wallet",
		Name:     "Test wallet",
		Type:     101,
		UserId:   userId,
	}
}

func createToken(userId, key string ) string {
	expirationTime := time.Now().Add(24 * 60 * time.Minute)

	claims := &Claims{
		UserID: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(key))

	if err != nil {
		log.Panicln(err)
	}
	return tokenString
}

/**
func TestTransact_ok(t *testing.T) {
	clearWalletTable()
	ctx := context.Background()
	sqlLayer := v1internals.NewMysqlLayer(testDb)
	s := NewflairsWalletServer(sqlLayer)

	// createThe wallet
	testDb.Save(v1.Wallet{
		AccountBal: 0.00,
		UserId:"userID",
		ID:"xxxx-id",
	})
	// create perform transact req
	 req:= &v1.PerformTransactionReq{
		 Amount:-99.10,
		 WalletID: "xxxx-id",
	 }
	// Perform Transact
	 _, err :=s.Transact(ctx, req)
	// Test response
	if err != nil {
		t.Errorf("flairWalletServer. TestTransact_ok() failed because user could not Get user got wallets returned error   %v", err)
		return
	}
	var w v1.Wallet
	testDb.Last(&w)

	if w.AccountBal != -99.10{
		t.Errorf("flairWalletServer. TestTransact_ok() failed account balance failed to sum up, expected 100 got  %v", w.AccountBal)
		return
	}
}
**/
func clearWalletTable() {
	testDb.Exec(setup.ClearWalletTable)
}
