package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"project/controllers"
)

func HandleRequests() {

	router := httprouter.New()
	router.GET("/ping", controllers.PingHandler)
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	HandleRequests()
}
