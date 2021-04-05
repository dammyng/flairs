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

var (
	userId   = "user1"
	wrongKey = "secret key two"
	md1      = metadata.Pairs("authorization", "")

	ctx      = context.Background()
	emptyCtx = metadata.NewIncomingContext(ctx, md1)

	wrongToken = "wrongtoken_obviously_wrong"
	md2        = metadata.Pairs("authorization", wrongToken)
	wrongCtx   = metadata.NewIncomingContext(ctx, md2)

	anotherUsersToken = createToken("user2", Key)
	md3               = metadata.Pairs("authorization", anotherUsersToken)
	anotherUserCtx    = metadata.NewIncomingContext(ctx, md3)

	wrongSignerToken = createToken("user2", wrongKey)
	md4              = metadata.Pairs("authorization", wrongSignerToken)
	wrongSignerCtx   = metadata.NewIncomingContext(ctx, md4)

	shouldWorkToken = createToken(userId, Key)
	md5             = metadata.Pairs("authorization", shouldWorkToken)
	shouldWorkCtx   = metadata.NewIncomingContext(ctx, md5)

	walletID    = "xxx-ddd-ccc"
	walletID_u1 = "xxx-ddd-ccc-u1"
	walletID_u2 = "xxx-ddd-ccc-u2"
)

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
	//testDb.Exec(setup.DropDB)
	testDb.Exec(setup.CreateDatabase)
	testDb.Exec(setup.UseAlphaWallet)
	testDb.Exec(setup.CreateWalletTable)
	testDb.Exec(setup.SQLMode)
}

func TestCreateWallet(t *testing.T) {
	clearWalletTable()
	sqlLayer := v1internals.NewMysqlLayer(testDb)
	s := NewflairsWalletServer(sqlLayer)

	tests := map[string]struct {
		input          *v1.NewWalletRequest
		_context       context.Context
		expectedOutput *v1.AddWalletResponse
		expectedError  error
	}{
		"NoAuthHeader":          {input: GenerateNewWallet(userId), _context: ctx, expectedOutput: nil, expectedError: NoAuthMetaDataError},
		"EmptyAuthHeader":       {input: GenerateNewWallet(userId), _context: emptyCtx, expectedOutput: nil, expectedError: InvalidTokenError},
		"WrongToken":            {input: GenerateNewWallet(userId), _context: wrongCtx, expectedOutput: nil, expectedError: WrongTokenStruct},
		"AuthorizationMisMatch": {input: GenerateNewWallet(userId), _context: anotherUserCtx, expectedOutput: nil, expectedError: InvalidTokenError},
		"WrongSigner":           {input: GenerateNewWallet(userId), _context: wrongSignerCtx, expectedOutput: nil, expectedError: WrongTokenStruct},
		"ShouldWork":            {input: GenerateNewWallet(userId), _context: shouldWorkCtx, expectedOutput: nil, expectedError: nil},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actualOutput, actualError := s.AddNewWallet(tc._context, tc.input)
			if actualError != nil {
				require.Equal(t, tc.expectedError.Error(), actualError.Error())
			}
			if actualOutput != nil {
				require.Len(t, actualOutput.ID, len(uuid.NewV4().String()))
			}
		})
	}
}

