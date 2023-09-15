package main

import (
	"fmt"
	"github.com/cannondaleSL4/bank-person-store/common/config"
	"github.com/cannondaleSL4/bank-person-store/common/logger"
	"github.com/cannondaleSL4/bank-person-store/common/proto"
	"github.com/cannondaleSL4/bank/db"
	"github.com/cannondaleSL4/bank/service"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"net"
)

func main() {

	config.Init("config.yaml")
	conf := config.GetConfig()

	logger.Init(conf.Server.ServiceName)
	defer logger.Sync()

	log := logger.GetLogger()
	sugar := log.Sugar()

	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Database.Host, conf.Database.Port, conf.Database.User, conf.Database.Password, conf.Database.DbName)

	store := db.New(connection, sugar)
	defer store.Close()

	grpcServer := grpc.NewServer()

	bankService := service.New(store, sugar)

	proto.RegisterBankServiceServer(grpcServer, bankService)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", conf.Server.Port))
	if err != nil {
		sugar.Fatalf("Failed to listen on port %s: %v", conf.Server.Port, err)
	}
	sugar.Infoln(fmt.Sprintf("gRPC server started on port %s", conf.Server.Port))
	if err := grpcServer.Serve(lis); err != nil {
		sugar.Fatalf("Failed to serve gRPC server: %v", err)
	}

}
