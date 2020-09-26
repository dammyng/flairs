module auth

replace shared => ../shared-libraries

go 1.14

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gorilla/mux v1.8.0
	github.com/jinzhu/gorm v1.9.16
	golang.org/x/net v0.0.0-20200822124328-c89045814202
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/grpc v1.32.0
	shared v0.0.0-00010101000000-000000000000
)
