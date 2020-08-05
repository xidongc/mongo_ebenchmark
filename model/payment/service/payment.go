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
	"github.com/xidongc/mongo_ebenchmark/model/payment/paymentpb"
	"github.com/xidongc/mongo_ebenchmark/pkg/proxy"
)

const ns = "payment"

type Service struct {
	Storage   proxy.Client
	Amplifier proxy.Amplifier
}

// New Charge
func (s Service) NewCharge(ctx context.Context, req *paymentpb.ChargeRequest) (*paymentpb.Charge, error) {
	charge := paymentpb.Charge{
		Currency:     req.GetCurrency(),
		ChargeAmount: req.GetAmount(),
	}

	var docs []interface{}
	docs = append(docs, charge)

	param := &proxy.InsertParam{
		Docs: docs,
		Amp:  s.Amplifier,
	}

	if err := s.Storage.Insert(ctx, param); err != nil {
		log.Error(err)
		return nil, err
	}
	return &charge, nil
}

// Refund Charge
func (s Service) RefundCharge(ctx context.Context, req *paymentpb.RefundRequest) (*paymentpb.Charge, error) {
	charge := paymentpb.Charge{
		ChargeAmount: -req.GetAmount(),
	}

	var docs []interface{}
	docs = append(docs, charge)

	param := &proxy.InsertParam{
		Docs: docs,
		Amp:  s.Amplifier,
	}

	if err := s.Storage.Insert(ctx, param); err != nil {
		log.Error(err)
		return nil, err
	}
	return &charge, nil
}

func (s Service) Get(ctx context.Context, req *paymentpb.GetRequest) (charge *paymentpb.Charge, err error) {
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
		return charge, errors.New("no result found")
	}

	err = mapstructure.Decode(results[0], charge)
	if err != nil {
		log.Fatal(err)
	}
	return charge, nil
}
