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

// CreateTransaction -> create a new transaction
func (db *MysqlLayer) CreateTransaction(arg *v1.Transaction) error {
	return db.Session.Create(arg).Error
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
