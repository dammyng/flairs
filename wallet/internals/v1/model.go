package v1

import "time"

type Wallet struct {
	ID            string    `json:"id" gorm:"type:varchar(255); not null"`
	UserId        string    `json:"user_id" gorm:"not null"`
	AccountBal    string    `json:"available_balance"`
	LedgerBal     string    `json:"ledger_balance"`
	WalletType    string    `gorm:"type:varchar(1);" json:"wallet_type"`
	TermID        string    `gorm:"type:varchar(1);" json:"term_id"`
	Name          string    `gorm:"type:varchar(255);" json:"name"`
	Memo          string    `gorm:"type:varchar(255);" json:"memo"`
	Currency      string    `gorm:"type:varchar(20);not null" json:"currency"`
	Status        bool      `json:"status"`
	DateCreated   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DateBalUpdate time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	LastUpdate    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
