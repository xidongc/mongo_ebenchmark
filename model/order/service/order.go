package service

import (
	"context"
	"github.com/xidongc/mongodb_ebenchmark/model/order/orderpb"
	"github.com/xidongc/mongodb_ebenchmark/pkg/proxy"
)

const ns = "order"

type Service struct {
	Storage proxy.Client
	Amplifier proxy.Amplifier
}

func (s Service) New(context.Context, *orderpb.NewRequest) (*orderpb.Order, error) {
	panic("implement me")
}

func (s Service) Get(context.Context, *orderpb.GetRequest) (*orderpb.Order, error) {
	panic("implement me")
}

func (s Service) Pay(context.Context, *orderpb.PayRequest) (*orderpb.Order, error) {
	panic("implement me")
}

func (s Service) Return(context.Context, *orderpb.ReturnRequest) (*orderpb.Order, error) {
	panic("implement me")
}


