package rest

import (
	"fmt"
	"log"
	"main/config"
	"main/rest/middleware"
	"net/http"
	"strconv"
)

func Start(cnf config.Config) {
	manager := middleware.NewManager()
	manager.Use(middleware.Logger, middleware.HandleCors, middleware.HandlePreflight)

	mux := http.NewServeMux()
	initRoutes(mux, manager)
	wrappedMux := manager.WrapMux(mux)
	fmt.Println("Server started on :" + strconv.Itoa(cnf.HttpPort))

	err := http.ListenAndServe(":"+strconv.Itoa(cnf.HttpPort), wrappedMux)
	if err != nil {
		log.Fatal(err)
	}
}
