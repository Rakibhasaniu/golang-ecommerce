package user

import (
	"main/domain"
	userHandler "main/rest/handlers/user"
)

type Service interface {
	userHandler.Service
}
type UserRepo interface {
	CreateUser(u domain.User) (*domain.User, error)
	GetUserByEmail(email string, password string) (*domain.User, error)
}
