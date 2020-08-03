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

package sku

import (
	"context"
	"errors"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
	"github.com/xidongc-wish/mgo/bson"
	"github.com/xidongc/mongodb_ebenchmark/model/sku/skupb"
	"github.com/xidongc/mongodb_ebenchmark/pkg/proxy"
)

const ns = "sku"

type Service struct {
	Storage   proxy.Client
	Amplifier proxy.Amplifier
}

func (s *Service) Get(ctx context.Context, req *skupb.GetRequest) (sku *skupb.Sku, err error) {
	param := &proxy.QueryParam{
		Filter:      bson.M{"_id": req.Id},
		FindOne:     true,
		Amp:         s.Amplifier,
	}

	results, err := s.Storage.Find(ctx, param)

	if err != nil || len(results) > 1 {
		log.Error(err)
		return
	} else if len(results) == 0 {
		return sku, errors.New("no result found")
	}

	err = mapstructure.Decode(results[0], sku)
	if err != nil {
		log.Fatal(err)
	}
	return sku, nil
}

func (s *Service) Update(ctx context.Context, req *skupb.UpdateRequest) (sku *skupb.Sku, err error) {
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

func (s *Service) Delete(ctx context.Context, req *skupb.DeleteRequest) (*skupb.Empty, error) {
	removeQuery := &proxy.RemoveParam{
		Filter: bson.M{"_id": req.Id},
		Amp:    s.Amplifier,
	}
	_, err := s.Storage.Remove(ctx, removeQuery)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &skupb.Empty{}, nil
}

func (s *Service) New(ctx context.Context, req *skupb.NewRequest) (sku *skupb.Sku, err error) {

	updateQuery := bson.M{
		"name":              req.GetName(),
		"price":             req.GetPrice(),
		"currency":          req.GetCurrency(),
		"active":            req.GetActive(),
		"productid":         req.GetParent(),
		"image":             req.GetImage(),
		"metadata":          req.GetMetadata(),
		"packagedimensions": req.GetPackageDimensions(),
		"attibutes":         req.GetAttributes(),
	}

	param := &proxy.UpdateParam{
		Filter: bson.M{"name": req.GetName()},
		Update: updateQuery,
		Upsert: true,
		Multi:  false,
		Amp:    s.Amplifier,
	}

	changeInfo, err := s.Storage.Update(ctx, param)
	if err != nil {
		log.Error("error")
	}

	log.Info(changeInfo)

	if err != nil {
		log.Errorf("sku error: storage failed with %s", err)
	}

	sku = &skupb.Sku{
		Name:              req.GetName(),
		Price:             req.GetPrice(),
		Currency:          req.GetCurrency(),
		Active:            req.GetActive(),
		ProductId:         req.GetParent(),
		Image:             req.GetImage(),
		Metadata:          req.GetMetadata(),
		PackageDimensions: req.GetPackageDimensions(),
		Attributes:        req.GetAttributes(),
	}

	return
}

func NewClient(config *proxy.Config, cancel context.CancelFunc) (client *proxy.Client) {
	client, _ = proxy.NewClient(config, ns, cancel)
	return
}
