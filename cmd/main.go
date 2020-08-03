package main

import (
	"context"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
	"github.com/xidongc-wish/mgo/bson"
	sku "github.com/xidongc/mongodb_ebenchmark/model/sku/service"
	"github.com/xidongc/mongodb_ebenchmark/model/sku/skupb"
	"github.com/xidongc/mongodb_ebenchmark/pkg/proxy"
)

func test() {
	var result skupb.Sku

	param := &proxy.QueryParam{
		Filter:      bson.M{"_id" : bson.ObjectIdHex("5f1905221110eb58159ecd1e")},
		FindOne:     true,
		Amp:         proxy.MicroAmplifier(),
	}

	ctx, cancel := context.WithCancel(context.Background())

	storage := *sku.NewClient(proxy.DefaultConfig(), cancel)

	defer func() {
		err := storage.Close()
		if err != nil {
			log.Error("error")
		}
	}()

	results, err := storage.Find(ctx, param)
	if err != nil {
		panic(err)
	}
	log.Info("start output... ")
	log.Info(results[0])
	err = mapstructure.Decode(results[0], &result)
	if err != nil {
		panic(err)
	}
	log.Infof("%+v", result)
}
