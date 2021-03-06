package main

import (
	"github.com/aiwantaozi/go_http/controller"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func HandleRequests() {

	router := httprouter.New()
	router.GET("/ping", controller.TestPingHandler)
	log.Fatal(http.ListenAndServe(":8082", router))
}

func main() {
	HandleRequests()
}
