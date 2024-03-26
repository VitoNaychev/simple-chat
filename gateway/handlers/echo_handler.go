package handlers

import (
	"encoding/json"
	"net/http"
)

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	var echoMessage EchoMessage
	json.NewDecoder(r.Body).Decode(&echoMessage)

	json.NewEncoder(w).Encode(echoMessage)
}
