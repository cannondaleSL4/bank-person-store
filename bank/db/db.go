package db

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type Store struct {
	*sql.DB
	logger *zap.SugaredLogger
}

func New(connection string, logger *zap.SugaredLogger) *Store {
	db, err := sql.Open("postgres", connection)
	if err != nil {
		logger.Fatalf("Could not open config file: %v", err)
	}
	return &Store{db, logger}
}

func (s *Store) IncreaseBalance(clientID int64, amount decimal.Decimal) error {
	_, err := s.DB.Exec("UPDATE Bank SET account_balance = account_balance + $1 WHERE client_id = $2", amount, clientID)
	return err
}

func (s *Store) DecreaseBalance(clientID int64, amount decimal.Decimal) error {
	_, err := s.DB.Exec("UPDATE Bank SET account_balance = account_balance - $1 WHERE client_id = $2", amount, clientID)
	return err
}
