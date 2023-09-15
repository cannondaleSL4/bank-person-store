package service

import (
	"github.com/cannondaleSL4/bank-person-store/common/proto"
	"github.com/cannondaleSL4/bank-person-store/store/db"
	"github.com/cannondaleSL4/bank-person-store/store/models"
	"go.uber.org/zap"
)

type Store struct {
	Products          []*models.Product
	Store             *db.Store
	Sugar             *zap.SugaredLogger
	bankServiceClient proto.BankServiceClient
}
