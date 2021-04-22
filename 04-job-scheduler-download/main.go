package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// refer https://github.com/jasonlvhit/gocron

func task() {
	fmt.Println("I am running task.")
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}

// 目標 URL
var loginEndpoint string = "http://35.201.243.111:8080/download"
var count int = 0

func taskDownloadRequest() {
	// conntransport := &http.Transport{
	// 	DisableKeepAlives:  true, // 這樣才會 send FIN 封包
	// 	DisableCompression: true,
	// }
	count++
	fmt.Println(strconv.Itoa(count) + " = start time " + time.Now().Format("2006-01-02 15:04:05.000000"))
	url := loginEndpoint
	req, err := http.NewRequest("GET", url, nil)
	//req.Header.Set("Content-Type", "application/json")
	//client := &http.Client{Transport: conntransport}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	ioutil.ReadAll(resp.Body) // 取得回應 body 資料
	//loginBodyString := string(body)

	//fmt.Println(loginBodyString)

	//conntransport.CloseIdleConnections() // close connection
	defer func() {
		resp.Body.Close()
		fmt.Println(strconv.Itoa(count) + " = ent time " + time.Now().Format("2006-01-02 15:04:05.000000"))
	}()
	//resp.Body.Close()
}

func main() {
	fmt.Println("begin")

	fmt.Println("test")
	taskDownloadRequest()

	fmt.Println("end")

	//gocron.Every(1).Second().Do(taskDownloadRequest)
	//<-gocron.Start()

}
