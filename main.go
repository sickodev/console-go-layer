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

	print("Server started at port 8080")
	err := http.ListenAndServe(":8080", router)

	if err != nil {
		panic("Something went wrong" + err.Error())
	}
}
