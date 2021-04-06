package v1

import (
	"log"
	v1 "transaction/pkg/api/v1"
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

// GetTransaction -> get a new transaction
func (db *MysqlLayer) GetTransaction(arg *v1.Transaction) (*v1.Transaction, error) {
	var t v1.Transaction

	if db.Session.Where(&v1.Transaction{ID: arg.ID}).First(&t).RecordNotFound() {
		return nil, gorm.ErrRecordNotFound
	}
	return &t, nil
}

// CreateTransaction -> create a new transaction
func (db *MysqlLayer) CreateTransaction(arg *v1.Transaction) (string, error) {
	err := db.Session.Create(arg).Error
	return arg.ID, err
}

// GetWalletTransactions -> get all transactions in a wallet
func (db *MysqlLayer) GetWalletTransactions(arg *v1.Transaction) ([]v1.Transaction, error) {
	var ts []v1.Transaction
	err := db.Session.Where("wallet_id = ?", arg.WalletId).Find(&ts).Error
	if err != nil {
		return nil, err
	}
	return ts, err
}


// GetUserTransactions -> get all transaction owned by a user
func (db *MysqlLayer) GetUserTransactions(arg *v1.GetMyTransactionsRequest) ([]v1.Transaction, error) {

	rows, err := db.Session.Model(&v1.Transaction{}).Where("user_id = ?", arg.UserId).Select("id,user_id").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var _w v1.Transaction
	var ws []v1.Transaction
	for rows.Next() {
		if err := rows.Scan(&_w.ID, &_w.CustomerId); err != nil {
			log.Fatalln(err.Error())
		}

		ws = append(ws, _w)
	}
	return ws, err
}
