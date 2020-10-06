package rest

import "auth/libs/sqlmodel"

type (
	EmailPayload struct {
		Email string `json:"email, omitempty"`
	}

	AuthPayload struct {
		Email    string `json:"email, omitempty"`
		Password string `json:"password, omitempty"`
	}

	UpdateUserDataPayload struct {
		FirstName   string `json:"first_name,omitempty"`
		LastName    string `json:"last_name,omitempty"`
		Username    string `json:"username,omitempty"`
		PhoneNumber string `json:"phone_number,omitempty"`
		BVN         string `json:"bvn,omitempty"`

		PhoneVerifiedAt string `json:"phone_verified_at,omitempty"`
		DOB             string `json:"dob,omitempty"`
		Address         string `json:"address,omitempty"`
		Street          string `json:"street,omitempty"`

		City string `json:"city,omitempty"`

		PostalCode string `json:"postal_code,omitempty"`

		State string `json:"state,omitempty"`

		Country string `json:"country,omitempty"`

		Gender                  string `json:"gender,omitempty"`
		How_did_u_hear_about_us string `json:"how_did_u_hear_about_us,omitempty"`
		Passport                string `json:"passport,omitempty"`
		Photo                   string `json:"photo,omitempty"`

		IDCard string                `json:"id_card,omitempty"`
		Type   sqlmodel.ACCOUNT_TYPE `json:"type" gorm:"type:varchar(100);not null"`

		Pin string `json:"pin,omitempty"`
	}
)
