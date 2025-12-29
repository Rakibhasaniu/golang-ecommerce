package cmd

import (
	"log"
	"main/config"
	"main/infra/db"
	"main/repo"
	"main/rest"
	"main/rest/handlers/product"
	"main/rest/handlers/user"
	"main/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()
	db, err := db.NewConnection(cnf.DBConfig)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	middleware := middleware.NewMiddlewares(cnf)
	productRepo := repo.NewProductRepo(db)
	productHandler := product.NewHandler(middleware, productRepo)
	userRepo := repo.NewUserRepo(db)
	userHandler := user.NewHandler(userRepo, cnf)

	rest.NewServer(productHandler, userHandler, cnf).Start()

}
