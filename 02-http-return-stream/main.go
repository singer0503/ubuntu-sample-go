// Start a mock HTTP server that returns 2GB of data in the response. Make a
// HTTP request to this server and print the amount of data read from the
// response.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

const oneMB = 1024 * 1024
const oneGB = 1024 * oneMB
const responseSize = 1 * oneGB

const serverAddr = "localhost:9999"

func startServer() {
	// Mock HTTP server that always returns 1GB of data / 開啟一個 goroutine 模擬 http Server
	go http.ListenAndServe(serverAddr, http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-length", fmt.Sprintf("%d", responseSize))

		// 1MB buffer that'll be copied multiple times to the response
		buf := []byte(strings.Repeat("x", oneMB))

		for i := 0; i < responseSize/len(buf); i++ {
			if _, err := w.Write(buf); err != nil {
				log.Fatal("Failed to write to response. Error: ", err.Error())
			}
		}
	}))

	// Some grace period for the server to start
	time.Sleep(100 * time.Millisecond)
}

func main() {
	startServer()

	// HTTP client
	req, err := http.NewRequest("GET", "http://"+serverAddr, nil)
	if err != nil {
		log.Fatal("Error creating HTTP request: ", err.Error())
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error making HTTP request: ", err.Error())
	}

	// Read the response header
	fmt.Println("Response: Content-length:", resp.Header.Get("Content-length"))

	bytesRead := 0
	buf := make([]byte, oneMB)

	// Read the response body
	for {
		n, err := resp.Body.Read(buf)
		bytesRead += n

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal("Error reading HTTP response: ", err.Error())
		}
	}

	fmt.Println("Response: Read", bytesRead, "bytes")
}
