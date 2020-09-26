package authserver

import (
	"fmt"
	"log"
	"net"
	"shared/models/appuser"
	"google.golang.org/grpc"
)

// Start GRPC Server
func Start()  {

	listenter, err := net.Listen("tcp", fmt.Sprintf(":%d", 9011))
	if err != nil {
		log.Fatalf("failed to listen :%v", err)
	}

	auth := AuthServer{}

	grpcServer := grpc.NewServer()
	defer grpcServer.Stop()
	appuser.RegisterUserServiceServer(grpcServer, &auth)

	if err := grpcServer.Serve(listenter); err != nil{
		log.Fatalf("failed to serve: %s", err)
	}
	log.Println("grpc started")
}