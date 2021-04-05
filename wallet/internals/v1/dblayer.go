package v1

import (
	v1 "wallet/pkg/api/v1"
	"github.com/jinzhu/gorm"
)

// MysqlLayer Application MYSQL layer
// Handles database processes
type MysqlLayer struct {
	Session *gorm.DB
}

// NewMysqlLayer create SQL layer
func NewMysqlLayer(session *gorm.DB) DatabaseHandler {
	return &MysqlLayer{Session: session}
}

// CreateWallet -> create a new wallet
func (db *MysqlLayer) CreateWallet(arg *v1.Wallet) (string ,error) {
	db.Session.AutoMigrate(&v1.Wallet{})
	operation := db.Session.Create(arg)
	return arg.ID, operation.Error
}

// GetWallet -> get a new wallet
func (db *MysqlLayer) GetWallet(arg *v1.GetOneWalletReq) (*v1.Wallet, error) {
	var w v1.Wallet

	if db.Session.Where(&v1.Wallet{ID: arg.WalletId}).First(&w).RecordNotFound() {
		return nil, gorm.ErrRecordNotFound
	}
	return &w, nil
}

// UpdateWallet -> update a new wallet
func (db *MysqlLayer) UpdateWallet(oldArg *v1.Wallet, newArg *v1.Wallet) error {
	return db.Session.Model(&oldArg).Updates(newArg).Error
}

// GetUserWallets -> get all wallets owned by a wallet
func (db *MysqlLayer) GetUserWallets(arg *v1.GetMyWalletsRequest) ([]v1.Wallet, error) {
	var ws []v1.Wallet
	err := db.Session.Where("user_id = ?", arg.UserId).Find(&ws).Error
	if err != nil {
		return nil, err
	}
	return ws, err
}
