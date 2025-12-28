package user

import (
	"main/config"
	"main/repo"
)

type Handler struct {
	userRepo repo.UserRepo
	config   *config.Config
}

func NewHandler(userRepo repo.UserRepo, config *config.Config) *Handler {
	return &Handler{userRepo: userRepo, config: config}
}
