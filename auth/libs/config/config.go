package config

import (
	"flairs/auth/libs/persistence"
)
var (
	DBConfig = persistence.DBConfig{
		Hosts:    "localhost",
		Database: "alpha_plus",
		Username: "",
		Password: "",
		Port:     "3306",
	}
	AmqpMessageBroker string
	SecretKey         string
	AesKey            string
	RedisHost         = "localhost:6379"
)