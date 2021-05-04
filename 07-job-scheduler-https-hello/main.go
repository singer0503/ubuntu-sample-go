package main

import (
	"crypto/tls"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jasonlvhit/gocron"
)

// refer https://github.com/jasonlvhit/gocron

func task() {
	fmt.Println("I am running task.")
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}

// 目標 URL
var loginEndpoint string = "https://maxhuang.me:8443/hello"
var count int = 0

var w *csv.Writer

func taskHttpRequest() {
	// disable security checks globally for all requests of the default client:
	//http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	// refer http://www.honlyc.com/post/golang-x509-certificate-unknown-authority/
	conntransport := &http.Transport{
		DisableKeepAlives:  true, // 這樣才會 send FIN 封包
		DisableCompression: true,
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}
	count++
	startTime := time.Now()
	url := loginEndpoint
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Transport: conntransport}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	ioutil.ReadAll(resp.Body) // 取得回應 body 資料
	//loginBodyString := string(body)
	//fmt.Println(count)
	//fmt.Println(loginBodyString)
	conntransport.CloseIdleConnections()
	defer func() {
		resp.Body.Close()
		endTime := time.Now()
		deferTime := endTime.Sub(startTime)
		fmt.Println(strconv.Itoa(count) + "," + startTime.Format("2006-01-02 15:04:05.000") + "," + endTime.Format("2006-01-02 15:04:05.000") + "," + strconv.FormatInt(deferTime.Milliseconds(), 10) + "ms")
		w.Write([]string{strconv.Itoa(count), startTime.Format("2006-01-02 15:04:05.000"), endTime.Format("2006-01-02 15:04:05.000"), strconv.FormatInt(deferTime.Milliseconds(), 10) + "ms"})
		w.Flush()
	}()
}

func main() {

	// 不存在則建立;存在則清空;讀寫模式;
	// file, err := os.Create("call_https_hello_list.csv")
	// if err != nil {
	// 	fmt.Println("open file is failed, err: ", err)
	// }

	// 不存在則建立;存在繼續往下寫;讀寫模式;
	file, err := os.OpenFile("3_call_https_hello_list.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file is failed, err: ", err)
	}

	// 延遲關閉
	defer file.Close()

	// 寫入UTF-8 BOM，防止中文亂碼
	//file.WriteString("\xEF\xBB\xBF")

	w = csv.NewWriter(file)
	// 寫入資料
	w.Write([]string{"發送編號", "起始時間", "結束時間", "消耗時間"})
	w.Flush()

	fmt.Println("No,StartTime,EndTime,UseTime")
	taskHttpRequest()

	gocron.Every(1).Minute().Do(taskHttpRequest) // 改成 1 分鐘一次
	<-gocron.Start()

}
