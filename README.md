<p align="left">
  <a href="https://goreportcard.com/report/github.com/xidongc/mongo_ebenchmark"><img src="https://goreportcard.com/badge/github.com/xidongc/mongo_ebenchmark"></a>
</p>

# Mongodb eBenchmark

Mongodb grpc proxy benchmark for e-commerce workload (still in dev)

-------------------------
- [Design](#design)
- [Introduction](#introduction)
- [Usage](#usage)
- [License](#license)

-------------------------
## design

Mongodb eBenchmark is targeted to use and amplify common e-commerce traffic to database grpc proxy. by design, it can be divided into following five layers:

- Server: register model services to grpc server, and handle requests
- Model: Normal e-commerce model consist of order, product, sku, payment and user
- Middleware: Amplify Database Workload based on AmplifyOptions
- Proxy: Expose grpc requests to middleware, and route traffic to Database
- Database: Support MongoDB 3.2 and higher

-------------------------
## Introduction

Mongodb eBenchmark start a grpc server with basic e-commerce modules: order, product,
payment, sku, user, please refer `model` folder, data structure is inspired by digota
`github.com/digota/digota`, with a service for each module 

in Middleware, Mongodb eBenchmark uses ghz `github.com/bojand/ghz` to amplify db request to proxy via grpc protocol
The granularity of workload amplification is per function, via AmplifyOptions settings

```go
type AmplifyOptions struct {
	Connections    uint			
	Concurrency    uint			
	TotalRequest   uint			
	QPS            uint			
	Timeout        time.Duration	
	CPUs           uint			
}
```

client will call database proxy api via grpc, besides from real request, it fake benchmark 
request based on given AmplifyOptions, undo is designed to calculate opposite request to make, 
in order to keep database clean after doing benchmark, proxy rpc server is written in a private
repo `github.com/xidongc-wish/mp-server` with limited access only. 

the project is managed by go mod, and can be installed by running

```bash
go get github.com/xidongc/mongo_ebenchmark
```

-------------------------
## Usage

To get started with, `go run cmd/server.go` with start an e-commerce server with benchmark
To create a new service:
```go
service := &module.Service{
                  Storage: storageClient,
                  Amplifier: amplifyOptions,
	      }
servicepb.RegisterServiceServer(svr, service)
```

when in turbo mode with `storageClient.Turbo` enabled, mongo client connection will use eventual 
consistency mode to maximaize throughput with consistency trade off, database driver uses 
`github.com/xidongc/mgo`, originally fork from `github.com/go-mgo/mgo` eg:

```go
if client.Turbo {
		readConcern = "local"
		prefetch = 0.75
		readPref = mgo.Nearest
	} else {
		readConcern = "linearizable"
		prefetch = 0.25
		readPref = mgo.Primary
	}
```

please refer to  `pkg/client/client.go` for database grpc client api support:

-------------------------
## License

The Mongodb eBenchmark is licensed under the [Apache License](LICENSE).
