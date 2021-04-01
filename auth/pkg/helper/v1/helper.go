package v1

import (
	v1 "auth/pkg/api/v1"
	"crypto/md5"
	"encoding/hex"
	"log"
	"math/rand"
	"regexp"
	"time"

	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
)

// DecodeToSQLUser -> Decode to SQL
// Parse protocol buffer generated object into an acceptable SQL object
func DecodeToSQLUser(in interface{}) *v1.User {
	m := structs.Map(in)
	var user v1.User
	mapstructure.Decode(m, &user)
	return &user
}


var seedRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

const intset = "0123456789"

// RandStringWithCharSet -> Random Character set
// Generate Random Character set of specific length
func RandStringWithCharSet(length int, charset string) string {

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seedRand.Intn(len(charset))]
	}
	return string(b)
}

// RandInt -> Random Number
// Generate random number of specific length
func RandInt(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = intset[seedRand.Intn(len(intset))]
	}
	return string(b)
}

// RandString -> Random string
// Generate random string of specific length
func RandString(length int) string {
	return RandStringWithCharSet(length, charset)
}

// GenerateToken -> Generate token
// Generates token passed along as responses
func GenerateToken(word string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(word), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	hasher := md5.New()
	hasher.Write(hash)
	return hex.EncodeToString(hasher.Sum(nil))
}


var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// IsEmailValid checks if the email provided passes the required structure and length.
func IsEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}
