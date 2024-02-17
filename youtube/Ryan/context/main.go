package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {

	startNow := time.Now()

	timeoutContext, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()

	// create HTTP request
	url := "http://placehold.it/2000x2000"
	// req, err := http.NewRequest(http.MethodGet, url, nil)
	req, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("http.NewRequestWithContext() fail: %v\n", err)
		return
	}

	// perform HTTP request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("http.DefaultClient.Do() fail: %v\n", err)
		return
	}
	defer res.Body.Close()

	// get data from HTTP response
	imageData, _ := io.ReadAll(res.Body)

	fmt.Printf("downloaded image of size %d\n", len(imageData))

	fmt.Println("This operation took:", time.Since(startNow))

}
