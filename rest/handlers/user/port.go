package user

import "main/domain"

type Service interface {
	CreateUser(u domain.User) (*domain.User, error)
	GetUserByEmail(email string, password string) (*domain.User, error)
}
