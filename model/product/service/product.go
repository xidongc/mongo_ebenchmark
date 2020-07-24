package service

import (
	"context"
	"github.com/xidongc/mongodb_ebenchmark/model/product/productpb"
	"github.com/xidongc/mongodb_ebenchmark/pkg/proxy"
)

const ns = "product"

type Service struct {
	Storage proxy.Client
	Amplifier proxy.Amplifier
}

func (s Service) New(context.Context, *productpb.NewRequest) (*productpb.Product, error) {
	panic("implement me")
}

func (s Service) Get(context.Context, *productpb.GetRequest) (*productpb.Product, error) {
	panic("implement me")
}

func (s Service) Update(context.Context, *productpb.UpdateRequest) (*productpb.Product, error) {
	panic("implement me")
}

func (s Service) Delete(context.Context, *productpb.DeleteRequest) (*productpb.Empty, error) {
	panic("implement me")
}


