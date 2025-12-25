package user

import (
	"net/http"

	"main/rest/middleware"
)

func (h *Handler) UserRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.HandleFunc("POST /users/create", h.CreateUser)
	mux.HandleFunc("POST /users/login", h.Login)

}
