package models

type Product struct {
	ID    int64   `db:"id"`
	Name  string  `db:"last_name"`
	Price float64 `db:"service"`
}
