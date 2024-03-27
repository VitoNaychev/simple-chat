package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

var echoMessage = EchoMessage{"Hello from Gateway!"}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %v request from %v", r.Method, r.RemoteAddr)

	json.NewEncoder(w).Encode(echoMessage)
}
