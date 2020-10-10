package v1

import v1 "auth/pkg/api/v1"

// DatabaseHandler - Module's Database  interface
type DatabaseHandler interface {
	FindUser(*v1.User) (*User, error)
	CreateUser(*v1.User) error
	UpdateUser(*v1.User, *v1.User) error
	UpdateUserMap(*v1.User, map[string]interface{}) error
}
