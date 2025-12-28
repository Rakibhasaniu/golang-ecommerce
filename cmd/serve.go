package cmd

import (
	"main/config"
	"main/repo"
	"main/rest"
	"main/rest/handlers/product"
	"main/rest/handlers/user"
	"main/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()
	middleware := middleware.NewMiddlewares(cnf)
	productRepo := repo.NewProductRepo()
	productHandler := product.NewHandler(middleware, productRepo)
	userRepo := repo.NewUserRepo()
	userHandler := user.NewHandler(userRepo, cnf)

	rest.NewServer(productHandler, userHandler, cnf).Start()

}
