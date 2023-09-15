package db

import (
	"database/sql"
	"go.uber.org/zap"
)

type CommonStore struct {
	*sql.DB
	logger *zap.SugaredLogger
}

func New(connection string, logger *zap.SugaredLogger) *CommonStore {
	db, err := sql.Open("postgres", connection)
	if err != nil {
		logger.Fatalf("Could not open config file: %v", err)
	}
	return &CommonStore{db, logger}
}
