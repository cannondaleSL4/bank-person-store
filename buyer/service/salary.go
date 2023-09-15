package service

import (
	"context"
	"github.com/cannondaleSL4/bank-person-store/buyer/db"
	"github.com/cannondaleSL4/bank-person-store/buyer/models"
	"github.com/cannondaleSL4/bank-person-store/common/proto"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

type Salary struct {
	Persons           []*models.Person
	Store             *db.Store
	Sugar             *zap.SugaredLogger
	bankServiceClient proto.BankServiceClient
}

func New(store *db.Store, logger *zap.SugaredLogger, bankService proto.BankServiceClient) *Salary {
	listOfPerson, err := store.GetAllPerson()

	if err != nil {
		logger.Fatalf("Failed to get list of persons: %v", err)
	}

	return &Salary{
		Store:             store,
		Sugar:             logger,
		Persons:           listOfPerson,
		bankServiceClient: bankService}
}

func (s *Salary) GetWage() error {

	shuffledList := s.makeShuffle()

	for _, k := range shuffledList {

		_, err := s.bankServiceClient.IncreaseBalance(context.Background(), &proto.IncreaseBalanceRequest{
			ClientId: k.ID,
			Amount:   k.Salary,
		})

		if err != nil {
			s.Sugar.Errorf("Could not update bank accout of service: %v", err)
			return err
		}
	}
	return nil
}

func (s *Salary) makeShuffle() []*models.Person {
	listOfPerson := s.Persons

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(listOfPerson), func(i, j int) {
		listOfPerson[i], listOfPerson[j] = listOfPerson[j], listOfPerson[i]
	})
	return listOfPerson
}
