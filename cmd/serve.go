package cmd

import (
	"log"
	"main/config"
	"main/infra/db"
	"main/repo"
	"main/rest"
	productHandler "main/rest/handlers/product"
	userHandler "main/rest/handlers/user"
	"main/rest/middleware"
	"main/user"
	"os"
)

func Serve() {
	cnf := config.GetConfig()
	database, err := db.NewConnection(cnf.DBConfig)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	err = db.MigrateDB(database, "./migrations")
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
		os.Exit(1)
	}
	middleware := middleware.NewMiddlewares(cnf)
	// repos
	productRepo := repo.NewProductRepo(database)
	userRepo := repo.NewUserRepo(database)

	//domain
	usrSvc := user.NewService(userRepo)
	// handlers
	productHandler := productHandler.NewHandler(middleware, productRepo)
	userHandler := userHandler.NewHandler(usrSvc, cnf)

	rest.NewServer(productHandler, userHandler, cnf).Start()

}
