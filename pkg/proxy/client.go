package proxy

import (
	"fmt"
	"github.com/bojand/ghz/printer"
	"github.com/bojand/ghz/runner"
	log "github.com/sirupsen/logrus"
	"github.com/xidongc-wish/mgo/bson"
	"github.com/xidongc-wish/mongoproxy/mprpc"
	"os"
)

const (
	Find          = "mprpc.MongoProxy.Find"
	FindIter      = "mprpc.MongoProxy.FindIter"
	Count         = "Count"
	Explain       = "Explain"
	Aggregate     = "Aggregate"
	Bulk          = "Bulk"
	Update        = "Update"
	Remove        = "Remove"
	Insert        = "mprpc.MongoProxy.Insert"
	FindAndModify = "FindAndModify"
	Distinct      = "Distinct"
	Healthcheck   = "mprpc.MongoProxy.Healthcheck"
)

const (
	Database 	  = "mpc"
	Collection 	  = "mpc"
)

// TODO define as const, should be in env
const (
	ProtoFile     = "/Users/derekchen/go/src/github.com/xidongc/mongodb_ebenchmark/pkg/proxy/rpc.protoset"
)

type Client struct {
	config *Config
	Host string
	Collection *mprpc.Collection
	Turbo bool
}

type Empty struct {}

type Documents []byte

func NewClient(config *Config) (client *Client, err error) {
	if config == nil {
		config = DefaultConfig()
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
	client.Collection = &mprpc.Collection{
		Database: Database,
		Collection: Collection,
	}
	return
}

func (client *Client) FindIter(filter bson.M, amp Amplifier) (err error) {
	stream, err := bson.Marshal(filter)
	if err != nil {
		log.Errorf("%s: marshall filter error", "FindIter")
	}
	request := mprpc.FindQuery{
		Collection:  client.Collection,
		Filter:      stream,
		Skip:        0,
		Maxtimems:   -1,
		Batchsize:   client.config.BatchSize,
		Readpref:    client.config.ReadPref,
		Findone:     false,
		Partial:     false,
		Comment:     FindIter,
		Rpctimeout:  client.config.RpcTimeout,
	}
	report, err := runner.Run(
		FindIter,
		client.Host,
		runner.WithProtoset(ProtoFile),
		runner.WithConcurrency(amp().Concurrency),
		runner.WithConnections(amp().Connections),
		runner.WithCPUs(amp().CPUs),
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

// TODO detect dup obj id, add amplify with different obj id
func (client *Client) Insert(docs []interface{}, amp Amplifier) (err error) {

	var rpcDocs []*mprpc.Document
	for _, doc := range docs {
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
	report, err := runner.Run(
		Insert,
		client.Host,
		runner.WithProtoset(ProtoFile),
		runner.WithConcurrency(amp().Concurrency),
		runner.WithConnections(amp().Connections),
		runner.WithCPUs(amp().CPUs),
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
