package main

import (
	"log"
	"net/http"

	"github.com/VitoNaychev/simple-chat/gateway/handlers"
)

func main() {
	handler := http.HandlerFunc(handlers.EchoHandler)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
