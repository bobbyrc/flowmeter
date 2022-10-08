package sample

import (
	"fmt"
	"net/http"

	"github.com/bobbyrc/flowmeter/src/cmd/sample"
	"github.com/gin-gonic/gin"
)

func TakeSampleHandler(c *gin.Context) {
	connection := sample.GetConnection()
	serverList := sample.FindServers(connection)
	server := sample.GetBestServer(serverList)

	server.PingTest()
	server.DownloadTest(false)
	server.UploadTest(false)

	fmt.Printf("Latency: %s, Download: %f, Upload: %f\n", server.Latency, server.DLSpeed, server.ULSpeed)

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
