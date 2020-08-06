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

package service

import (
	"context"
	"errors"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
	"github.com/xidongc-wish/mgo/bson"
	"github.com/xidongc/mongo_ebenchmark/model/product/productpb"
	skuService "github.com/xidongc/mongo_ebenchmark/model/sku/service"
	"github.com/xidongc/mongo_ebenchmark/model/sku/skupb"
	"github.com/xidongc/mongo_ebenchmark/pkg/proxy"
	"strings"
)

const ns = "product"

type Service struct {
	Storage   proxy.Client
	Amplifier  proxy.Amplifier
	SkuService *skuService.Service
}

// Create Product
func (s Service) New(ctx context.Context, req *productpb.NewRequest) (*productpb.Product, error) {
	if _, err := s.Get(ctx, &productpb.GetRequest{Id: req.Id}); err == nil {
		return nil, errors.New("please use update for exist record")
	}

	product := productpb.Product{
		Id:			 req.GetId(),
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Shippable:   req.GetShippable(),
		Images:      req.GetImages(),
		Attributes:  req.GetAttributes(),
		Metadata:    req.GetMetadata(),
		Active:      req.GetActive(),
		Url:         req.GetUrl(),
	}

	var docs []interface{}
	docs = append(docs, product)

	param := &proxy.InsertParam{
		Docs: docs,
		Amp:  s.Amplifier,
	}

	if err := s.Storage.Insert(ctx, param); err != nil {
		log.Error(err)
		return nil, err
	}
	log.Infof("%+v", product)
	return &product, nil
}

// Get product
func (s Service) Get(ctx context.Context, req *productpb.GetRequest) (product *productpb.Product, err error) {

	param := &proxy.QueryParam{
		Filter:  bson.M{"id": req.Id},
		FindOne: true,
		Amp:     s.Amplifier,
	}

	results, err := s.Storage.Find(ctx, param)

	if err != nil || len(results) > 1 {
		log.Error(err)
		return
	} else if len(results) == 0 {
		return product, errors.New("no result found")
	}

	err = mapstructure.Decode(results[0], &product)
	if err != nil || product == nil {
		log.Fatal(err)
	}

	skus, err := s.SkuService.GetProductSkus(ctx, &skupb.GetProductSkusRequest{ProductId: req.Id})
	if skus != nil {
		product.Skus = skus.GetSkus()
	}
	return product, nil
}

// Update product and return updated
func (s Service) Update(ctx context.Context, req *productpb.UpdateRequest) (product *productpb.Product, err error) {
	var updateParams bson.M
	if err = mapstructure.Decode(req, &updateParams); err != nil {
		log.Error(err)
		return
	}
	updateLowerParams := make (bson.M, len(updateParams))
	for key, val := range updateParams {
		r := []rune(key)
		key = strings.ToLower(string(r[0])) + string(r[1:])
		updateLowerParams[strings.ToLower(key)] = val
	}
	updateQuery := &proxy.UpdateParam{
		Filter: bson.M{"id": req.Id},
		Update: updateLowerParams,
		Upsert: false,
		Multi:  false,
		Amp:    s.Amplifier,
	}
	changeInfo, err := s.Storage.Update(ctx, updateQuery)
	if err != nil {
		return
	}
	log.Info(changeInfo)
	product, err = s.Get(ctx, &productpb.GetRequest{Id: req.Id})
	if err != nil {
		log.Error(err)
	}
	return
}

// Delete product
func (s Service) Delete(ctx context.Context, req *productpb.DeleteRequest) (*productpb.Empty, error) {
	removeQuery := &proxy.RemoveParam{
		Filter: bson.M{"id": req.Id},
		Amp:    s.Amplifier,
	}
	_, err := s.Storage.Remove(ctx, removeQuery)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	skus, err := s.SkuService.GetProductSkus(ctx, &skupb.GetProductSkusRequest{ProductId: req.Id})
	if skus == nil {
		log.Fatal(err)
		return nil, err
	}
	for _, sku := range skus.GetSkus() {
		_, err = s.SkuService.Delete(ctx, &skupb.DeleteRequest{Name: sku.GetName()})
	}
	return &productpb.Empty{}, nil
}

// Create Product Service client
func NewClient(config *proxy.Config, cancel context.CancelFunc) (client *proxy.Client) {
	client, _ = proxy.NewClient(config, ns, cancel)
	return
}
