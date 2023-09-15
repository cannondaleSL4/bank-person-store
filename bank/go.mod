module github.com/cannondaleSL4/bank

go 1.20

require (
	github.com/cannondaleSL4/bank-person-store/common v0.0.1
	github.com/lib/pq v1.10.9
	github.com/shopspring/decimal v1.3.1
	go.uber.org/zap v1.25.0
	google.golang.org/grpc v1.57.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230525234030-28d5490b6b19 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

replace github.com/cannondaleSL4/bank-person-store/common => /home/dima/go/src/github.com/cannondaleSL4/bank-person-store/common
