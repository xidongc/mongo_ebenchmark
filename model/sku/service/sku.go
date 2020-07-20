package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/xidongc-wish/mgo/bson"
	"github.com/xidongc/mongodb_ebenchmark/model/sku/skupb"
	"github.com/xidongc/mongodb_ebenchmark/pkg/proxy"
	"time"
)

type skuService struct {
	storage proxy.Client
	amplifier proxy.Amplifier
}

func (s *skuService) New(req *skupb.NewRequest) (sku *skupb.Sku, err error){
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var documents []interface{}

	stream, err := s.storage.FindIter(ctx, bson.M{"name": req.GetName()}, s.amplifier)
	if err != nil {
		log.Error("error")
	}
	for {
		if stream == nil {
			break
		}
		rs, err := stream.Recv()
		if err != nil {
			break
		}
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
	log.Info(documents)


	if err != nil {
		log.Errorf("sku error: storage failed with %s", err)
	}

	sku = &skupb.Sku{
		Name:			   req.GetName(),
		Currency:          req.GetCurrency(),
		Active:            req.GetActive(),
		Price:             req.GetPrice(),
		ProductId:		   req.GetParent(),
		Image:             req.GetImage(),
		Metadata:          req.GetMetadata(),
		PackageDimensions: req.GetPackageDimensions(),
		Attributes: 	   req.GetAttributes(),
	}
	return
}
