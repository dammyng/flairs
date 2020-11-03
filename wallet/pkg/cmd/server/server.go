package cmd

import (
	"wallet/libs/setup"
	"wallet/pkg/protocol/grpc"
	"wallet/pkg/protocol/rest"
	v1 "wallet/pkg/service/v1"
	v1internals "wallet/internals/v1"

	"context"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"

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


	v1API := v1.NewflairsWalletServer(sqlLayer)

	// run HTTP gateway
	go func() {
		_ = rest.RunServer(ctx, cfg.GRPCPort, cfg.HTTPPort)
	}()

	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
