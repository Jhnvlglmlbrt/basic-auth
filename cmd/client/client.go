package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	client := http.Client{Timeout: 5 * time.Second}

	req, err := http.NewRequest(http.MethodGet, "https://localhost:4000/protected", http.NoBody)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("alice", "p8fnxeqj5a7zbrqp")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Status: %d\n", res.StatusCode)
	fmt.Printf("Body: %s\n", string(resBody))
}
