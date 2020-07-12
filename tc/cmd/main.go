package main

import (
	"github.com/ljb021027/tcc/tc"
	"github.com/ljb021027/tcc/util"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := util.InitServiceConfig()
	if err != nil {
		log.Fatal()
	}
	util.InitDb()
	//启动tc service
	tcService := tc.NewTcService()
	tcService.StartListen(util.SConfig.TcPort)
}
