package models

import "math/big"

type Bank struct {
	ID             int64      `db:"id"`
	ClientID       int64      `db:"client_id"`
	FirstName      string     `db:"first_name"`
	LastName       string     `db:"last_name"`
	AccountBalance *big.Float `db:"account_balance"`
}
