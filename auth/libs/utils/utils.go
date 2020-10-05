package utils

import (
	"auth/libs/sqlmodel"
	"log"
	"shared/models/appuser"
	"time"
)

//GRPCModelToSQL as the name implies
func GRPCModelToSQL(in *appuser.User) *sqlmodel.User {
	out := sqlmodel.User{
		ID:                      in.ID,
		FirstName:               in.FirstName,
		LastName:                in.LastName,
		Address:                 in.Address,
		Street:                  in.Street,
		City:                    in.City,
		PostalCode:              in.PostalCode,
		State:                   in.State,
		Country:                 in.Country,
		Referrer:                in.Referrer,
		RefCode:                 in.RefCode,
		How_did_u_hear_about_us: in.HowDidUHearAboutUs,
		Username:                in.UserName,
		LastCardRequest:         in.LastCardRequested,
		Passport:                in.Passport,
		IDCard:                  in.IDCard,
		BVN:                     in.BVN,
		Wallet:                  0,
		Customer:                0,
		Gender:                  in.Gender,
		DOB:                     usRFC3339Time(in.DOB),
		Email:                   in.Email,
		EmailVerifiedAt:         usRFC3339Time(in.EmailVerifiedAt),
		Password:                in.Password,
		Pin:                     in.Pin,
		IsProfileComplete:       false,
		PhoneNumber:             in.PhoneNumber,
		PhoneVerifiedAt:         usRFC3339Time(in.PhoneVerifiedAt),
		Photo:                   in.Photo,
		Type:                    sqlmodel.ACCOUNT_TYPE(in.ACCOUNT_TYPE),
		CountryID:               0,
		CreatedAt:               usRFC3339Time(in.CreatedAt),
		UpdatedAt:               usRFC3339Time(in.UpdatedAt),
	}
	return &out
}

func usRFC3339Time(in string) time.Time {
	if len(in) < 2 {
		in = "2006-01-02T15:04:05Z"
	}
	t, err := time.Parse(time.RFC3339, in)
	if err != nil {
		log.Panic("invalid date string", err)
	}
	return t
}
