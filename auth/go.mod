module auth

replace shared => ../shared

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fatih/structs v1.1.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.2
	github.com/gomodule/redigo v1.8.2
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.15.0
	github.com/jinzhu/gorm v1.9.16
	github.com/joho/godotenv v1.3.0
	github.com/mitchellh/mapstructure v1.3.3
	github.com/satori/go.uuid v1.2.0
	github.com/streadway/amqp v1.0.0
	github.com/urfave/negroni v1.0.0
	golang.org/x/crypto v0.0.0-20201002170205-7f63de1d35b0
	golang.org/x/net v0.0.0-20201006153459-a7d1128ccaa0
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
	google.golang.org/grpc v1.32.0
	google.golang.org/protobuf v1.25.0
	shared v0.0.0-00010101000000-000000000000
)
