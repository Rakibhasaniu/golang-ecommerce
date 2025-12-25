package rest

import (
	"fmt"
	"log"
	"main/config"
	"main/rest/handlers/product"
	"main/rest/handlers/user"
	"main/rest/middleware"
	"net/http"
	"strconv"
)

type Server struct {
	cnf            *config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(
	productHandler *product.Handler,
	userHandler *user.Handler,
	cnf *config.Config,
) *Server {
	return &Server{
		productHandler: productHandler,
		userHandler:    userHandler,
		cnf:            cnf,
	}
}
func (s *Server) Start() {
	manager := middleware.NewManager()
	manager.Use(middleware.Logger, middleware.HandleCors, middleware.HandlePreflight)

	mux := http.NewServeMux()
	// initRoutes(mux, manager)
	wrappedMux := manager.WrapMux(mux)
	s.productHandler.ProductRoutes(mux, manager)
	s.userHandler.UserRoutes(mux, manager)
	fmt.Println("Server started on :" + strconv.Itoa(s.cnf.HttpPort))

	err := http.ListenAndServe(":"+strconv.Itoa(s.cnf.HttpPort), wrappedMux)
	if err != nil {
		log.Fatal(err)
	}
}
