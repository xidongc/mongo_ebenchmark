/*
 * mongodb_ebenchmark - Mongodb grpc proxy benchmark for e-commerce workload (still in dev)
 *
 * Copyright (c) 2020 - Chen, Xidong <chenxidong2009@hotmail.com>
 *
 * All rights reserved.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

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
	ctx, cancel := context.WithCancel(context.Background())
	client, _ := NewClient(nil, "test", cancel)

	defer func() {
		err := client.Close()
		if err != nil {
			t.Error(err)
		}
	}()

	if err := client.HealthCheck(); err != nil {
		panic(err)
	}

	t.Run(Insert, func(t *testing.T) {
		sku := skupb.Inventory{
			SkuId: 123,
			WarehouseId: 345,
			Quantity: 234,
			Type: 123,
		}
		skus := make([]interface{}, 1)
		skus[0] = sku
		param := &InsertParam{
			Docs: skus,
			Amp:  MicroAmplifier(),
		}
		if err := client.Insert(ctx, param); err != nil {
			t.Error("error")
		}
	})

	t.Run(FindIter, func(t *testing.T) {
		queryParam := &QueryParam{
			Filter:      bson.M{},
			Amp:         MicroAmplifier(),
		}

		var documents []interface{}

		stream, err := client.FindIter(ctx, queryParam)
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

	t.Run(Update, func(t *testing.T) {
		updateParam := &UpdateParam{
			Filter:      bson.M{"skuid": 123},
			Update:      bson.M{"skuid": 124},
			Multi: 		 false,
			Amp:         MicroAmplifier(),
		}
		changeInfo, err := client.Update(ctx, updateParam)
		if err != nil {
			t.Error(err)
		}
		t.Log(changeInfo)
	})
}
