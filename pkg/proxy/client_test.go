package proxy

import (
	"context"
	"github.com/xidongc-wish/mgo/bson"
	"github.com/xidongc/mongodb_ebenchmark/model/sku/skupb"
	"testing"
	"time"
)

// Test on basic proxy operations
func TestClientOps(t *testing.T) {
	client, _ := NewClient(nil)

	t.Run(Insert, func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		sku := skupb.Inventory{
			SkuId: 123,
			WarehouseId: 345,
			Quantity: 234,
			Type: 123,
		}
		skus := make([]interface{}, 1)
		skus[0] = sku
		if err := client.Insert(ctx, skus, MicroAmplifier); err != nil {
			t.Error("error")
		}
	})

	t.Run(FindIter, func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		var filter bson.M
		var documents []interface{}
		filter = bson.M{}
		stream, err := client.FindIter(ctx, filter, MicroAmplifier)
		if err != nil {
			t.Error("error")
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
						t.Fatal(err)
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
		t.Log(documents)
	})
}
