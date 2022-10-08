package sample

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/showwin/speedtest-go/speedtest"
)

func TakeSampleHandler(c *gin.Context) {
	user, _ := speedtest.FetchUserInfo()

	serverList, _ := speedtest.FetchServers(user)
	targets, _ := serverList.FindServer([]int{})

	for _, s := range targets {
		s.PingTest()
		s.DownloadTest(false)
		s.UploadTest(false)

		fmt.Printf("Latency: %s, Download: %f, Upload: %f\n", s.Latency, s.DLSpeed, s.ULSpeed)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
