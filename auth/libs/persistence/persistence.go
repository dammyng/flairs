package persistence

import (
	"shared/models/appuser"

	"github.com/jinzhu/gorm"
)

//DBConfig Application DBConfig
type DBConfig struct {
	Hosts string
	Database string
	Username string
	Password string
	Port string
}
// DatabaseHandler - module app interface
type DatabaseHandler interface{
	AddUser(appuser.User) error
	AllUsers() ([]appuser.User, error)
	GetUser(*appuser.User) (*appuser.User, error)
	UpdateUser(*appuser.UpdateArg)  error
	FindUsers() ([]appuser.User, error)
	//
	GetFreshSession() *gorm.DB
	DoMigrations()
	Close()
	//
	AddCardRequest(appuser.CardRequest) error
	FindCardRequestById(string) (appuser.CardRequest, error)
	FindUserCardRequests(string) ([]appuser.CardRequest, error)
	FindWalletById(string) (appuser.Wallet, error)
	FindUserWallets(string) ([]appuser.Wallet, error)
}