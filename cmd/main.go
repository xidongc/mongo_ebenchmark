package main

import (
	"github.com/xidongc/mongodb_ebenchmark/pkg/proxy"
)

func main() {
	client, _ := proxy.NewClient(nil)
	if err := client.HealthCheck(); err != nil {
		panic(err)
	}
}
