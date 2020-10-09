package authclient

import (
	"log"
	"os"
	"shared/models/appuser"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func Connect() (interface{}, error) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9011", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	authClient := appuser.NewUserServiceClient(conn)

	return authClient.AllUsers(context.Background(), &appuser.Empty{})
	//if err != nil {
	//	log.Fatalf("Error when calling SayHello: %s", err)
	//}
	//log.Printf("Response from server:   %v", response.Results)

}

// AddNewUser grpc client to add new user
func AddNewUser(in *appuser.UserArg) (*appuser.User, error) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(os.Getenv("GRPCPort"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	authClient := appuser.NewUserServiceClient(conn)

	result, err := authClient.AddUser(context.Background(), in)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetAUser(in *appuser.UserArg) (*appuser.User, error) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(os.Getenv("GRPCPort"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	authClient := appuser.NewUserServiceClient(conn)

	result, err := authClient.GetUser(context.Background(), in)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateUser(in *appuser.UpdateArg) (*appuser.Empty, error) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(os.Getenv("GRPCPort"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	authClient := appuser.NewUserServiceClient(conn)

	_, err = authClient.UpdateUser(context.Background(), in)

	if err != nil {
		return new(appuser.Empty), err
	}
	return new(appuser.Empty), nil
}

func UpdateUserMap(in *appuser.UpdateArg) (*appuser.Empty, error) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(os.Getenv("GRPCPort"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	authClient := appuser.NewUserServiceClient(conn)

	_, err = authClient.UpdateUser(context.Background(), in)

	if err != nil {
		return new(appuser.Empty), err
	}
	return new(appuser.Empty), nil
}