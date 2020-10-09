package persistence

import (
	"fmt"
	"log"
	"reflect"
	"shared/models/appuser"

	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
)

func IsZero(v interface{}) (bool, error) {
	t := reflect.TypeOf(v)
	if !t.Comparable() {
		return false, fmt.Errorf("type is not comparable: %v", t)
	}
	return v == reflect.Zero(t).Interface(), nil
}

func CleanJson(u appuser.User) appuser.User {
	toMap := structs.Map(u)
	delete(toMap, "Password")
	delete(toMap, "Pin")
	var newUser appuser.User
	err := mapstructure.Decode(toMap, &newUser)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return newUser
}


//IsProfileComplete check for complete ptofile
func IsProfileComplete(u appuser.User) bool {
	arr := [9]interface{}{u.FirstName, u.LastName, u.HowDidYouHearAboutUs, u.BVN, u.DOB, u.PhoneNumber, u.Gender, u.ACCOUNT_TYPE.Type(), u.UserName}
	arr_name := [9]interface{}{"FirstName", "LastName", "How_did_you_hear_about_us", "BVN", "DOB", "PhoneNumber", "Gender", "Type", "Username"}

	for i, v := range arr {
		fmt.Printf("comparing %s \n", arr_name[i])
		isZ, err := IsZero(v)
		if err != nil {
			log.Fatalf(err.Error())
		}
		if isZ == true {
			fmt.Printf("%s:%v  is true \n", arr_name[i], v)
			return false
		}
	}
	zeroCard, err := IsZero(u.IDCard)
	zeroPassport, err := IsZero(u.Passport)

	if err != nil {
		log.Fatalf(err.Error())
	}

	bothZero := zeroCard && zeroPassport

	if bothZero {
		fmt.Printf("card %v or passport %v is true", u.IDCard, u.Passport)
		return false
	}

	//emptyByteVar := make([]byte, 128)

	if string(u.Pin) == "" {
		fmt.Printf("pin:%v  is true\n", u.Pin)
		return false
	}
	fmt.Printf("pin:%v", string(u.Pin))

	return true
}
