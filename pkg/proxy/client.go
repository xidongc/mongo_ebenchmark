package proxy

import (
	"context"
	"fmt"
	"github.com/bojand/ghz/printer"
	"github.com/bojand/ghz/runner"
	log "github.com/sirupsen/logrus"
	"github.com/xidongc-wish/mgo"
	"github.com/xidongc-wish/mgo/bson"
	"github.com/xidongc-wish/mongoproxy/mprpc"
	"google.golang.org/grpc"
	"os"
)

// Function List in Proxy
const (
	Find          = "mprpc.MongoProxy.Find"
	FindIter      = "mprpc.MongoProxy.FindIter"
	Count         = "mprpc.MongoProxy.Count"
	Explain       = "mprpc.MongoProxy.Explain"
	Aggregate     = "mprpc.MongoProxy.Aggregate"
	Bulk          = "mprpc.MongoProxy.Bulk"
	Update        = "mprpc.MongoProxy.Update"
	Remove        = "mprpc.MongoProxy.Remove"
	Insert        = "mprpc.MongoProxy.Insert"
	FindAndModify = "mprpc.MongoProxy.FindAndModify"
	Distinct      = "mprpc.MongoProxy.Distinct"
	Healthcheck   = "mprpc.MongoProxy.Healthcheck"
)

// Database used for ebenchmark
const (
	Database 	  = "ebenchmark"
	Collection 	  = "default"
)

// TODO define as const, should be in env
const (
	ProtoFile     = "/Users/derekchen/go/src/github.com/xidongc/mongodb_ebenchmark/pkg/proxy/rpc.protoset"
)

type FindAndModifyMode int

// Mode for FindAndModify
const (
	FindAndDelete 	FindAndModifyMode = 0
	FindAndUpdate   FindAndModifyMode = 1
	FindAndUpsert   FindAndModifyMode = 2
)

// Proxy Client
type Client struct {
	config *Config
	Host string
	Collection *mprpc.Collection
	Turbo bool
	activeCon *grpc.ClientConn
	rpcClient mprpc.MongoProxyClient
	cancelFunc context.CancelFunc
	isHealthy bool
}

type Empty struct {}

// Query param for upper services
type QueryParam struct {
	Filter 			bson.M
	Fields 			bson.M
	Limit			int64
	Skip 			int64
	Sort 			[]string
	Distinctkey 	string
	FindOne			bool
	UsingIndex		[]string
	Amp 			Amplifier
}

// Insert param for upper services
type InsertParam struct {
	Docs 			[]interface{}
	Amp 			Amplifier
}

// Remove param for upper services
type RemoveParam struct {
	Filter			bson.M
	Amp 			Amplifier
}

// Update param for upper services
type UpdateParam struct {
	Filter			bson.M
	Update			bson.M
	Upsert          bool
	Multi           bool
	Amp 			Amplifier
}

// FindAndModify param for upper services
type FindModifyParam struct {
	Filter			bson.M
	Desired			bson.M
	Mode			FindAndModifyMode
	SortRule		[]string
	Fields 			bson.M
	Amp 			Amplifier
}

// Aggregate param for upper services
type AggregateParam struct {
	Pipeline 		bson.M 
}

// Create a new client based on provided config
func NewClient(config *Config, namespace string, cancel context.CancelFunc) (client *Client, err error) {
	if config == nil {
		config = DefaultConfig()
	}
	if namespace == "" {
		namespace = Collection
	}
	if config.Port < 0 || config.Port == 0 {
		log.Warning("port not defined properly, set back to default")
		config.Port = 50051
	}

	host := fmt.Sprintf("%s:%d", config.ServerIp, config.Port)
	client = &Client{
		config: config,
		Host: host,
	}
	if client.Collection == nil {
		client.Collection = &mprpc.Collection{
			Database:   Database,
			Collection: namespace,
		}
	}
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect to rpc server error: %s", err)
	}
	client.activeCon = conn
	client.rpcClient = mprpc.NewMongoProxyClient(conn)
	if cancel == nil {
		log.Warning("handle context exit in application")
	}
	client.cancelFunc = cancel
	return
}

// Close a proxy client
func (client *Client) Close() (err error) {
	if err := client.activeCon.Close(); err != nil {
		log.Fatalf("clean up failed: %s", err)
	}
	if client.cancelFunc != nil {
		client.cancelFunc()
	}
	return
}

func (client *Client) Find(ctx context.Context, query *QueryParam) (docs []interface{}, err error) {
	return
}

