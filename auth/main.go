package main

import (
	"log"
	"net/http"
	"auth/rest"
	"auth/grpc/authserver"
)

func main() {
	
	go authserver.Start()	

	r := rest.ServerRoute()
	log.Fatal(http.ListenAndServe("0.0.0.0:9000", r))

}
