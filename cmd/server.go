package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	sku "github.com/xidongc/mongodb_ebenchmark/model/sku/service"
	"github.com/xidongc/mongodb_ebenchmark/model/sku/skupb"
	"github.com/xidongc/mongodb_ebenchmark/pkg/proxy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
	// gRPC server
	maxSendMsgSize := 1024 * 1024 * 500
	maxRecvMsgSize := 1024 * 1024 * 100
	maxSendMsgSizeOpt := grpc.MaxSendMsgSize(maxSendMsgSize)
	maxRecvMsgSizeOpt := grpc.MaxRecvMsgSize(maxRecvMsgSize)

	server := grpc.NewServer(maxSendMsgSizeOpt, maxRecvMsgSizeOpt)
	skuService := sku.NewSKUService(nil, proxy.MicroAmplifier, nil)
	skupb.RegisterSkuServiceServer(server, skuService)
	reflection.Register(server)

	go func() {
		addr := fmt.Sprintf(":%d", 50053)
		lis, err := net.Listen("tcp4", addr)
		if err != nil {
			log.Fatal(err)
		}
		if err = server.Serve(lis); err != nil {
			log.Fatal(err)
		}
		cancel()
	}()
	select {
	case <-sigs:
	case <-ctx.Done():
	}
}
