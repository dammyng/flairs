package v1

import (
	v1internals "auth/internals/v1"
	v1 "auth/pkg/api/v1"
	"fmt"
	"log"
	"reflect"

	"github.com/dgrijalva/jwt-go"
	"github.com/gomodule/redigo/redis"
	amqp "shared/events/amqp"

)


type flairsServiceServer struct {
	Db v1internals.DatabaseHandler
	RedisConn    redis.Conn
	EventEmitter amqp.EventEmitter
}

// Claims jwt custom authentication claims 
type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

// NewFlairsServiceServer creates ToDo service
func NewFlairsServiceServer(db v1internals.DatabaseHandler,redisConn redis.Conn, eventEmitter amqp.EventEmitter) v1.FlairsServiceServer {
	return &flairsServiceServer{
		Db: db, 
		RedisConn:    redisConn, 
		EventEmitter: eventEmitter,
	}
}

// DecodeJwt - decodes JWT token from request
func DecodeJwt(token string, claims *Claims) error {
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_FLAIRS_KEY")), nil
	})
	return err
}


// IsProfileCompled -> check if user profile contains all required property
func IsProfileCompled(u *v1.User) bool {
	arr := [9]interface{}{u.FirstName, u.LastName, u.HowDidYouHearAboutUs, u.BVN, u.DOB, u.PhoneNumber, u.Gender, u.ACCOUNT_TYPE, u.UserName}
	arrName := [9]interface{}{"FirstName", "LastName", "HowDidYouHearAboutUs", "BVN", "DOB", "PhoneNumber", "Gender", "ACCOUNT_TYPE", "UserName"}

	for i, v := range arr {
		isZ, err := IsZero(v)
		if err != nil {
			log.Fatalf(err.Error())
		}
		if isZ == true {
			fmt.Printf("%s:%v  is true \n", arrName[i], v)
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
		return false
	}

	if string(u.Pin) == "" {
		return false
	}
	return true
}

func IsZero(v interface{}) (bool, error) {
	t := reflect.TypeOf(v)
	if !t.Comparable() {
		return false, fmt.Errorf("type is not comparable: %v", t)
	}
	return v == reflect.Zero(t).Interface(), nil
}

