package v1

import (
	"log"
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


// GetUserWallets -> get all wallets owned by a wallet
func (db *MysqlLayer) GetUserWallets(arg *v1.GetMyWalletsRequest) ([]v1.Wallet, error) {

	rows, err := db.Session.Model(&v1.Wallet{}).Where("user_id = ?", arg.UserId).Select("id,user_id").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var _w v1.Wallet
	var ws []v1.Wallet
	for rows.Next() {
		if err := rows.Scan(&_w.ID, &_w.UserId); err != nil {
			log.Fatalln(err.Error())
		}

		ws = append(ws, _w)
	}
	return ws, err
}
