package service

import (
	"database/sql"
	"github.com/cannondaleSL4/bank-person-store/common/proto"
	"go.uber.org/zap"
)

type CommonService struct {
	*sql.DB
	Sugar             *zap.SugaredLogger
	bankServiceClient proto.BankServiceClient
}
