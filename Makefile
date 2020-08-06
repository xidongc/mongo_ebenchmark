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
		github.com/xidongc/mongo_ebenchmark/cmd

build/ebenchmark.darwin:
	@echo "$@"
	@GOOS=darwin CGO_ENABLED=0 go build -o bin/ebenchmark.darwin \
    	github.com/xidongc/mongo_ebenchmark/cmd

pb.payment:
	protoc -I include/googleapis -I model/payment --go_out=plugins=grpc:$(go env GOPATH)/src model/payment/paymentpb/payment.proto

pb.order:
	protoc -I include/googleapis -I model -I model/order/orderpb --go_out=plugins=grpc:$(go env GOPATH)/src model/order/orderpb/order.proto

pb.sku:
	protoc -I include/googleapis -I model -I model/sku/skupb --go_out=plugins=grpc:$(go env GOPATH)/src model/sku/skupb/sku.proto
	protoc --include_imports -I ./include/googleapis -I model -I model/sku/skupb --descriptor_set_out=./model/sku/sku.protoset ./model/sku/skupb/sku.proto

pb.user:
	protoc -I include/googleapis -I model -I model/user/userpb --go_out=plugins=grpc:$(go env GOPATH)/src model/user/userpb/user.proto
	protoc --include_imports -I ./include/googleapis -I model -I model/user/userpb --descriptor_set_out=./model/user/user.protoset ./model/user/userpb/user.proto

pb.product:
	protoc -I include/googleapis -I model -I model/product/productpb --go_out=plugins=grpc:$(go env GOPATH)/src model/product/productpb/product.proto
	protoc --include_imports -I ./include/googleapis -I model -I model/product/productpb --descriptor_set_out=./model/product/product.protoset ./model/product/productpb/product.proto

sku.new:
	ghz --insecure --protoset ./model/sku/sku.protoset --call skupb.SkuService.New -d '{"productId":"1234567","image":"wertw","price":123,"active":false,"name":"xidong", "inventory": {"skuId": 123, "warehouseId": 12345}, "packageDimensions": {"height": 10, "length": 10, "weight": 10.3, "width":10.23}, "hasLiquid": false, "hasBattery": false, "hasSensitive": false, "description":"this is only a test"}' -c 1 -n 1 0.0.0.0:50053

sku.get:
	ghz --insecure --protoset ./model/sku/sku.protoset --call skupb.SkuService.Get -d '{"name": "xidong"}' -c 1 -n 1 0.0.0.0:50053

sku.delete:
	ghz --insecure --protoset ./model/sku/sku.protoset --call skupb.SkuService.Delete -d '{"name": "xidong"}' -c 1 -n 1 0.0.0.0:50053

product.new:
	ghz --insecure --protoset ./model/product/product.protoset --call productpb.ProductService.New -d '{"id":"1234567", "name": "xidong", "active": false, "attributes": ["hello", "world"], "description": "hello world", "images": ["test", "test1"], "metadata": {"name": "xidongc", "product": "xidongc"}, "shippable": false, "url": "www.google.com"}' -c 1 -n 1 0.0.0.0:50053

product.update:
	ghz --insecure --protoset ./model/product/product.protoset --call productpb.ProductService.Update -d '{"id":"1234567", "name": "xidong", "active": false, "attributes": ["hello", "world"], "description": "hello world", "images": ["test", "test1"], "metadata": {"name": "xidongc", "product": "xidongc"}, "shippable": false, "url": "www.google.com"}' -c 1 -n 1 0.0.0.0:50053

product.get:
	ghz --insecure --protoset ./model/product/product.protoset --call productpb.ProductService.Get -d '{"id":"1234567"}' -c 1 -n 1 0.0.0.0:50053

product.delete:
	ghz --insecure --protoset ./model/product/product.protoset --call productpb.ProductService.Delete -d '{"id":"1234567"}' -c 1 -n 1 0.0.0.0:50053

user.new:
	ghz --insecure --protoset ./model/user/user.protoset --call userpb.UserService.New -d '{"name":"xidongc", "nickname": "xidongc", "email": "chenxidong2009@hotmail.com", "active": true, "balance": 10, "currency": 10, "image": "www.google.com", "pwd": "pwd", "metadata": {"name": "xidongc", "sex": "male"}}' -c 1 -n 1 0.0.0.0:50053

user.get:
	ghz --insecure --protoset ./model/user/user.protoset --call userpb.UserService.Get -d '{"nickname": "xidongc"}' -c 1 -n 1 0.0.0.0:50053

db.insert:
	ghz --insecure --protoset ./pkg/proxy/rpc.protoset --call mprpc.MongoProxy.Insert -d '[{"documents":[{"val":"WQAAAAdfaWQAXw\/F2PfpmwABuq0PAnN0YXRlAAcAAABhY3RpdmUAAm5hbWUACQAAAHhpZG9uZ2MzAAJtc2dzAAgAAABzdWNjZXNzABBudW1iZXIAAwAAAAA="}],"writeoptions":{"rpctimeout":30000,"writetimeout":null,"j":null,"fsync":null,"writeconcern":1},"collection":{"collection":"mpc","database":"mpc"}}]' -c 1 -n 1 0.0.0.0:50051

db.finditer:
	ghz --insecure --protoset ./pkg/proxy/rpc.protoset --call mprpc.MongoProxy.FindIter -d '[{"comment":"[172.17.0.2]delete @ crud\/main.py:133","filter":"IQAAAANzdGF0ZQAVAAAAAiRlcQAHAAAAYWN0aXZlAAAA","readpref":5,"findone":false,"distinctkey":null,"skip":0,"batchsize":10000,"collection":{"collection":"mpc","database":"mpc"},"rpctimeout":25000,"limit":0,"maxscan":null,"maxtimems":-1,"partial":false,"prefetch":null}]' -c 5 -n 20 0.0.0.0:50051

db.healthcheck:
	ghz --insecure --protoset ./pkg/proxy/rpc.protoset --call mprpc.MongoProxy.Healthcheck -d {} -c 5 -n 20 0.0.0.0:50051
