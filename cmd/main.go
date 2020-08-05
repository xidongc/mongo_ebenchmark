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

package main

import (
	"context"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
	"github.com/xidongc-wish/mgo/bson"
	sku "github.com/xidongc/mongo_ebenchmark/model/sku/service"
	"github.com/xidongc/mongo_ebenchmark/model/sku/skupb"
	"github.com/xidongc/mongo_ebenchmark/pkg/proxy"
)

func test() {
	var result skupb.Sku

	param := &proxy.QueryParam{
		Filter:  bson.M{"_id": bson.ObjectIdHex("5f1905221110eb58159ecd1e")},
		FindOne: true,
		Amp:     proxy.MicroAmplifier(),
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
