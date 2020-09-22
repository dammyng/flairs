package persistence

import "shared/models/appuser"

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
}