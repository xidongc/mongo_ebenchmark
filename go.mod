module github.com/xidongc/mongodb_ebenchmark

go 1.14

require (
	github.com/bojand/ghz v0.55.0
	github.com/golang/protobuf v1.4.2
	github.com/sirupsen/logrus v1.2.0
	github.com/xidongc-wish/mgo v0.0.0-20200417061821-13161a071d79
	github.com/xidongc-wish/mongoproxy/mprpc v0.0.0-00010101000000-000000000000
	golang.org/x/sys v0.0.0-20200625212154-ddb9806d33ae // indirect
	google.golang.org/genproto v0.0.0-20200707001353-8e8330bf89df
	google.golang.org/protobuf v1.25.0
)

replace github.com/xidongc-wish/mongoproxy/mprpc => ../../xidongc-wish/mongoproxy/mprpc
