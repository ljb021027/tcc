package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ljb021027/tcc/client"
	tcc "github.com/ljb021027/tcc/proto/go"
	"github.com/ljb021027/tcc/test"
	log "github.com/sirupsen/logrus"
)

func main() {
	tcClient := client.NewTcClient()
	serviceA := test.NewServiceA(&tcc.RmResource{
		Uri: "grpc://127.0.0.1:8001",
	})
	serviceB := test.NewServiceB(&tcc.RmResource{
		Uri: "grpc://127.0.0.1:8002",
	})
	rmServiceA, err := client.NewRmService(tcClient, serviceA)
	if err != nil {
		log.Fatal(err)
	}
	go rmServiceA.Listen()

	rmServiceB, err := client.NewRmService(tcClient, serviceB)
	if err != nil {
		log.Fatal(err)
	}
	go rmServiceB.Listen()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
}
