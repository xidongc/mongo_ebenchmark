package main

import (
	"context"
	"fmt"
	flags "github.com/jessevdk/go-flags"
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
	var config proxy.Config

	parser := flags.NewParser(&config, flags.Default)
	if _, err := parser.Parse(); err != nil {
		os.Exit(1)
	}
	log.Infof("%+v", config)

	ctx, cancel := context.WithCancel(context.Background())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	maxSendMsgSize := 1024 * 1024 * 500
	maxRecvMsgSize := 1024 * 1024 * 100

	maxSendMsgSizeOpt := grpc.MaxSendMsgSize(maxSendMsgSize)
	maxRecvMsgSizeOpt := grpc.MaxRecvMsgSize(maxRecvMsgSize)

	server := grpc.NewServer(maxSendMsgSizeOpt, maxRecvMsgSizeOpt)
	storageClient := *sku.NewClient(&config, cancel)
	defer func() {
		err := storageClient.Close()
		if err != nil {
			log.Error("error")
		}
	}()
	skuService := &sku.Service{
		Storage: storageClient,
		Amplifier: proxy.MicroAmplifier,
	}
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

	log.Warn("Got shutdown signal")
	server.GracefulStop()
}
