module flairs/auth

replace shared => ../shared-libraries

go 1.14

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gorilla/mux v1.8.0
	github.com/jinzhu/gorm v1.9.16
	golang.org/x/tools v0.0.0-20200923182640-463111b69878 // indirect
	shared v0.0.0-00010101000000-000000000000
)
