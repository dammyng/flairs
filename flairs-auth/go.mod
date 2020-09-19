module flairs/flairs-auth

go 1.14

replace shared => /Users/kdnotes/src/flairs/shared-libraries

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gorilla/mux v1.8.0
	github.com/jinzhu/gorm v1.9.16
	shared v0.0.0-00010101000000-000000000000
)
