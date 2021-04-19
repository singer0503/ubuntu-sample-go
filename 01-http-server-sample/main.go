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

func downloadFileSample(w http.ResponseWriter, r *http.Request) {
	file := "FileZilla_3.53.1_win64-setup.exe"
	fmt.Println("in call : " + file)
	// 設定此 Header 告訴瀏覽器下載檔案。 如果沒設定則會在新的 tab 開啟檔案。
	w.Header().Set("Content-Disposition", "attachment; filename="+file)

	http.ServeFile(w, r, file)
}

func main() {
	port := ":9000"
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/download", downloadFileSample)
	fmt.Println("ok")
	fmt.Println("port = " + port)
	http.ListenAndServe(port, nil)

}
