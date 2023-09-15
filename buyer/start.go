package main

import (
	"fmt"
	bankConfig "github.com/cannondaleSL4/bank-person-store/buyer/config"
	"github.com/cannondaleSL4/bank-person-store/buyer/db"
	"github.com/cannondaleSL4/bank-person-store/buyer/service"
	"github.com/cannondaleSL4/bank-person-store/common/config"
	"github.com/cannondaleSL4/bank-person-store/common/logger"
	"github.com/cannondaleSL4/bank-person-store/common/proto"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"time"
)

func main() {

	config.Init("config.yaml")
	conf := config.GetConfig()

	content, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Errorf("Error reading config file: %v", err)
	}

	var bankConf bankConfig.BankServiceConfig
	err = yaml.Unmarshal(content, &bankConf)
	if err != nil {
		fmt.Errorf("Error unmarshalling config: %v", err)
	}

	logger.Init(conf.Server.ServiceName)
	defer logger.Sync()

	log := logger.GetLogger()
	sugar := log.Sugar()

	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Database.Host, conf.Database.Port, conf.Database.User, conf.Database.Password, conf.Database.DbName)

	store := db.New(connection, sugar)
	defer store.Close()

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", bankConf.Host, bankConf.Port), grpc.WithInsecure())
	if err != nil {
		sugar.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	bankClient := proto.NewBankServiceClient(conn)

	salaryService := service.New(store, sugar, bankClient)

	StartSalaryWageRoutine(salaryService)

	for {

	}
}

// StartSalaryWageRoutine starts the GetWage method for the provided Salary instance every minute in a separate goroutine.
func StartSalaryWageRoutine(salary *service.Salary) {
	ticker := time.Tick(1 * time.Minute)

	go func() {
		for range ticker {
			err := salary.GetWage()
			if err != nil {
				salary.Sugar.Errorf("Error in GetWage: %v", err)
			}
		}
	}()
}
