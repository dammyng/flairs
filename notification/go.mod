module notification

replace shared => ../shared

go 1.14

require (
	github.com/GeertJohan/go.rice v1.0.0
	github.com/joho/godotenv v1.3.0
	github.com/sendgrid/rest v2.6.1+incompatible // indirect
	github.com/sendgrid/sendgrid-go v3.6.4+incompatible
	github.com/streadway/amqp v1.0.0
	shared v0.0.0-00010101000000-000000000000
)
