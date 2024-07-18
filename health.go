package main

import (
	"encoding/json"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	bytes, _ := json.Marshal("API RUNNING")

	w.Write(bytes)
}

func ErrorCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	bytes, _ := json.Marshal("API ERROR")

	w.Write(bytes)
}
