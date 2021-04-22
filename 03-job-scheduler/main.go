package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
var loginEndpoint string = "http://35.201.243.111:8090/hello"
var count int = 0

func taskHttpRequest() {
	conntransport := &http.Transport{
		DisableKeepAlives:  true, // 這樣才會 send FIN 封包
		DisableCompression: true,
	}
	count++
	startTime := time.Now()
	fmt.Print(strconv.Itoa(count) + "," + startTime.Format("2006-01-02 15:04:05.000000"))
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
		deferTime := startTime.Sub(startTime)
		fmt.Print("," + endTime.Format("2006-01-02 15:04:05.000000"))
		fmt.Println("," + strconv.FormatInt(deferTime.Milliseconds(), 10))
	}()
}

func main() {
	fmt.Println("No,StartTime,EndTime,UseTime")
	taskHttpRequest()

	gocron.Every(1).Minute().Do(taskHttpRequest) // 改成 1 分鐘一次
	<-gocron.Start()

	//fmt.Println("begin")

	// taskHttpRequest()
	// time.Sleep(time.Duration(1) * time.Second)
	// taskHttpRequest()
	// time.Sleep(time.Duration(1) * time.Second)
	// taskHttpRequest()

	//fmt.Println("end")
	// Do jobs without params
	// gocron.Every(1).Day().At("00:02").Do(task)
	// gocron.Every(1).Day().At("00:32").Do(task)
	// gocron.Every(1).Day().At("01:02").Do(task)
	// gocron.Every(1).Day().At("01:32").Do(task)
	// gocron.Every(1).Day().At("02:02").Do(task)
	// gocron.Every(1).Day().At("02:32").Do(task)
	// gocron.Every(1).Day().At("03:02").Do(task)
	// gocron.Every(1).Day().At("03:32").Do(task)
	// gocron.Every(1).Day().At("04:02").Do(task)
	// gocron.Every(1).Day().At("04:32").Do(task)
	// gocron.Every(1).Day().At("05:02").Do(task)
	// gocron.Every(1).Day().At("05:32").Do(task)
	// gocron.Every(1).Day().At("06:02").Do(task)
	// gocron.Every(1).Day().At("06:32").Do(task)
	// gocron.Every(1).Day().At("07:02").Do(task)
	// gocron.Every(1).Day().At("07:32").Do(task)
	// gocron.Every(1).Day().At("08:02").Do(task)
	// gocron.Every(1).Day().At("08:32").Do(task)
	// gocron.Every(1).Day().At("09:02").Do(task)
	// gocron.Every(1).Day().At("09:32").Do(task)
	// gocron.Every(1).Day().At("10:02").Do(task)
	// gocron.Every(1).Day().At("10:32").Do(task)
	// gocron.Every(1).Day().At("11:02").Do(task)
	// gocron.Every(1).Day().At("11:32").Do(task)

	// gocron.Every(1).Day().At("12:02").Do(task)
	// gocron.Every(1).Day().At("12:32").Do(task)
	// gocron.Every(1).Day().At("13:02").Do(task)
	// gocron.Every(1).Day().At("13:32").Do(task)
	// gocron.Every(1).Day().At("14:02").Do(task)
	// gocron.Every(1).Day().At("14:32").Do(task)
	// gocron.Every(1).Day().At("15:02").Do(task)
	// gocron.Every(1).Day().At("15:32").Do(task)
	// gocron.Every(1).Day().At("16:02").Do(task)
	// gocron.Every(1).Day().At("16:32").Do(task)
	// gocron.Every(1).Day().At("17:02").Do(task)
	// gocron.Every(1).Day().At("17:32").Do(task)
	// gocron.Every(1).Day().At("18:02").Do(task)
	// gocron.Every(1).Day().At("18:32").Do(task)
	// gocron.Every(1).Day().At("19:02").Do(task)
	// gocron.Every(1).Day().At("19:32").Do(task)
	// gocron.Every(1).Day().At("20:02").Do(task)
	// gocron.Every(1).Day().At("20:32").Do(task)
	// gocron.Every(1).Day().At("21:02").Do(task)
	// gocron.Every(1).Day().At("21:32").Do(task)
	// gocron.Every(1).Day().At("22:02").Do(task)
	// gocron.Every(1).Day().At("22:32").Do(task)
	// gocron.Every(1).Day().At("23:02").Do(task)
	// gocron.Every(1).Day().At("23:32").Do(task)

	//gocron.Every(2).Second().Do(taskHttpRequest) // 原本 2 秒一次

	// gocron.Every(1).Minute().Do(task)
	// gocron.Every(2).Minutes().Do(task)
	//gocron.Every(1).Hour().Do(task)
	//gocron.Every(2).Hours().Do(task)
	// gocron.Every(1).Day().Do(task)
	// gocron.Every(2).Days().Do(task)
	// gocron.Every(1).Week().Do(task)
	// gocron.Every(2).Weeks().Do(task)

	// // Do jobs with params
	// gocron.Every(1).Second().Do(taskWithParams, 1, "hello")

	// // Do jobs on specific weekday
	// gocron.Every(1).Monday().Do(task)
	// gocron.Every(1).Thursday().Do(task)

	// // Do a job at a specific time - 'hour:min:sec' - seconds optional
	// gocron.Every(1).Day().At("10:30").Do(task)
	// gocron.Every(1).Monday().At("18:30").Do(task)
	// gocron.Every(1).Tuesday().At("18:30:59").Do(task)

	// // Begin job immediately upon start
	// gocron.Every(1).Hour().From(gocron.NextTick()).Do(task)

	// // Begin job at a specific date/time
	// t := time.Date(2019, time.November, 10, 15, 0, 0, 0, time.Local)
	// gocron.Every(1).Hour().From(&t).Do(task)

	// // NextRun gets the next running time
	// _, time := gocron.NextRun()
	// fmt.Println(time)

	// // Remove a specific job
	// gocron.Remove(task)

	// // Clear all scheduled jobs
	// gocron.Clear()

	// // Start all the pending jobs
	// <-gocron.Start()

	// // also, you can create a new scheduler
	// // to run two schedulers concurrently
	// s := gocron.NewScheduler()
	// s.Every(3).Seconds().Do(task)
	// <-s.Start()
}
