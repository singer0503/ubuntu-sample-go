package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetDownloadFileSample(c *gin.Context) {
	fileName := "FileZilla_3.53.1_win64-setup.exe"
	// 設定此 Header 告訴瀏覽器下載檔案。 如果沒設定則會在新的 tab 開啟檔案。
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName)) //fmt.Sprintf("attachment; filename=%s", filename) Downloaded file renamed
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(fileName)
}

func main() {
	port := ":8090"
	server := gin.Default()
	server.GET("/download", GetDownloadFileSample)
	logrus.Fatal(server.Run(port))
}
