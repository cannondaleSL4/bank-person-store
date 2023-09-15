# Generate proto files
1. `cd ./proto` - change dir
2. go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
2. protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative bank.proto
