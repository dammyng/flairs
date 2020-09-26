package authclient

import (
	"log"
	"shared/models/appuser"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
)

func Connect() (interface{}, error) {
	var conn  *grpc.ClientConn
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