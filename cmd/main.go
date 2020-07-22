package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/xidongc/mongodb_ebenchmark/pkg/proxy"
)

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
