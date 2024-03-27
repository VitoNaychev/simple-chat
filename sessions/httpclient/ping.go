package httpclient

import (
	"fmt"
	"log"
	"net/http"
)

func Ping(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error pinging %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error: Status code %d received from %s\n", resp.StatusCode, url)
		return
	}

	log.Printf("Ping successful at %s\n", url)
}
