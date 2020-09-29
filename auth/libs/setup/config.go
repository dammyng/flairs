package setup

import (
	"auth/libs/persistence"
)
var (
	DBConfig = persistence.DBConfig{
		Hosts:    "localhost",
		Database: "alpha_plus",
		Username: "root",
		Password: "password",
		Port:     "3306",
	}
	AmqpMessageBroker string
	SecretKey         string
	AesKey            string
	RedisHost         = "localhost:6379"
)