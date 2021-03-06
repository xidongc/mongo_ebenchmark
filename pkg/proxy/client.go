/*
 * mongodb_ebenchmark - Mongodb grpc proxy benchmark for e-commerce workload (still in dev)
 * Copyright (c) 2020 - Chen, Xidong <chenxidong2009@hotmail.com>
 *
 * All rights reserved.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 */

package proxy

import (
	"context"
	"fmt"
	"github.com/bojand/ghz/printer"
	"github.com/bojand/ghz/runner"
	log "github.com/sirupsen/logrus"
	"github.com/xidongc-wish/mgo"
	"github.com/xidongc-wish/mgo/bson"
	"github.com/xidongc/mongo_ebenchmark/mprpc"
	"github.com/xidongc/mongo_ebenchmark/pkg/cfg"
	"google.golang.org/grpc"
	"os"
	"sync"
	"sync/atomic"
	"time"
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
	Database   = "ebenchmark"
	Collection = "default"
)

// Mode for FindAndModify
const (
	FindAndDelete FindAndModifyMode = 0
	FindAndUpdate FindAndModifyMode = 1
	FindAndUpsert FindAndModifyMode = 2
)

// Client represents a middleware with the actual database driver
//
// Currently client will use ebenchmark database as defined in const
// to avoid interfere with production workload. client support multi
// api. See the documentation on const for more details.
type Client struct {
	config      *cfg.ProxyConfig
	Host        string
	Collection  *mprpc.Collection
	Turbo       bool
	activeCon   *grpc.ClientConn
	rpcClient   mprpc.MongoProxyClient
	cancelFunc  context.CancelFunc
	Healthy     int32
	ProtoFile   string
	amplifierWG sync.WaitGroup // add for lock prep work
}

// NewClient Creates a new proxy client based on provided cfg
// This method is generally called just once for con establish,
// does health check in the background task, and check for env:
// PROTOSET_FILE, this env represent grpc interface with driver
//
// Once Client is not useful anymore, Close must be called to
// release the resources appropriately
func NewClient(config *cfg.ProxyConfig, namespace string, cancel context.CancelFunc) (client *Client, err error) {
	if config == nil {
		config = cfg.DefaultConfig()
	}
	if namespace == "" {
		namespace = Collection
	}
	if config.ProxyPort < 0 || config.ProxyPort == 0 {
		log.Warning("port not defined properly, set back to default")
		config.ProxyPort = 50051
	}

	host := fmt.Sprintf("%s:%d", config.ProxyAddr, config.ProxyPort)
	client = &Client{
		config: config,
		Host:   host,
	}

	val, ok := os.LookupEnv("PROTOSET_FILE")
	if !ok {
		atomic.StoreInt32(&client.Healthy, 0)
		log.Panic("not found env PROTOSET_FILE")
		return
	}
	client.amplifierWG = sync.WaitGroup{}
	client.ProtoFile = val

	if client.Collection == nil {
		client.Collection = &mprpc.Collection{
			Database:   Database,
			Collection: namespace,
		}
	}
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect to rpc cfg error: %s", err)
	}
	client.activeCon = conn
	client.rpcClient = mprpc.NewMongoProxyClient(conn)
	if cancel == nil {
		log.Warning("handle context exit in application")
	}
	client.cancelFunc = cancel

	go func() {
		for {
			_ = client.HealthCheck()
			if atomic.LoadInt32(&client.Healthy) != 1 {
				panic("cfg not healthy error")
			}
			time.Sleep(time.Second * 300)
		}
	}()
	return
}

// Close will release the resources in a proxy client, and call
// cancelFunc if specified
//
// Call Close after NewClient, See NewClient for more details
func (client *Client) Close() (err error) {
	if err := client.activeCon.Close(); err != nil {
		log.Fatalf("clean up failed: %s", err)
	}
	if client.cancelFunc != nil {
		client.cancelFunc()
	}
	return
}

