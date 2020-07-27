default: \
	vendor \
	build/ebenchmark.linux \
	build/ebenchmark.darwin

vendor: go.mod go.sum
	go mod tidy
	go mod vendor

build/ebenchmark.linux:
	@echo "$@"
	@GOOS=linux CGO_ENABLED=0 go build -o bin/ebenchmark.linux \
		github.com/xidongc/mongodb_ebenchmark/cmd

build/ebenchmark.darwin:
	@echo "$@"
	@GOOS=darwin CGO_ENABLED=0 go build -o bin/ebenchmark.darwin \
    	github.com/xidongc/mongodb_ebenchmark/cmd

pb:
	protoc -I include/googleapis -I model -I model/sku/skupb --go_out=plugins=grpc:$(go env GOPATH)/src model/sku/skupb/sku.proto
	protoc -I include/googleapis -I model/payment --go_out=plugins=grpc:$(go env GOPATH)/src model/payment/paymentpb/payment.proto
	protoc -I include/googleapis -I model -I model/product/productpb --go_out=plugins=grpc:$(go env GOPATH)/src model/product/productpb/product.proto
	protoc -I include/googleapis -I model -I model/user/userpb --go_out=plugins=grpc:$(go env GOPATH)/src model/user/userpb/user.proto
	protoc -I include/googleapis -I model -I model/order/orderpb --go_out=plugins=grpc:$(go env GOPATH)/src model/order/orderpb/order.proto

sku:
	ghz --insecure --protoset ./model/sku/sku.protoset --call skupb.SkuService.New -d '{"parent":"parent","image":"wertw","price":123,"active":false,"name":"xidongc"}' -c 1 -n 1 0.0.0.0:50053

db.insert:
	ghz --insecure --protoset ./pkg/proxy/rpc.protoset --call mprpc.MongoProxy.Insert -d '[{"documents":[{"val":"WQAAAAdfaWQAXw\/F2PfpmwABuq0PAnN0YXRlAAcAAABhY3RpdmUAAm5hbWUACQAAAHhpZG9uZ2MzAAJtc2dzAAgAAABzdWNjZXNzABBudW1iZXIAAwAAAAA="}],"writeoptions":{"rpctimeout":30000,"writetimeout":null,"j":null,"fsync":null,"writeconcern":1},"collection":{"collection":"mpc","database":"mpc"}}]' -c 1 -n 1 0.0.0.0:50051

db.finditer:
	ghz --insecure --protoset ./pkg/proxy/rpc.protoset --call mprpc.MongoProxy.FindIter -d '[{"comment":"[172.17.0.2]delete @ crud\/main.py:133","filter":"IQAAAANzdGF0ZQAVAAAAAiRlcQAHAAAAYWN0aXZlAAAA","readpref":5,"findone":false,"distinctkey":null,"skip":0,"batchsize":10000,"collection":{"collection":"mpc","database":"mpc"},"rpctimeout":25000,"limit":0,"maxscan":null,"maxtimems":-1,"partial":false,"prefetch":null}]' -c 5 -n 20 0.0.0.0:50051

db.healthcheck:
	ghz --insecure --protoset ./pkg/proxy/rpc.protoset --call mprpc.MongoProxy.Healthcheck -d {} -c 5 -n 20 0.0.0.0:50051