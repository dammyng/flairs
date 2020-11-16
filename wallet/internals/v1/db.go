package v1

import v1 "wallet/pkg/api/v1"

// DatabaseHandler - Module's Database  interface
type DatabaseHandler interface {
	CreateWallet(*v1.Wallet) error
	GetWallet(*v1.GetOneWalletReq) (v1.Wallet, error)
	GetUserWallets(*v1.GetMyWalletsRequest) ([]v1.Wallet,  error)
	UpdateWallet(*v1.Wallet, *v1.Wallet) (error)
}
