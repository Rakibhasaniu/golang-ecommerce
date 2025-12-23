package cmd

import (
	"main/config"
	"main/rest"
)

func Serve() {
	cnf := config.GetConfig()

	rest.Start(cnf)

}
