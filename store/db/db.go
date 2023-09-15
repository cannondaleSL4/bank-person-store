package db

import (
	"database/sql"
	"github.com/cannondaleSL4/bank-person-store/store/models"
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

func (s *Store) GetProduct() ([]*models.Product, error) {
	var list []*models.Product
	rows, err := s.Query("SELECT * FROM Person")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		p := &models.Product{}
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		list = append(list, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}
