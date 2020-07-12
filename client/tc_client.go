package client

import (
	tcc "github.com/ljb021027/tcc/proto/go"
	"github.com/ljb021027/tcc/util"
	log "github.com/sirupsen/logrus"
)

func NewTcClient() tcc.TcClient {
	util.InitClientConfig()
	util.InitGrpcCli()
	cli, err := util.GloalGrpcCliCache.GetCli(util.CConfig.TcAddr)
	if err != nil {
		log.Fatal(err)
	}
	tcClient := tcc.NewTcClient(cli)
	return tcClient
}

func StartTm(tcClient tcc.TcClient, services ...TccService) Tm {
	rms := make([]*RmProxyClient, len(services))
	for index, tccService := range services {
		client, err := NewRmProxyClient(tccService.GetRmResource())
		if err != nil {

		}
		rms[index] = client
	}
	return NewTmService(tcClient, rms)
}
