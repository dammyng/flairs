package main

import (
	"log"
	"net/http"
	"auth/rest"
	"auth/libs/persistence"
	"auth/libs/config"
	"auth/grpc/authserver"
)

func main() {
	
	go authserver.Start()

	dbHandler := persistence.NewMysqlLayer(config.DBConfig)
	dbHandler.Session.Exec(config.CreateDatabase)
	dbHandler.Session.Exec(config.UseAlphaPlus)
	

	r := rest.ServerRoute(dbHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:9000", r))

}
