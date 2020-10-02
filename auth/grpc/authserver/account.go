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

// AllUsers - get the list of all registered user
func (a *AuthServer) AllUsers(ctx context.Context, in *appuser.Empty) (*appuser.Users, error) {
	var u = []*appuser.User{}
	//data, err := a.DbHandler.AllUsers()
	//if err != nil {
	//	log.Print("Error querying Db")
	//}
	//for _, t := range data {
	//	u = append(u, &t)
	//}
	//
	//log.Print("Receive message body from client")
	return &appuser.Users{
		Results: u,
	}, nil
}

// GetUser - get a single user
func (a *AuthServer) GetUser(ctx context.Context, in *appuser.UserArg) (*appuser.User, error) {
	result, err := a.DbHandler.GetUser(in.UserPayload)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// FindUsers filter users by an obj
func (a *AuthServer) FindUsers(ctx context.Context, in *appuser.UserArg) (*appuser.Users, error) {
	log.Print("Receive message body from client")
	return &appuser.Users{
		Results: []*appuser.User{},
	}, nil
}

// AddUser add a new user
func (a *AuthServer) AddUser(ctx context.Context, in *appuser.UserArg) (*appuser.User, error) {
	log.Println("Receive add user message from client")
	err := a.DbHandler.AddUser(*in.UserPayload)
	if err != nil {
		return nil, err
	}
	return in.UserPayload, nil
}

// UpdateUser update a user record
func (a *AuthServer) UpdateUser(ctx context.Context, in *appuser.UpdateArg) (*appuser.Empty, error) {
	log.Print("Receive message body from client")
	return nil, a.DbHandler.UpdateUser(in)
}

// DeleteUser assuming it is possible
func (a *AuthServer) DeleteUser(ctx context.Context, in *appuser.UserArg) (*appuser.Empty, error) {
	log.Print("Receive message body from client")
	return &appuser.Empty{}, nil
}
