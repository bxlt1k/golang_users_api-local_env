package models

type User struct {
	ID        int64  `db:"id" json:"id"`
	Email     string `db:"email" json:"email"`
	FirstName string `db:"firstName" json:"firstName"`
	LastName  string `db:"lastName" json:"lastName"`
	Password  string `db:"password" json:"password"`
}
