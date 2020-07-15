protoc -I include/googleapis -I model -I model/sku/skupb --go_out==plugins=grpc:$(go env GOPATH)/src model/sku/skupb/sku.proto
protoc -I include/googleapis -I model/payment --go_out==plugins=grpc:$(go env GOPATH)/src model/payment/payment.proto
protoc -I include/googleapis -I model -I model/product/productpb --go_out==plugins=grpc:$(go env GOPATH)/src model/product/productpb/product.proto
protoc -I include/googleapis -I model -I model/user/userpb --go_out==plugins=grpc:$(go env GOPATH)/src model/user/userpb/user.proto

