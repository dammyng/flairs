package sqlmodel

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ACCOUNT_TYPE string
type REFERRAL_TYPE string

const (
	PERSONAL ACCOUNT_TYPE = "personal"
	BUSINESS ACCOUNT_TYPE = "bussiness"
)

const (
	REGISTRATION REFERRAL_TYPE = "registration"
)

type User struct {
	// id
	ID string `gorm:"primary_key"`
	// first name
	FirstName string `json:"first_name" gorm:"size:255;"`
	// last name
	LastName string `json:"last_name" gorm:"size:255;"`
	// address
	Address string `json:"address" gorm:"size:255;"`
	// street
	Street string `json:"street" gorm:"size:255;"`
	// city
	City string `json:"city" gorm:"size:255;"`
	// postal code
	PostalCode string `json:"postal_code" gorm:"size:255;"`
	// state
	State string `json:"state" gorm:"size:255;"`
	// country
	Country string `json:"country" gorm:"size:255;"`
	// referrer
	Referrer string `json:"referrer" gorm:"size:255;"`
	// refcode
	RefCode                 string `json:"refcode" gorm:"size:255;"`
	How_did_u_hear_about_us string `json:"how_did_u_hear_about_us" gorm:"size:255;"`

	Username string `json:"username" gorm:"type:varchar(30);"`

	LastCardRequest string `json:"last_card_request" gorm:"size:255;"`
	Passport        string `json:"passport" sql:"size:999999" `
	IDCard          string `json:"id_card" sql:"size:999999"`

	BVN      string `json:"bvn" gorm:"type:varchar(11)"`
	Wallet   int64  `json:"wallet" gorm:"size:255"`
	Customer int64  `json:"customer" gorm:"size:255"`
	//	Cards    string `json:"cards" gorm:"type:varchar(255)"`
	Cards           []Card    `json:"cards"`
	Gender          string    `json:"gender" gorm:"type:varchar(20)"`
	DOB             time.Time `json:"dob"`
	Email           string    `json:"email" gorm:"type:varchar(100);unique_index;not null"`
	EmailVerifiedAt time.Time `json:"email_verified_at"`
	Password        []byte    `json:"-"`
	Pin             []byte    `json:"-"`

	IsProfileComplete bool `json:"is_profile_complete" gorm:"type:boolean;column:is_profile_complete;default:false"`

	PhoneNumber     string       `json:"phone_number" gorm:"size:255;not null"`
	PhoneVerifiedAt time.Time    `json:"phone_verified_at"`
	Photo           string       `json:"photo" sql:"size:999999"`
	Type            ACCOUNT_TYPE `json:"type" gorm:"type:varchar(100);not null"`
	CountryID       uint
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}


type CardRequest struct {
	ID        string    `gorm:"primary_key"`
	UserID    string    `gorm:"primary_key"`
	Color     string    `gorm:"size:255;not null"`
	Currency  string    `gorm:"size:255;not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type Card struct {
	gorm.Model
	// owner
	UserID string

	// card_id
	CardId string `gorm:"type:varchar(50)" json:"card_id"`
	// details
	Details string `gorm:"type:text" json:"details"`
}