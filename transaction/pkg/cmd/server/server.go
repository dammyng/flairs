package cmd

import (
	"log"
	v1internals "transaction/internals/v1"
	"transaction/libs/setup"
	"transaction/pkg/protocol/grpc"
	"transaction/pkg/protocol/rest"
	v1 "transaction/pkg/service/v1"

	"context"
	"fmt"
	"os"
	e_amqp "shared/events/amqp"

	_ "github.com/go-sql-driver/mysql"
	"github.com/streadway/amqp"

	"github.com/jinzhu/gorm"
)

// Config is configuration for Server
type Config struct {
	// gRPC server start parameters section
	// gRPC is TCP port to listen by gRPC server
	GRPCPort string
	// HTTP/REST gateway start parameters section
	// HTTPPort is TCP port to listen by HTTP/REST gateway
	HTTPPort string // DB Datastore parameters section
	// DatastoreDBHost is host of database
	DatastoreDBHost string
	// DatastoreDBUser is username to connect to database
	DatastoreDBUser string
	// DatastoreDBPassword password to connect to database
	DatastoreDBPassword string
	// DatastoreDBSchema is schema of database
	DatastoreDBSchema string
}

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	// get configuration
	var cfg Config
	cfg.GRPCPort = os.Getenv("GRPCPort")
	cfg.HTTPPort = os.Getenv("HTTPPort")
	cfg.DatastoreDBHost = os.Getenv("DatastoreDBHost")
	cfg.DatastoreDBUser = os.Getenv("DatastoreDBUser")
	cfg.DatastoreDBPassword = os.Getenv("DatastoreDBPassword")
	cfg.DatastoreDBSchema = os.Getenv("DatastoreDBSchema")

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	if len(cfg.HTTPPort) == 0 {
		return fmt.Errorf("invalid TCP port for HTTP gateway: '%s'", cfg.HTTPPort)
	}

	// add MySQL driver specific parameter to parse date/time
	// Drop it for another database
	param := "charset=utf8&parseTime=True&loc=Local"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		cfg.DatastoreDBUser,
		cfg.DatastoreDBPassword,
		cfg.DatastoreDBHost,
		cfg.DatastoreDBSchema,
		param)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	db.Exec(setup.SQLMode)
	sqlLayer := v1internals.NewMysqlLayer(db)

	defer db.Close()

	conn, err := amqp.Dial(os.Getenv("AMQP_URL"))
	if err != nil {
		log.Fatal("could not establish amqp connection :" + err.Error())
	}

	eventEmitter, err := e_amqp.NewAMQPEventEmitter(conn, "auth")
	if err != nil {
		log.Fatal("could not establish amqp connection :" + err.Error())
	}

	v1API := v1.NewflairsTransactionServer(sqlLayer, eventEmitter)

	// run HTTP gateway
	go func() {
		_ = rest.RunServer(ctx, cfg.GRPCPort, cfg.HTTPPort)
	}()

	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
