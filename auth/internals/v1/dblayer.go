package v1

import (
	v1 "auth/pkg/api/v1"
	"errors"

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

// FindUser Find a user in DB,
// Returns Not found as error
//
func (db *MysqlLayer) FindUser(arg *v1.User) (*User, error) {
	var user User
	err := db.Session.Where(arg).First(&user).Error
	if errors.Is( err,  gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &user, err
}

// CreateUser -> create a new user
func (db *MysqlLayer) CreateUser(arg *v1.User) error {
	return db.Session.Create(arg).Error
}

// UpdateUser -> Update User
// update a new user with a model user object
func (db *MysqlLayer) UpdateUser(old *v1.User, new *v1.User, ) error {
	return db.Session.Model(&old).Update(&new).Error
}
// UpdateUserMap -> Update user with map
// update a new user with dynamics map object
func (db *MysqlLayer) UpdateUserMap(arg *v1.User, dict map[string]interface{}) error {
	return db.Session.Model(&arg).Update(dict).Error
}

