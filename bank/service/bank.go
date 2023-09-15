package service

import (
	"context"
	"github.com/cannondaleSL4/bank-person-store/common/proto"
	"github.com/cannondaleSL4/bank/db"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type BankServer struct {
	proto.UnimplementedBankServiceServer // Note: this is an embedding, not a field!
	store                                *db.Store
	logger                               *zap.SugaredLogger
}

func New(store *db.Store, logger *zap.SugaredLogger) *BankServer {
	return &BankServer{
		logger: logger,
		store:  store,
	}
}

func (s *BankServer) IncreaseBalance(ctx context.Context, req *proto.IncreaseBalanceRequest) (*proto.BalanceResponse, error) {

	// Check for cancellation before starting a potentially long-running operation
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	err := s.store.IncreaseBalance(req.ClientId, decimal.NewFromFloat(req.Amount))
	if err != nil {
		return nil, err
	}
	return &proto.BalanceResponse{Message: "Successfully increased balance."}, nil
}

func (s *BankServer) DecreaseBalance(ctx context.Context, req *proto.DecreaseBalanceRequest) (*proto.BalanceResponse, error) {

	// Check for cancellation before starting a potentially long-running operation
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	err := s.store.DecreaseBalance(req.ClientId, decimal.NewFromFloat(req.Amount))
	if err != nil {
		return nil, err
	}
	return &proto.BalanceResponse{Message: "Successfully decreased balance."}, nil
}
