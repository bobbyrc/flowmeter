package sample

import (
	"github.com/showwin/speedtest-go/speedtest"
	log "github.com/sirupsen/logrus"
)

func GetConnection() *speedtest.User {
	connection, err := speedtest.FetchUserInfo()
	if err != nil {
		log.Error(err)
	}
	return connection
}

func FindServers(connection *speedtest.User) speedtest.Servers {
	serverList, err := speedtest.FetchServers(connection)
	if err != nil {
		log.Error(err)
	}
	return serverList
}

func GetBestServer(serverList speedtest.Servers) speedtest.Server {
	bestServerList, err := serverList.FindServer([]int{})
	if err != nil {
		log.Error(err)
	}
	bestServer := bestServerList[0]

	return *bestServer
}
