package main

type Employee struct {
	ID        int64  `db:"id" json:"id"`
	Firstname string `db:"firstname" json:"firstname"`
	Lastname  string `db:"lastname" json:"lastname"`
	Salary    string `db:"salary" json:"salary"`
}
