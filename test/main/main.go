package main

import (
	"context"
	"net/http"

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

	tm := client.StartTm(tcClient, serviceA, serviceB)

	http.HandleFunc("/commit", func(writer http.ResponseWriter, request *http.Request) {
		ctx := context.Background()
		err := tm.StartTransaction(ctx)
		if err != nil {
			log.Error(err)
		}
	})
	http.HandleFunc("/rollback", func(writer http.ResponseWriter, request *http.Request) {

	})
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}

}
