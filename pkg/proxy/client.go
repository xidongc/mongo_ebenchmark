package proxy

import (
	"github.com/xidongc-wish/mongoproxy/mprpc"
	"google.golang.org/grpc"
)

type Client struct {
	Config *Config
	activeConnection *grpc.ClientConn
	mprpc.MongoProxyClient
}
