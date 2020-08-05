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
	"github.com/xidongc/mongo_ebenchmark/pkg/proxy"
)

const ns = "product"

type Service struct {
	Storage   proxy.Client
	Amplifier proxy.Amplifier
}

// Create a product
func (s Service) New(ctx context.Context, req *productpb.NewRequest) (*productpb.Product, error) {
	product := productpb.Product{
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
	return &product, nil
}

// get a product
func (s Service) Get(ctx context.Context, req *productpb.GetRequest) (product *productpb.Product, err error) {

	param := &proxy.QueryParam{
		Filter:  bson.M{"_id": req.Id},
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

	err = mapstructure.Decode(results[0], product)
	if err != nil {
		log.Fatal(err)
	}
	return product, nil
}

// update a product
func (s Service) Update(ctx context.Context, req *productpb.UpdateRequest) (product *productpb.Product, err error) {
	var updateParams bson.M
	if err = mapstructure.Decode(req, &updateParams); err != nil {
		log.Error(err)
		return
	}
	updateQuery := &proxy.UpdateParam{
		Filter: bson.M{"_id": req.Id},
		Update: updateParams,
		Upsert: false,
		Multi:  true,
		Amp:    s.Amplifier,
	}
	_, err = s.Storage.Update(ctx, updateQuery)
	if err != nil {
		return
	}
	// TODO not return product
	return
}

// Delete a product
func (s Service) Delete(ctx context.Context, req *productpb.DeleteRequest) (*productpb.Empty, error) {
	removeQuery := &proxy.RemoveParam{
		Filter: bson.M{"_id": req.Id},
		Amp:    s.Amplifier,
	}
	_, err := s.Storage.Remove(ctx, removeQuery)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &productpb.Empty{}, nil
}
