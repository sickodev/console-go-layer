package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	bytes, _ := json.Marshal("Created an Entry")
	w.Write(bytes)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	bytes, _ := json.Marshal("Deleted an Entry")
	w.Write(bytes)
}

func GetAllEntries(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	bytes, _ := json.Marshal("Got All Entries")
	w.Write(bytes)
}

func GetEntryById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	w.WriteHeader(http.StatusOK)
	msg := "Get entry by id: " + id
	bytes, _ := json.Marshal(msg)
	w.Write(bytes)
}
