package main

import (
    "fmt"
    "net/http"
)

// 測試主機是否有回應
func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello\n")
}

// 回傳主機看到的 headers
func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {
	port := ":80"
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)

	fmt.Println("ok")
	fmt.Println("port = "+port)
    http.ListenAndServe(port, nil)
	
}