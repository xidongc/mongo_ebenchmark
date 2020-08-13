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
	log "github.com/sirupsen/logrus"
	"github.com/xidongc/mongo_ebenchmark/model/payment/paymentpb"
	"github.com/xidongc/mongo_ebenchmark/model/payment/service/provider"
	"github.com/xidongc/mongo_ebenchmark/pkg/proxy"
)

const ns = "payment"

type Service struct {
	Storage   proxy.Client
	Amplifier  proxy.Amplifier
}

// New Charge
func (s Service) NewCharge(ctx context.Context, req *paymentpb.ChargeRequest) (*paymentpb.Charge, error) {
	providerId := req.GetPaymentProviderId()
	var provide Provider

	// TODO add more
	switch providerId{
	case paymentpb.PaymentProviderId_AliPay:
		provide = &provider.AliPay{}
	default:
		provide = &provider.AliPay{}
	}

	charge, err := provide.Charge(req)

	if err != nil {
		log.Warning("do refund")
		// TODO do refund
	}

	if charge == nil {
		return nil, errors.New("error")
	}

	var docs []interface{}
	docs = append(docs, *charge)

	param := &proxy.InsertParam{
		Docs: docs,
		Amp:  s.Amplifier,
	}

	if err := s.Storage.Insert(ctx, param); err != nil {
		log.Error(err)
		return nil, err
	}
	return charge, nil
}

// Refund Charge
func (s Service) RefundCharge(ctx context.Context, req *paymentpb.RefundRequest) (charge *paymentpb.Charge, err error) {
	return
}

// Get Charge
func (s Service) Get(ctx context.Context, req *paymentpb.GetRequest) (charge *paymentpb.Charge, err error) {
	return
}

// Create Payment Service client
func NewClient(config *proxy.Config, cancel context.CancelFunc) (client *proxy.Client) {
	client, _ = proxy.NewClient(config, ns, cancel)
	return
}
