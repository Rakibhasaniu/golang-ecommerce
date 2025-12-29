package user

import (
	"main/config"
)

type Handler struct {
	svc    Service
	config *config.Config
}

func NewHandler(svc Service, config *config.Config) *Handler {
	return &Handler{svc: svc, config: config}
}
