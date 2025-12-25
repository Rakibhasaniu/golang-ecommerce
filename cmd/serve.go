package cmd

import (
	"main/config"
	"main/rest"
	"main/rest/handlers/product"
	"main/rest/handlers/user"
	"main/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()
	middleware := middleware.NewMiddlewares(cnf)
	productHandler := product.NewHandler(middleware)
	userHandler := user.NewHandler()

	rest.NewServer(productHandler, userHandler, cnf).Start()

}
