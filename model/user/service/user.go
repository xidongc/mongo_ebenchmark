package service

import (
	"context"
	"github.com/xidongc/mongodb_ebenchmark/model/user/userpb"
	"github.com/xidongc/mongodb_ebenchmark/pkg/proxy"
)

const ns = "user"

type Service struct {
	Storage proxy.Client
	Amplifier proxy.Amplifier
}

func (s Service) New(context.Context, *userpb.NewRequest) (*userpb.User, error) {
	panic("implement me")
}

func (s Service) Get(context.Context, *userpb.GetRequest) (*userpb.User, error) {
	panic("implement me")
}

func (s Service) Update(context.Context, *userpb.UpdateRequest) (*userpb.User, error) {
	panic("implement me")
}

func (s Service) Delete(context.Context, *userpb.DeleteRequest) (*userpb.Empty, error) {
	panic("implement me")
}

