package repository

import (
	"github.com/jmoiron/sqlx"
	"users_api/internal/models"
)

type UserRepository struct {
	MySqlConn *sqlx.DB
}

func NewUserRepository(mySqlConn *sqlx.DB) *UserRepository {
	return &UserRepository{
		MySqlConn: mySqlConn,
	}
}

func (ur *UserRepository) GetUsers(page int64) ([]*models.User, error) {
	sql := `
		SELECT 
		    * 
		FROM users
		LIMIT ?, ?
	`

	var count int64 = 2
	var offset int64 = 0

	if page > 0 {
		offset = page * count
	}

	var users []*models.User
	if err := ur.MySqlConn.Select(&users, sql, offset, count); err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) GetUserByID(id int64) (*models.User, error) {
	sql := `
		SELECT 
		    * 
		FROM users
		WHERE id = ?
	`

	var user []*models.User
	if err := ur.MySqlConn.Select(&user, sql, id); err != nil {
		return nil, err
	}

	if len(user) < 1 {
		return nil, nil
	}

	return user[0], nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	sql := `
		SELECT 
		    * 
		FROM users
		WHERE email = ?
	`

	var user []*models.User
	if err := ur.MySqlConn.Select(&user, sql, email); err != nil {
		return nil, err
	}

	if len(user) < 1 {
		return nil, nil
	}

	return user[0], nil
}

func (ur *UserRepository) SaveUser(user *models.User) (int64, error) {
	sql := `
		INSERT INTO users
		(email, firstName, lastName, password) VALUE 
		(?, ?, ?, ?)
	`

	var args []interface{}
	args = append(args,
		user.Email, user.FirstName,
		user.LastName, user.Password,
	)

	result, err := ur.MySqlConn.Exec(sql, args...)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}

func (ur *UserRepository) DeleteUserByID(id int64) (*models.User, error) {
	sql := `
		DELETE FROM users
		WHERE id = ?
	`

	var user []*models.User
	if err := ur.MySqlConn.Select(&user, sql, id); err != nil {
		return nil, err
	}

	if len(user) < 1 {
		return nil, nil
	}

	return user[0], nil
}

func (ur *UserRepository) UpdateUser(id int64, user *models.User) error {
	sql := `
		UPDATE users 
		SET email = ?, firstName = ?, lastName = ?, password = ?
		WHERE id = ?
	`
	var args []interface{}
	args = append(args,
		user.Email, user.FirstName,
		user.LastName, user.Password,
		id,
	)

	_, err := ur.MySqlConn.Exec(sql, args...)
	if err != nil {
		return err
	}

	return nil
}
