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
	"errors"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
	"github.com/xidongc-wish/mgo/bson"
	"github.com/xidongc/mongo_ebenchmark/model/sku/skupb"
	"github.com/xidongc/mongo_ebenchmark/pkg/proxy"
)

const ns = "sku"

// SKU Service
type Service struct {
	Storage   proxy.Client
	Amplifier proxy.Amplifier
}

// Find SKU
func (s *Service) Get(ctx context.Context, req *skupb.GetRequest) (*skupb.Sku, error) {
	var sku skupb.Sku

	param := &proxy.QueryParam{
		Filter:  bson.M{"Name": req.GetName()},
		FindOne: true,
		Amp:     s.Amplifier,
	}

	results, err := s.Storage.Find(ctx, param)

	if err != nil || len(results) > 1 {
		log.Error(err)
		return &sku, err
	} else if len(results) == 0 {
		return &sku, errors.New("no result found")
	}

	err = mapstructure.Decode(results[0], &sku)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("received sku: %+v", sku)
	return &sku, nil
}

// Delete SKU
func (s *Service) Delete(ctx context.Context, req *skupb.DeleteRequest) (*skupb.Empty, error) {
	removeQuery := &proxy.RemoveParam{
		Filter: bson.M{"Name": req.GetName()},
		Amp:    s.Amplifier,
	}
	changeInfo, err := s.Storage.Remove(ctx, removeQuery)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Info(changeInfo)
	return &skupb.Empty{}, nil
}

// Create SKU inserts if SKU not recorded, or update sku if exist
func (s *Service) New(ctx context.Context, req *skupb.UpsertRequest) (sku *skupb.Sku, err error) {

	var inventories []*skupb.Inventory

	query := proxy.QueryParam{
		Filter:  bson.M{"Name": req.Name},
		FindOne: true,
		Amp:     nil,
	}

	result, err := s.Storage.Find(ctx, &query)

	if err != nil || len(result) > 1 {
		log.Fatal(err)
		return
	} else if len(result) == 0 {
		inventories = append(inventories, req.GetInventory())
	} else if _, ok := result[0]["Inventory"]; ok && len(result) == 1 {
		if err := mapstructure.Decode(result[0], &sku); err == nil {
			inventories = append(inventories, sku.Inventory...)
			inventories = append(inventories, req.GetInventory())
		}
	}

	sku = &skupb.Sku{
		Name:              req.GetName(),
		Price:             req.GetPrice(),
		Currency:          req.GetCurrency(),
		Active:            req.GetActive(),
		ProductId:         req.GetProductId(),
		Image:             req.GetImage(),
		SkuLabel:          req.GetProductId(),
		Metadata:          req.GetMetadata(),
		Inventory:         inventories,
		PackageDimensions: req.GetPackageDimensions(),
		Attributes:        req.GetAttributes(),
		HasBattery:        req.GetHasBattery(),
		HasSensitive:      req.GetHasSensitive(),
		HasLiquid:         req.GetHasLiquid(),
		Description:       req.GetDescription(),
		Supplier:          req.GetSupplier(),
	}

	var updateQuery bson.M

	if err = mapstructure.Decode(sku, &updateQuery); err != nil {
		log.Error(err)
		return
	}

	param := &proxy.UpdateParam{
		Filter: bson.M{"Name": req.GetName()},
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

	return
}

// Create SKU Service client
func NewClient(config *proxy.Config, cancel context.CancelFunc) (client *proxy.Client) {
	client, _ = proxy.NewClient(config, ns, cancel)
	return
}
