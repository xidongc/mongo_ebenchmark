package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/xidongc-wish/mgo/bson"
	"github.com/xidongc/mongodb_ebenchmark/pkg/proxy"
	"time"
)

func main() {
	client, _ := proxy.NewClient(nil)
	if err := client.HealthCheck(); err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var filter bson.M
	var documents []interface{}
	filter = bson.M{}
	stream, err := client.FindIter(ctx, filter, proxy.MicroAmplifier)
	if err != nil {
		log.Error("error")
	}
	for {
		if stream == nil {
			break
		}
		rs, err := stream.Recv()
		if err != nil {
			break
		}
		tmpDocs := make(chan interface{}, len(rs.Results))
		if len(rs.Results) > 0 {
			for _, document := range rs.Results {
				var doc bson.D
				if err := bson.Unmarshal(document.Val, &doc); err != nil {
					log.Fatal(err)
				}
				tmpDocs <- doc
			}

			for i := 0; i < len(rs.Results); i++ {
				doc := <-tmpDocs
				documents = append(documents, doc)
			}
		} else {
			time.Sleep(5 * time.Millisecond)
		}
	}
	log.Info(documents)
}
