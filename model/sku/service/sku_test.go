/*
 * mongodb_ebenchmark - Mongodb grpc proxy benchmark for e-commerce workload (still in dev)
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
 *
 */

package sku

import (
	"context"
	"github.com/mitchellh/mapstructure"
	"github.com/xidongc-wish/mgo/bson"
	"github.com/xidongc/mongo_ebenchmark/model/sku/skupb"
	"github.com/xidongc/mongo_ebenchmark/pkg/proxy"
	"testing"
)

func TestSkuServiceApi(t *testing.T) {
	var result skupb.Sku

	param := &proxy.QueryParam{
		Filter:  bson.M{"Name": "xidong"},
		FindOne: true,
		Amp:     proxy.MicroAmplifier(),
	}

	ctx, cancel := context.WithCancel(context.Background())

	storage := NewClient(proxy.DefaultConfig(), cancel)

	defer func() {
		err := storage.Close()
		if err != nil {
			t.Error("error")
		}
	}()

	results, err := storage.Find(ctx, param)
	if err != nil {
		panic(err)
	}
	t.Log("start output... ")
	t.Log(results[0])
	err = mapstructure.Decode(results[0], &result)
	if err != nil {
		panic(err)
	}
	t.Logf("%+v", result)
}