func (client *Client) FindIter(ctx context.Context, query *QueryParam) (stream mprpc.MongoProxy_FindIterClient, err error) {
	filterBytes, err := bson.Marshal(query.Filter)
	if err != nil {
		log.Errorf("%s: marshall filter error", FindIter)
	}
	var readConcern string
	var prefetch float64
	var readPref mgo.Mode

	if client.Turbo {
		readConcern = "local"
		prefetch = 0.75
		readPref = mgo.Nearest
	} else {
		readConcern = "linearizable"
		prefetch = 0.25
		readPref = mgo.Primary
	}

	request := mprpc.FindQuery{
		Collection:  client.Collection,
		Filter:      filterBytes,
		Skip:        0,
		Maxtimems:   -1,
		Maxscan: 	 0,
		Prefetch:	 prefetch,
		Batchsize:   client.config.BatchSize,
		Readpref:    int32(readPref),
		Findone:     false,
		Partial:     client.config.AllowPartial,
		Readconcern: readConcern,
		Comment:     FindIter,
		Rpctimeout:  client.config.RpcTimeout,
	}
	stream, err = client.rpcClient.FindIter(ctx, &request)
	if err != nil {
		log.Fatalf("find iter call failed with: %s", err)
		return
	}
	report, err := runner.Run(
		FindIter,
		client.Host,
		runner.WithProtoset(ProtoFile),
		runner.WithConcurrency(query.Amp.Concurrency),
		runner.WithConnections(query.Amp.Connections),
		runner.WithCPUs(query.Amp.CPUs),
		runner.WithData(request),
		runner.WithInsecure(client.config.Insecure),
		)
	if err != nil {
		log.Fatal(err.Error())
	}
	p := printer.ReportPrinter{
		Out:    os.Stdout,
		Report: report,
	}

	_ = p.Print("pretty")
	return
}

func (client *Client) Count(ctx context.Context, query *QueryParam) (count uint64, err error) {
	return
}

func (client *Client) Explain(ctx context.Context, query *QueryParam) (explainFields bson.M, err error){
	return
}

func (client *Client) Aggregate(ctx context.Context, query *AggregateParam) (documents []interface{}, err error){
	return
}

func (client *Client) Update(ctx context.Context, param *UpdateParam) (changeInfo *mprpc.ChangeInfo, err error) {
	var wOptions *mprpc.WriteOptions
	if client.Turbo {
		wOptions = getTurboWriteOptions()
	} else {
		wOptions = getSafeWriteOptions()
	}

	filter, err := bson.Marshal(param.Filter)
	if err != nil {
		log.Fatal(err)
	}
	update, err := bson.Marshal(param.Update)
	if err != nil {
		log.Fatal(err)
	}

	request := &mprpc.UpdateOperation{
		Collection:   client.Collection,
		Filter:       filter,
		Update:       update,
		Upsert:       param.Upsert,
		Multi:        param.Multi,
		Writeoptions: wOptions,
	}

	changeInfo, err = client.rpcClient.Update(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (client *Client) Remove(ctx context.Context, param *RemoveParam) (changeInfo *mprpc.ChangeInfo, err error) {
	return
}

func (client *Client) Insert(ctx context.Context, param *InsertParam) (err error) {

	var rpcDocs []*mprpc.Document
	for _, doc := range param.Docs {
		val, err := bson.Marshal(doc)
		if err != nil {
			log.Panicf("unable to marshall error: %s", err)
		}
		rpcDocs = append(rpcDocs, &mprpc.Document{
			Val: val,
		})
	}
	var wOptions *mprpc.WriteOptions
	if client.Turbo {
		wOptions = getTurboWriteOptions()
	} else {
		wOptions = getSafeWriteOptions()
	}
	request := mprpc.InsertOperation{
		Collection:   client.Collection,
		Documents:  rpcDocs,
		Writeoptions: wOptions,
	}
	if _, err = client.rpcClient.Insert(ctx, &request); err != nil {
		log.Errorf("rpc insert error with: %s", err)
		return
	}

	report, err := runner.Run(
		Insert,
		client.Host,
		runner.WithProtoset(ProtoFile),
		runner.WithConcurrency(param.Amp.Concurrency),
		runner.WithConnections(param.Amp.Connections),
		runner.WithCPUs(param.Amp.CPUs),
		runner.WithData(request),
		runner.WithInsecure(client.config.Insecure),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	p := printer.ReportPrinter{
		Out:    os.Stdout,
		Report: report,
	}

	_ = p.Print("pretty")
	return
}

func (client *Client) FindAndModify(ctx context.Context, param *FindModifyParam) (singleDoc interface{}, err error) {
	filterBytes, err := bson.Marshal(param.Filter)
	if err != nil {
		log.Errorf("%s: marshall filter error", FindAndModify)
	}

	updateBytes, err := bson.Marshal(param.Desired)
	if err != nil {
		log.Errorf("%s: marshall filter error", FindAndModify)
	}

	var wOptions *mprpc.WriteOptions
	if client.Turbo {
		wOptions = getTurboWriteOptions()
	} else {
		wOptions = getSafeWriteOptions()
	}

	request := mprpc.FindAndModifyOperation{
		Collection:   client.Collection,
		Filter:       filterBytes,
		Update:       updateBytes,
		Upsert:       true,
		New:          true,
		Fields:       nil,
		Writeoptions: wOptions,
	}

	singleDoc, err = client.rpcClient.FindAndModify(ctx, &request)

	log.Info(singleDoc)

	return
}

func (client *Client) Distinct(ctx context.Context, query *QueryParam) (distinctKeys []interface{}, err error) {
	return
}

func (client *Client) HealthCheck() (err error) {
	log.Info("Start doing health check")
	empty := Empty{}
	_, err = runner.Run(
		Healthcheck,
		client.Host,
		runner.WithProtoset(ProtoFile),
		runner.WithData(empty),
		runner.WithInsecure(client.config.Insecure),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	return
}
