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
	"github.com/xidongc/mongo_ebenchmark/model/order/orderpb"
	"github.com/xidongc/mongo_ebenchmark/model/payment/paymentpb"
	payment "github.com/xidongc/mongo_ebenchmark/model/payment/service"
	"github.com/xidongc/mongo_ebenchmark/pkg/proxy"
)

const ns = "order"

type Service struct {
	Storage   proxy.Client
	Payment   payment.Service
	Amplifier proxy.Amplifier
}

// Create Order
func (s Service) New(ctx context.Context, req *orderpb.NewRequest) (*orderpb.Order, error) {
	order := orderpb.Order{
		Currency: req.Currency,
		Items:    req.Items,
		Metadata: req.Metadata,
		Shipping: req.Shipping,
	}
	var orders []interface{}
	orders = append(orders, order)
	insertQuery := &proxy.InsertParam{
		Docs: orders,
		Amp:  s.Amplifier,
	}
	err := s.Storage.Insert(ctx, insertQuery)
	if err != nil {
		log.Error(err)
	}
	return &order, nil
}

// Get Order by id
func (s Service) Get(ctx context.Context, req *orderpb.GetRequest) (order *orderpb.Order, err error) {

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
		return order, errors.New("no result found")
	}

	err = mapstructure.Decode(results[0], order)
	if err != nil {
		log.Fatal(err)
	}
	return order, nil
}

// Pay
func (s Service) Pay(ctx context.Context, req *orderpb.PayRequest) (order *orderpb.Order, err error) {
	chargeRequest := &paymentpb.ChargeRequest{
		PaymentProviderId: req.GetPaymentProviderId(),
		Card:              req.GetCard(),
		// Amount:  uint64(o.GetAmount()),  TODO need to define order id
	}
	charge, err := s.Payment.NewCharge(ctx, chargeRequest)
	if err != nil {
		// TODO do refund logic
		log.Fatal(err)
	}
	log.Info(charge)
	return
}

func (s Service) Return(ctx context.Context, req *orderpb.ReturnRequest) (order *orderpb.Order, err error) {
	// TODO
	return
}
