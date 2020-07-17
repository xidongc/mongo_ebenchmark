package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/xidongc-wish/mgo/bson"
	"github.com/xidongc/mongodb_ebenchmark/pkg/proxy"
)

func main() {
	client, _ := proxy.NewClient(nil)
	var filter bson.M
	filter = bson.M{}
	if err := client.FindIter(filter, proxy.MicroAmplifier); err != nil {
		log.Error("error")
	}
}
