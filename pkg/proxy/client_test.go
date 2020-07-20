package proxy

import (
	"github.com/xidongc-wish/mgo/bson"
	"github.com/xidongc/mongodb_ebenchmark/model/sku/skupb"
	"testing"
)

// Test on basic proxy operations
func TestClientOps(t *testing.T) {
	client, _ := NewClient(nil)

	t.Run(Insert, func(t *testing.T) {
		sku := skupb.Inventory{
			SkuId: 123,
			WarehouseId: 345,
			Quantity: 234,
			Type: 123,

		}
		skus := make([]interface{}, 1)
		skus[0] = sku
		if err := client.Insert(skus, MicroAmplifier); err != nil {
			t.Error("error")
		}
	})

	t.Run(FindIter, func(t *testing.T) {
		var filter bson.M
		filter = bson.M{}
		if err := client.FindIter(filter, MicroAmplifier); err != nil {
			t.Error("error")
		}
	})
}
