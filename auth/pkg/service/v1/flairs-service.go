package v1

import (
	v1 "auth/pkg/api/v1"
	v1internals "auth/internals/v1"

	"context"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/gomodule/redigo/redis"

)

const (
	apiVersion = "v1"
)

type flairsServiceServer struct {
	Db v1internals.DatabaseHandler
	RedisConn    redis.Conn
}

// NewFlairsServiceServer creates ToDo service
func NewFlairsServiceServer(db v1internals.DatabaseHandler,redisConn redis.Conn) v1.FlairsServiceServer {
	return &flairsServiceServer{Db: db, RedisConn:    redisConn,}
}

// checkAPI checks if the API version requested by client is supported by server
func (f *flairsServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

// connect returns SQL database connection from the pool
func (f *flairsServiceServer) connect(ctx context.Context) (*gorm.DB, error) {
	c, err := f.connect(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}

