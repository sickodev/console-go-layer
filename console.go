package main

import (
	"encoding/json"
	"net/http"
)

func TalkToSinestro(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	bytes, _ := json.Marshal("Talking to Sinestro")
	w.Write(bytes)
}
