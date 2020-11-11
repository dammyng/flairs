package v1

import v1 "wallet/pkg/api/v1"

// DatabaseHandler - Module's Database  interface
type DatabaseHandler interface {
	CreateWallet(*v1.Wallet) error
	GetUserWallets(*v1.GetMyWalletsRequest) ([]v1.Wallet,  error)
}
