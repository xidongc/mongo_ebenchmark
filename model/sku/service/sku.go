package sku

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/xidongc-wish/mgo/bson"
	"github.com/xidongc/mongodb_ebenchmark/model/sku/skupb"
	"github.com/xidongc/mongodb_ebenchmark/pkg/proxy"
)

const ns = "sku"

type Service struct {
	Storage proxy.Client
	Amplifier proxy.Amplifier
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
		Amp:    s.Amplifier,
	}

	changeInfo, err := s.Storage.Update(ctx, param)
	if err != nil {
		log.Error("error")
	}

	log.Info(changeInfo)


	if err != nil {
		log.Errorf("sku error: storage failed with %s", err)
	}

	sku = &skupb.Sku{
		Name:			     req.GetName(),
		Price:               req.GetPrice(),
		Currency:            req.GetCurrency(),
		Active:              req.GetActive(),
		ProductId:		     req.GetParent(),
		Image:               req.GetImage(),
		Metadata:            req.GetMetadata(),
		PackageDimensions:   req.GetPackageDimensions(),
		Attributes: 	     req.GetAttributes(),
	}

	return
}

func NewSKUService(config *proxy.Config, amplifier proxy.Amplifier, cancel context.CancelFunc) *Service{
	client, _ := proxy.NewClient(config, ns, cancel)
	return &Service {
		Storage: *client,
		Amplifier: amplifier,
	}
}
