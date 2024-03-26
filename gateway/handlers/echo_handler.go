package handlers

import (
	"encoding/json"
	"net/http"
)

var echoMessage = EchoMessage{"Hello from Gateway!"}

func EchoHandler(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(echoMessage)
}
