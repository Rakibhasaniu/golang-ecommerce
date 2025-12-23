package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

var configuration Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
		os.Exit(1)
	}
	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service name is required")
		os.Exit(1)
	}
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("HTTP port is required")
		os.Exit(1)
	}
	parsedHttpPort, err := strconv.Atoi(httpPort)
	if err != nil {
		fmt.Println("Invalid HTTP port")
		os.Exit(1)
	}

	configuration = Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    parsedHttpPort,
	}

}

func GetConfig() Config {
	loadConfig()
	return configuration
}
