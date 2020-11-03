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
func (db *MysqlLayer) CreateWallet(arg *v1.Wallet) error {
	return db.Session.Create(arg).Error
}

