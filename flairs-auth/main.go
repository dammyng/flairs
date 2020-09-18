package main

import (
	"flairs/flairs-auth/rest"
	"log"
	"net/http"
)

func main() {
	r := rest.ServerRoute()
	log.Println(http.ListenAndServe("0.0.0.0:9000", r))

}
