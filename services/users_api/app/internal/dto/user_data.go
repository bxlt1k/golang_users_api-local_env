package dto

type UserRegistrationData struct {
	Email     string `db:"email" json:"email"`
	FirstName string `db:"firstName" json:"firstName"`
	LastName  string `db:"lastName" json:"lastName"`
	Password  string `db:"password" json:"password"`
}

type UserUpdateData struct {
	Email     string `db:"email" json:"email"`
	FirstName string `db:"firstName" json:"firstName"`
	LastName  string `db:"lastName" json:"lastName"`
	Password  string `db:"password" json:"password"`
}
