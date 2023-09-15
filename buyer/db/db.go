package db

import (
	"github.com/cannondaleSL4/bank-person-store/buyer/models"
	"github.com/cannondaleSL4/bank-person-store/common/db"
	"go.uber.org/zap"
)

type Store struct {
	CommonStore *db.CommonStore
}

func New(connection string, logger *zap.SugaredLogger) *Store {
	return &Store{CommonStore: db.New(connection, logger)}
}

//func (s *Store) GetPersonByID(id int64) (*models.Person, error) {
//	var p models.Person
//	err := s.QueryRow("SELECT id, first_name, last_name, service FROM Person WHERE id=$1", id).Scan(&p.ID, &p.FirstName, &p.LastName, &p.Salary)
//	if err != nil {
//		return nil, err
//	}
//	return &p, nil
//}

func (s *Store) GetAllPerson() ([]*models.Person, error) {
	var list []*models.Person
	rows, err := s.CommonStore.Query("SELECT * FROM Person")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		p := &models.Person{}
		err := rows.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Salary)
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
