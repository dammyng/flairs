package v1

import v1 "transaction/pkg/api/v1"

// DatabaseHandler - Module's Database  interface
type DatabaseHandler interface {
	CreateTransaction(*v1.Transaction) error
	GetUserTransactions(*v1.GetMyTransactionsRequest) ([]v1.Transaction,  error)
}
