package httpclient

import (
	"fmt"
	"net/http"
	"time"
)

func Ping(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error pinging %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: Status code %d received from %s\n", resp.StatusCode, url)
		return
	}

	fmt.Printf("Ping successful at %s\n", time.Now().Format(time.RFC3339))
}
