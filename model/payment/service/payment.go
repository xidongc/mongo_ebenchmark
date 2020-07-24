package service

import (
	"context"
	"github.com/xidongc/mongodb_ebenchmark/model/payment/paymentpb"
	"github.com/xidongc/mongodb_ebenchmark/pkg/proxy"
)

const ns = "payment"

type Service struct {
	Storage proxy.Client
	Amplifier proxy.Amplifier
}

func (s Service) NewCharge(context.Context, *paymentpb.ChargeRequest) (*paymentpb.Charge, error) {
	panic("implement me")
}

func (s Service) RefundCharge(context.Context, *paymentpb.RefundRequest) (*paymentpb.Charge, error) {
	panic("implement me")
}

func (s Service) Get(context.Context, *paymentpb.GetRequest) (*paymentpb.Charge, error) {
	panic("implement me")
}


