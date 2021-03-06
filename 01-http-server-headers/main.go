package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 測試是否能正常連接
func GetHello(c *gin.Context) {
	// 在http包使用的時候，註冊了/這個根路徑的模式處理，瀏覽器會自動的請求 favicon.ico，要注意處理，否則會出現兩次請求
	if c.Request.RequestURI == "/favicon.ico" {
		return
	}
	//result := "{'msg':'test ok!'}"
	result := "Hello"
	c.JSON(http.StatusOK, result)
}

// 回傳主機看到的 headers
func GetHeaders(c *gin.Context) {
	for name, headers := range c.Request.Header {
		for _, h := range headers {
			fmt.Fprintf(c.Writer, "%v: %v\n", name, h)
		}
	}
}

func main() {
	port := ":8080"
	server := gin.Default()
	server.GET("/hello", GetHello)
	server.GET("/headers", GetHeaders)
	logrus.Fatal(server.Run(port))
}
