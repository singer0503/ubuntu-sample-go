package main

import (
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
var loginEndpoint string = "http://35.201.243.111:8090/download"
var count int = 0

var w *csv.Writer

func taskDownloadRequest() {
	conntransport := &http.Transport{
		DisableKeepAlives:  true, // 這樣才會 send FIN 封包
		DisableCompression: true,
	}
	count++
	startTime := time.Now()
	url := loginEndpoint
	req, err := http.NewRequest("GET", url, nil)
	//client := &http.Client{}
	client := &http.Client{Transport: conntransport}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	ioutil.ReadAll(resp.Body) // 取得回應 body 資料

	defer func() {
		resp.Body.Close()
		endTime := time.Now()
		deferTime := endTime.Sub(startTime)
		fmt.Println(strconv.Itoa(count) + "," + startTime.Format("2006-01-02 15:04:05.000") + "," + endTime.Format("2006-01-02 15:04:05.000") + "," + strconv.FormatInt(deferTime.Milliseconds(), 10) + "ms")
		w.Write([]string{strconv.Itoa(count), startTime.Format("2006-01-02 15:04:05.000"), endTime.Format("2006-01-02 15:04:05.000"), strconv.FormatInt(deferTime.Milliseconds(), 10) + "ms"})
		w.Flush()
	}()
	//resp.Body.Close()
}

func main() {
	// 不存在則建立;存在則清空;讀寫模式;
	file, err := os.Create("call_download_list.csv")
	if err != nil {
		fmt.Println("open file is failed, err: ", err)
	}
	// 延遲關閉
	defer file.Close()

	// 寫入UTF-8 BOM，防止中文亂碼
	file.WriteString("\xEF\xBB\xBF")

	w = csv.NewWriter(file)
	// 寫入資料
	w.Write([]string{"發送編號", "起始時間", "結束時間", "消耗時間"})
	w.Flush()

	fmt.Println("No,StartTime,EndTime,UseTime")

	taskDownloadRequest()

	gocron.Every(1).Minute().Do(taskDownloadRequest)
	<-gocron.Start()

}
