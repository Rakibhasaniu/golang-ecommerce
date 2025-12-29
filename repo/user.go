package repo

import (
	"database/sql"
	"main/domain"
	"main/user"

	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	user.UserRepo
}
type userRepo struct {
	dbCon *sqlx.DB
}

func NewUserRepo(dbCon *sqlx.DB) UserRepo {
	repo := &userRepo{
		dbCon: dbCon,
	}
	return repo
}

// CREATE TABLE users (
// 	id SERIAL PRIMARY KEY,
// 	first_name VARCHAR(100) NOT NULL,
// 	last_name VARCHAR(100) NOT NULL,
// 	email VARCHAR(255) UNIQUE NOT NULL,
// 	password VARCHAR(255) NOT NULL,
// 	is_shop_owner BOOLEAN DEFAULT FALSE,
// 	created_at TIMESTAMP WITH TIME ZONE  DEFAULT CURRENT_TIMESTAMP,
// 	updated_at TIMESTAMP WITH TIME ZONE  DEFAULT CURRENT_TIMESTAMP
// )

func (u *userRepo) CreateUser(user domain.User) (*domain.User, error) {
	query := `INSERT INTO users (
	first_name, 
	last_name, 
	email, 
	password, 
	is_shop_owner)
	VALUES (
	:first_name, 
	:last_name, 
	:email, 
	:password, 
	:is_shop_owner) RETURNING id`
	var userId int
	rows, err := u.dbCon.NamedQuery(query, user)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		rows.Scan(&userId)
	}
	user.ID = userId
	return &user, nil
}

func (u *userRepo) GetUserByEmail(email string, password string) (*domain.User, error) {
	query := `SELECT id, first_name, last_name, email, password, is_shop_owner FROM users WHERE email = $1 AND password = $2`
	var user domain.User
	err := u.dbCon.Get(&user, query, email, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
