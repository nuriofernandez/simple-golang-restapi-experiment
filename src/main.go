package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	. "./controllers"
)

func main() {
	apiRouter := mux.NewRouter().StrictSlash(true)
	apiRouter.HandleFunc("/", HomePage)
	apiRouter.HandleFunc("/servers", AllServers)
	apiRouter.HandleFunc("/server", NewServer).Methods("POST")
	apiRouter.HandleFunc("/server/{id}", DeleteServer).Methods("DELETE")
	apiRouter.HandleFunc("/server/{id}", SingleServer)
	log.Fatal(http.ListenAndServe(":8080", apiRouter))
}
