package authserver

import (
	"auth/libs/persistence"
	"golang.org/x/net/context"
	"log"
	"shared/models/appuser"
)

// AuthServer implements app user GRPC service
type AuthServer struct {
	DbHandler persistence.DatabaseHandler
}

func NewAuthServer(dbHandler persistence.DatabaseHandler) AuthServer {
	return AuthServer{
		DbHandler: dbHandler,
	}
}

// Find User
func (a *AuthServer) AllUsers(ctx context.Context, in *appuser.Empty) (*appuser.Users, error) {
	var u = []*appuser.User{}
	data, err := a.DbHandler.AllUsers()
	if err != nil {
		log.Print("Error querying Db")
	}
	for _, t := range data {
		u = append(u, &t)
	}

	log.Print("Receive message body from client")
	return &appuser.Users{
		Results: u,
	}, nil
}

// Find User
func (a *AuthServer) UpdateUser(ctx context.Context, in *appuser.UserArg) (*appuser.User, error) {
	log.Print("Receive message body from client")
	return &appuser.User{}, nil
}

// Find User
func (a *AuthServer) FindUser(ctx context.Context, in *appuser.UserArg) (*appuser.User, error) {
	log.Print("Receive message body from client")
	return &appuser.User{}, nil
}

// Find User
func (a *AuthServer) Welcome(ctx context.Context, in *appuser.UserArg) (*appuser.User, error) {
	log.Print("Receive message body from client")
	return &appuser.User{}, nil
}
