// To run:
// go build main.go
// ./main
// Then on a local machine, go to localhost:8080 in a browser

package main

import (
	"log"
	"net/http"

	"mongo-api/miner"

	"github.com/gorilla/mux"
)

func main() {
	//create a new router
	router := mux.NewRouter()

	log.Printf("Webpage started")

	//specify endpoints
	router.HandleFunc("/", miner.Home).Methods("GET")

	router.HandleFunc("/api-status", miner.ApiStatus).Methods("GET")

	router.HandleFunc("/mongo-mine/{ip_addr}", miner.MongoMine).Methods("GET")

	// router.HandleFunc("/addsearch/{regex}", miner.AddRegEx).Methods("GET")
	// router.HandleFunc("/clear", miner.Clear).Methods("GET")
	// router.HandleFunc("/reset", miner.ResetArray).Methods("GET")

	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)

}
