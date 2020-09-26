package authserver
import (
	"log"
	"shared/models/appuser"
	"golang.org/x/net/context"
)
// AuthServer implements app user GRPC service
type AuthServer struct {
	
}
// Find User
func (a *AuthServer) FindUser(ctx context.Context, in *appuser.UserArg) (*appuser.User, error) {
	log.Print("Receive message body from client")
	return &appuser.User{}, nil
}


// Find User
func (a *AuthServer) UpdateUser(ctx context.Context, in *appuser.UserArg) (*appuser.User, error) {
	log.Print("Receive message body from client")
	return &appuser.User{}, nil
}


// Find User
func (a *AuthServer) Welcome(ctx context.Context, in *appuser.UserArg) (*appuser.User, error) {
	log.Print("Receive message body from client")
	return &appuser.User{}, nil
}