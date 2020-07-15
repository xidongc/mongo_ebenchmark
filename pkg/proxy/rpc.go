package proxy

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/xidongc-wish/mgo/bson"
	"github.com/xidongc-wish/mongoproxy/mprpc"
	"google.golang.org/grpc"
	"strconv"
	"time"
)

type MPC interface {
	Connect() error
	Close() error
	IsHealthy(context.Context) bool
	Find(context.Context, *mprpc.FindQuery)([]interface{}, error)
	FindIter(context.Context, *mprpc.FindQuery) ([]interface{}, error)
	Count(context.Context, *mprpc.FindQuery) (*mprpc.ResultSetCount, error)
	Explain(context.Context, *mprpc.FindQuery) (bson.M, error)
	Aggregate(context.Context, *mprpc.AggregateQuery) ([]interface{}, error)
	Bulk(context.Context, *mprpc.BulkOperation) error
	Update(context.Context, *mprpc.UpdateOperation) (*mprpc.ChangeInfo, error)
	Remove(context.Context, *mprpc.RemoveOperation) (*mprpc.ChangeInfo, error)
	Insert(context.Context, *mprpc.InsertOperation) error
	FindAndModify(context.Context, *mprpc.FindAndModifyOperation) (interface{}, error)
	Distinct(context.Context, *mprpc.FindQuery, interface{}) error
}

func (c *Client) Connect() error {
	serverAddr := c.Config.ServerIp + ":" + strconv.Itoa(c.Config.Port)
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error: %+v", err)
		return err
	}
	c.activeConnection = conn
	c.MongoProxyClient = mprpc.NewMongoProxyClient(conn)
	return nil
}

func (c *Client) Close() error {
	if err := c.activeConnection.Close(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (c *Client) IsHealthy(ctx context.Context) (isHealth bool) {
	log.Debug("record checkHealth api call")
	hcRequest := &mprpc.HealthcheckRequest{}
	resp, err := c.MongoProxyClient.Healthcheck(ctx, hcRequest)
	if err != nil {
		log.Error(err)
	}
	if resp.Response == "OK" {
		isHealth = true
	}
	return
}

func (c *Client) Find(ctx context.Context, query *mprpc.FindQuery) (documents []interface{}, err error){
	log.Debug("record Find api call")
	rs, err := c.MongoProxyClient.Find(ctx, query)
	if err != nil {
		log.Fatalf("Unable to get initiate Find: %v", err)
	}
	log.Info(rs)
	return
}

func (c *Client) FindIter(ctx context.Context, query *mprpc.FindQuery) (documents []interface{}, err error) {
	log.Debug("record FindIter api call")
	stream, err := c.MongoProxyClient.FindIter(ctx, query)
	if err != nil {
		log.Fatalf("Unable to get initiate FindIter: %v", err)
	}
	for {
		rs, err := stream.Recv()
		if err != nil {
			break
		}
		// to avoid race condition, refer: https://medium.com/@cep21/gos-append-is-not-always-thread-safe-a3034db7975
		tmpDocs := make(chan interface{}, len(rs.Results))

		if len(rs.Results) > 0 {
			for _, document := range rs.Results {
				var doc bson.D
				if err := bson.Unmarshal(document.Val, &doc); err != nil {
					log.Fatal(err)
				}
				tmpDocs <- doc
			}

			for i := 0; i < len(rs.Results); i++ {
				doc := <-tmpDocs
				documents = append(documents, doc)
			}
		} else {
			time.Sleep(5 * time.Millisecond)
		}
	}
	return
}

func (c *Client) Count(ctx context.Context, query *mprpc.FindQuery) (results *mprpc.ResultSetCount, err error) {
	log.Debug("record count api call")
	results, err = c.MongoProxyClient.Count(ctx, query)
	if err != nil {
		log.Error(err)
	}
	return
}

func (c *Client) Explain(ctx context.Context, query *mprpc.FindQuery) (explainFields bson.M, err error) {
	log.Info("record explain api call")
	result, err := c.MongoProxyClient.Explain(ctx, query)
	if err != nil {
		log.Fatalf("Error doing the explain: %v", err)
	}

	if err := bson.Unmarshal(result.Val, &explainFields); err != nil {
		log.Fatalf("Error un-marshalling bson bytes to map: %v", err)
	}

	expectedFields := []string{"executionStats", "queryPlanner"}
	for _, field := range expectedFields {
		_, ok := explainFields[field]
		if !ok {
			log.Errorf("Field %s not found in explain result.", field)
		}
	}
	log.Info(explainFields["executionStats"])
	log.Info(explainFields["queryPlanner"])
	return
}

func (c *Client) Aggregate(ctx context.Context, query *mprpc.AggregateQuery) (documents []interface{}, err error) {
	log.Info("record aggregate api call ")
	resp, err := c.MongoProxyClient.Aggregate(ctx, query)
	if resp != nil {
		rs, _ := resp.Recv()

		for _, document := range rs.Results {
			var doc bson.D
			if err := bson.Unmarshal(document.Val, &doc); err != nil {
				log.Fatal(err)
			}
			documents = append(documents, doc)
		}
	}

	if err != nil {
		log.Error(err)
	}
	return
}

func (c *Client) Bulk(ctx context.Context, steps *mprpc.BulkOperation) (err error) {
	log.Info("record bulk api call")
	if _, err := c.MongoProxyClient.Bulk(ctx, steps); err != nil {
		log.Error(err)
	}
	return
}

func (c *Client) Update(ctx context.Context, steps *mprpc.UpdateOperation) (changeInfo *mprpc.ChangeInfo, err error) {
	log.Debug("record update api call")
	changeInfo, err = c.MongoProxyClient.Update(ctx, steps)
	if err != nil {
		log.Error(err)
	}
	return
}

func (c *Client) Remove(ctx context.Context, steps *mprpc.RemoveOperation) (changeInfo *mprpc.ChangeInfo, err error) {
	log.Debug("record remove api call ")
	changeInfo, err = c.MongoProxyClient.Remove(ctx, steps)
	if err != nil {
		log.Error(err)
	}
	return
}

func (c *Client) Insert(ctx context.Context, steps *mprpc.InsertOperation) (err error) {
	log.Debug("record Insert api call ")
	if resp, err := c.MongoProxyClient.Insert(ctx, steps); err != nil {
		log.Fatal("insert failed ", err, resp)
	}
	return
}

func (c *Client) FindAndModify(ctx context.Context, query *mprpc.FindAndModifyOperation, doc interface{}) (err error) {
	log.Debug("record find and modify api call ")
	resp, err := c.MongoProxyClient.FindAndModify(ctx, query)
	if err != nil {
		log.Error(err)
	}
	log.Info(resp)
	return
}

func (c *Client) Distinct(ctx context.Context, query *mprpc.FindQuery, results interface{}) (err error) {
	log.Debug("log distinct api call")
	resp, err := c.MongoProxyClient.Distinct(ctx, query)
	if err != nil {
		log.Error(err)
	}

	var doc = bson.Raw{
		Kind: byte(4),
		Data: resp.Results,
	}

	err = doc.Unmarshal(results)
	if err != nil {
		log.Error(err)
	}

	return
}
