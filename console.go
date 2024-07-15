package main

import "net/http"

func TalkToSinestro(w http.ResponseWriter, r *http.Request) {
	print("Talk to Sinestro!!")
}
