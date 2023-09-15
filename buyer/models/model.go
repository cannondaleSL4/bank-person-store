package models

type Person struct {
	ID        int64   `db:"id"`
	FirstName string  `db:"first_name"`
	LastName  string  `db:"last_name"`
	Salary    float64 `db:"service"`
}
