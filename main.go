package main

import (
	"github.com/dimayaschhu/mysound/src/rourers"
	"log"
	"net/http"
)

func main() {
	router := rourers.NewRouter()
	router.SetRoutes()
	if err := http.ListenAndServe(":80", router.Engine()); err != nil {
		log.Fatal(err)
	}
}
