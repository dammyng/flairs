package main

import (
	"flairs/auth/rest"
	"log"
	"net/http"
	"flairs/auth/libs/persistence"
	"flairs/auth/libs/config"
)

func main() {
	dbHandler := persistence.NewMysqlLayer(config.DBConfig)
	r := rest.ServerRoute(dbHandler)
	log.Println(http.ListenAndServe("0.0.0.0:9000", r))

}
