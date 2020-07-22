package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	sku "github.com/xidongc/mongodb_ebenchmark/model/sku/service"
	"github.com/xidongc/mongodb_ebenchmark/model/sku/skupb"
	"github.com/xidongc/mongodb_ebenchmark/pkg/proxy"
	"google.golang.org/grpc"
)

func registerServices(server *grpc.Server) {
	// TODO
	skupb.RegisterSkuServiceServer(server, &sku.Service{})

}

func main() {
	_, cancel := context.WithCancel(context.Background())
	client, _ := proxy.NewClient(nil, "main", cancel)
	defer func() {
		err := client.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	if err := client.HealthCheck(); err != nil {
		panic(err)
	}


}
