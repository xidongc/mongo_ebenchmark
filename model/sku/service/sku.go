package sku

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/xidongc-wish/mgo/bson"
	"github.com/xidongc/mongodb_ebenchmark/model/sku/skupb"
	"github.com/xidongc/mongodb_ebenchmark/pkg/proxy"
)

type Service struct {
	storage proxy.Client
	amplifier proxy.Amplifier
}

func (s *Service) Get(context.Context, *skupb.GetRequest) (*skupb.Sku, error) {
	panic("implement me")
}

func (s *Service) Update(context.Context, *skupb.UpdateRequest) (*skupb.Sku, error) {
	panic("implement me")
}

func (s *Service) Delete(context.Context, *skupb.DeleteRequest) (*skupb.Empty, error) {
	panic("implement me")
}

func (s *Service) New(ctx context.Context, req *skupb.NewRequest) (sku *skupb.Sku, err error){

	updateQuery := bson.M{
		"name":			     req.GetName(),
		"price":             req.GetPrice(),
		"currency":          req.GetCurrency(),
		"active":            req.GetActive(),
		"productid":		 req.GetParent(),
		"image":             req.GetImage(),
		"metadata":          req.GetMetadata(),
		"packagedimensions": req.GetPackageDimensions(),
		"attibutes": 	     req.GetAttributes(),
	}

	param := &proxy.UpdateParam{
		Filter: bson.M{"name": req.GetName()},
		Update: updateQuery,
		Upsert: true,
		Multi:  false,
		Amp:    s.amplifier,
	}

	changeInfo, err := s.storage.Update(ctx, param)
	if err != nil {
		log.Error("error")
	}

	log.Info(changeInfo)


	if err != nil {
		log.Errorf("sku error: storage failed with %s", err)
	}

	return
}

func RegisterService() {

}
