package main

import (
	"fmt"
	"net/http"
	"path/filepath"

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

func GetDownloadFileSample(c *gin.Context) {
	fileName := "FileZilla_3.53.1_win64-setup.exe"
	// 設定此 Header 告訴瀏覽器下載檔案。 如果沒設定則會在新的 tab 開啟檔案。
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName)) //fmt.Sprintf("attachment; filename=%s", filename) Downloaded file renamed
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(fileName)
}
func PostUploadFile(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("檔案上傳成功！ File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, email))
}

func main() {
	port := ":8443"
	router := gin.Default()

	router.GET("/hello", GetHello)
	router.GET("/headers", GetHeaders)
	router.GET("/download", GetDownloadFileSample)
	router.Static("/sanfran", "./public")
	router.POST("/upload", PostUploadFile)

	//logrus.Fatal(router.RunTLS(port, "server.crt", "server.key"))
	logrus.Fatal(router.RunTLS(port, "maxhuang_me.crt", "myserver.key"))
}