func TestGetWallet(t *testing.T) {
	clearWalletTable()
	ctx := context.Background()
	sqlLayer := v1internals.NewMysqlLayer(testDb)
	s := NewflairsWalletServer(sqlLayer)

	err := testDb.Save(v1.Wallet{ID: walletID, UserId: userId}).Error
	if err != nil {
		log.Panicln(err)
	}

	tests := map[string]struct {
		input          *v1.GetOneWalletReq
		_context       context.Context
		expectedOutput *v1.GetWalletResponse
		expectedError  error
	}{
		"NoAuthHeader":          {input: &v1.GetOneWalletReq{WalletId: walletID}, _context: ctx, expectedOutput: nil, expectedError: NoAuthMetaDataError},
		"EmptyAuthHeader":       {input: &v1.GetOneWalletReq{WalletId: walletID}, _context: emptyCtx, expectedOutput: nil, expectedError: InvalidTokenError},
		"WrongToken":            {input: &v1.GetOneWalletReq{WalletId: walletID}, _context: wrongCtx, expectedOutput: nil, expectedError: WrongTokenStruct},
		"WrongSigner":           {input: &v1.GetOneWalletReq{WalletId: walletID}, _context: wrongSignerCtx, expectedOutput: nil, expectedError: WrongTokenStruct},
		"AuthorizationMisMatch": {input: &v1.GetOneWalletReq{WalletId: walletID}, _context: anotherUserCtx, expectedOutput: nil, expectedError: UserIDClaimIDError},
		"ShouldWork":            {input: &v1.GetOneWalletReq{WalletId: walletID}, _context: shouldWorkCtx, expectedOutput: nil, expectedError: nil},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actualOutput, actualError := s.GetWallet(tc._context, tc.input)
			if actualError != nil {
				require.Equal(t, tc.expectedError.Error(), actualError.Error())
			}
			if actualOutput != nil {
				require.Equal(t, actualOutput.Result.ID, walletID)
				require.Equal(t, actualOutput.Result.UserId, userId)
			}
		})
	}
}

func TestGetMyWallets(t *testing.T) {
	clearWalletTable()
	sqlLayer := v1internals.NewMysqlLayer(testDb)
	s := NewflairsWalletServer(sqlLayer)

	user2 := "userid2"
	err := testDb.Save(v1.Wallet{ID: uuid.NewV4().String(), UserId: userId}).Error
	err = testDb.Save(v1.Wallet{ID: walletID, UserId: userId}).Error
	err = testDb.Save(v1.Wallet{ID: walletID_u1, UserId: userId}).Error
	err = testDb.Save(v1.Wallet{ID: uuid.NewV4().String(), UserId: user2}).Error

	if err != nil {
		log.Panicln(err)
	}

	tests := map[string]struct {
		input          *v1.GetMyWalletsRequest
		_context       context.Context
		expectedOutput *v1.WalletsResponse
		expectedError  error
	}{
		"NoAuthHeader": {input: &v1.GetMyWalletsRequest{UserId: userId}, _context: ctx, expectedOutput: nil, expectedError: NoAuthMetaDataError},
		"EmptyAuthHeader": {input: &v1.GetMyWalletsRequest{UserId: userId}, _context: emptyCtx, expectedOutput: nil, expectedError: InvalidTokenError},
		"WrongToken": {input: &v1.GetMyWalletsRequest{UserId: userId}, _context: wrongCtx, expectedOutput: nil, expectedError: WrongTokenStruct},
		"WrongSigner": {input: &v1.GetMyWalletsRequest{UserId: userId}, _context: wrongSignerCtx, expectedOutput: nil, expectedError: WrongTokenStruct},
		"AuthorizationMisMatch": {input: &v1.GetMyWalletsRequest{UserId: userId}, _context: anotherUserCtx, expectedOutput: nil, expectedError: UserIDClaimIDError},
		"ShouldWorkUser1": {input: &v1.GetMyWalletsRequest{UserId: userId}, _context: shouldWorkCtx, expectedOutput: nil, expectedError: UserIDClaimIDError},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actualOutput, actualError := s.GetMyWallets(tc._context, tc.input)
			if actualError != nil {
				require.Equal(t, tc.expectedError.Error(), actualError.Error())
			}
			if actualOutput != nil {
				require.Len(t, actualOutput.Wallets, 3)
			}
		})
	}
}

func TestGetUpdateWallet(t *testing.T) {
	clearWalletTable()
	sqlLayer := v1internals.NewMysqlLayer(testDb)
	s := NewflairsWalletServer(sqlLayer)

	tests := map[string]struct {
		input          *v1.UpdateWalletReq
		_context       context.Context
		expectedOutput *v1.UpdateWalletRes
		expectedError  error
	}{}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actualOutput, actualError := s.UpdateWallet(tc._context, tc.input)
			if actualError != nil {
				require.Equal(t, tc.expectedError.Error(), actualError.Error())
			}
			if actualOutput != nil {
			}
		})
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

func createToken(userId, key string) string {
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

func clearWalletTable() {
	testDb.Exec(setup.ClearWalletTable)
}
