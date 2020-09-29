package authserver

import (
	"log"
	"net"
	"shared/models/appuser"
	"google.golang.org/grpc"

)

// Start GRPC Server
func Start(auth AuthServer, port string)  {

	listenter, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen :%v", err)
	}

	grpcServer := grpc.NewServer()
	appuser.RegisterUserServiceServer(grpcServer, &auth)
defer grpcServer.Stop()
	if err := grpcServer.Serve(listenter); err != nil{
		log.Fatalf("failed to serve: %s", err)
	}
	log.Println("grpc started")
}