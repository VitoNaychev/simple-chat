package main

import (
	"time"

	"github.com/VitoNaychev/simple-chat/sessions/httpclient"
)

func main() {
	url := "http://gateway:8080"

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		httpclient.Ping(url)
		<-ticker.C
	}
}
