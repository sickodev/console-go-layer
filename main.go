package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Console AI Endpoints
	router.HandleFunc("/console", TalkToSinestro).Methods(http.MethodPost)

	// Database Endpoints
	router.HandleFunc("/database", GetAllEntries).Methods(http.MethodGet)
	router.HandleFunc("/database/{id}", GetEntryById).Methods(http.MethodGet)
	router.HandleFunc("/database", Create).Methods(http.MethodPost)
	router.HandleFunc("/database", Delete).Methods(http.MethodDelete)

	//Health Functions
	router.HandleFunc("/healthz", HealthCheck).Methods(http.MethodGet)
	router.HandleFunc("/error", ErrorCheck).Methods(http.MethodGet)

	print("Server started at port 8080\n")
	err := http.ListenAndServe(":8080", router)

	if err != nil {
		panic("Something went wrong" + err.Error())
	}
}
