module transaction

replace shared => ../shared

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/jinzhu/gorm v1.9.16
	github.com/satori/go.uuid v1.2.0
	github.com/streadway/amqp v1.0.0
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
	google.golang.org/grpc v1.33.2
	google.golang.org/protobuf v1.25.0
	shared v0.0.0-00010101000000-000000000000
)