// Find prepares a query using the provided document
//
// See proxy.QueryParam for customizing query param
//
// Relevant documentation:
//
//     http://www.mongodb.org/display/DOCS/Querying
//     http://www.mongodb.org/display/DOCS/Advanced+Queries
//
func (client *Client) Find(ctx context.Context, query *QueryParam) (docs []bson.M, err error) {
	filterBytes, err := bson.Marshal(query.Filter)
	if err != nil {
		log.Errorf("%s: marshal filter error", Find)
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

	// TODO add params
	request := mprpc.FindQuery{
		Collection:  client.Collection,
		Filter:      filterBytes,
		Skip:        query.Skip,
		Maxtimems:   -1,
		Maxscan:     0,
		Prefetch:    prefetch,
		Batchsize:   client.config.BatchSize,
		Readpref:    int32(readPref),
		Findone:     query.FindOne,
		Partial:     client.config.AllowPartial,
		Readconcern: readConcern,
		Comment:     FindIter,
		Rpctimeout:  client.config.RpcTimeout,
		Hint:        query.UsingIndex,
	}

	if query.Amp != nil {
		report, err := runner.Run(
			Find,
			client.Host,
			runner.WithProtoset(client.ProtoFile),
			runner.WithConcurrency(query.Amp.Concurrency),
			runner.WithConnections(query.Amp.Connections),
			runner.WithCPUs(query.Amp.CPUs),
			runner.WithData(request),
			runner.WithInsecure(true),
		)
		if err != nil {
			log.Fatal(err.Error())
		}
		file, err := os.Create("results/test_find.html")
		p := printer.ReportPrinter{
			Out:    file,
			Report: report,
		}

		_ = p.Print("html")
	}

	resultSet, err := client.rpcClient.Find(ctx, &request)
	if err != nil {
		log.Error(err)
	}
	for _, r := range resultSet.Results {
		var doc bson.M
		if err = bson.Unmarshal(r.Val, &doc); err != nil {
			log.Error(err)
		}
		docs = append(docs, doc)
	}
	return
}

// FindIter works like Find, but uses iterClient as stream function
//
// See proxy.QueryParam for customizing query param
//
// Relevant documentation:
//
//     http://www.mongodb.org/display/DOCS/Querying
//     http://www.mongodb.org/display/DOCS/Advanced+Queries
//
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
		Maxscan:     0,
		Prefetch:    prefetch,
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
		runner.WithProtoset(client.ProtoFile),
		runner.WithConcurrency(query.Amp.Concurrency),
		runner.WithConnections(query.Amp.Connections),
		runner.WithCPUs(query.Amp.CPUs),
		runner.WithData(request),
		runner.WithInsecure(true),
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

// Count returns the total number of documents in the collection.
//
// See proxy.QueryParam for customizing query param
func (client *Client) Count(ctx context.Context, query *QueryParam) (count uint64, err error) {
	return
}

// Explain returns a number of details about how the MongoDB cfg would
// execute the requested query, such as the number of objects examined,
// the number of times the read lock was yielded to allow writes to go in,
// and so on.
//
// See proxy.QueryParam for customizing query param
//
// Relevant documentation:
//
//     http://www.mongodb.org/display/DOCS/Optimization
//     http://www.mongodb.org/display/DOCS/Query+Optimizer
//
func (client *Client) Explain(ctx context.Context, query *QueryParam) (explainFields bson.M, err error) {
	return
}

// Aggregate processes data records and return computed results.
// Aggregation operations group values from multiple documents together,
// and can perform a variety of operations on the grouped data to return
// a single result.
//
// See proxy.AggregateParam for customizing aggregate param
//
// Relevant documentation:
//
//     http://docs.mongodb.org/manual/reference/aggregation
//     http://docs.mongodb.org/manual/applications/aggregation
//     http://docs.mongodb.org/manual/tutorial/aggregation-examples
//
func (client *Client) Aggregate(ctx context.Context, query *AggregateParam) (documents []interface{}, err error) {
	return
}

// Update finds a single document matching the provided selector document
// and modifies it according to the update document.
//
// See proxy.QueryParam for customizing query param
//
// Relevant documentation:
//
//     http://www.mongodb.org/display/DOCS/Updating
//     http://www.mongodb.org/display/DOCS/Atomic+Operations
//
func (client *Client) Update(ctx context.Context, param *UpdateParam) (changeInfo *mprpc.ChangeInfo, err error) {
	var wOptions *mprpc.WriteOptions
	if client.Turbo {
		wOptions = cfg.GetTurboWriteOptions()
	} else {
		wOptions = cfg.GetSafeWriteOptions()
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

// Remove finds documents matching the provided selector document
// and removes it from the database.
//
// See proxy.RemoveParam for customizing remove param
//
// Relevant documentation:
//
//     http://www.mongodb.org/display/DOCS/Removing
//
func (client *Client) Remove(ctx context.Context, param *RemoveParam) (changeInfo *mprpc.ChangeInfo, err error) {
	b, err := bson.Marshal(param.Filter)
	if err != nil {
		log.Error(err)
		return
	}
	var wOptions *mprpc.WriteOptions
	if client.Turbo {
		wOptions = cfg.GetTurboWriteOptions()
	} else {
		wOptions = cfg.GetSafeWriteOptions()
	}
	removeOps := &mprpc.RemoveOperation{
		Collection:   client.Collection,
		Filter:       b,
		Writeoptions: wOptions,
	}
	log.Info(removeOps)
	if changeInfo, err = client.rpcClient.Remove(ctx, removeOps); err != nil {
		log.Error(err)
	}
	return
}

// Insert inserts one or more documents in the respective collection.
//
// See proxy.InsertParam for customizing insert param
//
// Relevant documentation:
//
// http://www.mongodb.org/display/DOCS/Inserting
//
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
		wOptions = cfg.GetTurboWriteOptions()
	} else {
		wOptions = cfg.GetSafeWriteOptions()
	}
	request := mprpc.InsertOperation{
		Collection:   client.Collection,
		Documents:    rpcDocs,
		Writeoptions: wOptions,
	}

	if param.Amp != nil {
		client.amplifierWG.Add(1)
		report, err := runner.Run(
			Insert,
			client.Host,
			runner.WithProtoset(client.ProtoFile),
			runner.WithConcurrency(param.Amp.Concurrency),
			runner.WithConnections(param.Amp.Connections),
			runner.WithCPUs(param.Amp.CPUs),
			runner.WithData(request),
			runner.WithInsecure(!client.config.Secure),
		)
		if err != nil {
			log.Fatal(err.Error())
		}
		p := printer.ReportPrinter{
			Out:    os.Stdout,
			Report: report,
		}

		_ = p.Print("pretty")

		removeParams := UndoInsert(param)
		for _, p := range removeParams {
			if _, err := client.Remove(ctx, p); err != nil {
				log.Error(err)
			}
		}
		client.amplifierWG.Done()

	} else {
		log.Info("no amp specified")
	}

	if _, err = client.rpcClient.Insert(ctx, &request); err != nil {
		log.Errorf("rpc insert error with: %s", err)
		return
	}
	return
}

// FindAndModify allows updating, upserting or removing a document matching
// a query and atomically returning either the old version (the default) or
// the new version of the document (when ReturnNew is true). If no objects
// are found Apply returns ErrNotFound.
//
// See proxy.FindModifyParam for customizing findAndModify param
//
// Relevant documentation:
//
//     http://www.mongodb.org/display/DOCS/findAndModify+Command
//     http://www.mongodb.org/display/DOCS/Updating
//     http://www.mongodb.org/display/DOCS/Atomic+Operations
//
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
		wOptions = cfg.GetTurboWriteOptions()
	} else {
		wOptions = cfg.GetSafeWriteOptions()
	}

	// TODO add params back
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

// Distinct unmarshals into result the list of distinct values for the given key.
//
// See proxy.QueryParam for customizing distinct param
//
// Relevant documentation:
//
//     http://www.mongodb.org/display/DOCS/Aggregation
//
func (client *Client) Distinct(ctx context.Context, query *QueryParam) (distinctKeys []interface{}, err error) {
	return
}

// HealthCheck checks database driver, and atomically update client
func (client *Client) HealthCheck() (err error) {
	log.Info("Start doing health check")
	empty := Empty{}
	_, err = runner.Run(
		Healthcheck,
		client.Host,
		runner.WithProtoset(client.ProtoFile),
		runner.WithData(empty),
		runner.WithInsecure(true),
	)
	if err != nil {
		atomic.StoreInt32(&client.Healthy, 0)
		log.Fatal(err.Error())
	} else {
		atomic.StoreInt32(&client.Healthy, 1)
	}
	return
}
