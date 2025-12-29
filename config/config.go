package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configuration *Config

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  bool
}
type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	JwtSecret   string
	DBConfig    *DBConfig
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
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		fmt.Println("JWT secret is required")
		os.Exit(1)
	}
	host := os.Getenv("DB_HOST")
	if host == "" {
		fmt.Println("DB host is required")
		os.Exit(1)
	}
	parsedPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		fmt.Println("Invalid DB port")
		os.Exit(1)
	}
	user := os.Getenv("DB_USER")
	if user == "" {
		fmt.Println("DB user is required")
		os.Exit(1)
	}
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		fmt.Println("DB password is required")
		os.Exit(1)
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Println("DB name is required")
		os.Exit(1)
	}

	sslMode := os.Getenv("DB_SSLMode")
	fmt.Println(sslMode)
	if sslMode == "" {
		fmt.Println("DB SSL mode is required")
		os.Exit(1)
	}
	parsedSSLMode, err := strconv.ParseBool(sslMode)
	if err != nil {
		fmt.Println("Invalid SSL mode (must be true or false)")
		os.Exit(1)
	}
	dbConfig := &DBConfig{
		Host:     host,
		Port:     parsedPort,
		User:     user,
		Password: password,
		Name:     dbName,
		SSLMode:  parsedSSLMode,
	}
	configuration = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    parsedHttpPort,
		JwtSecret:   jwtSecret,
		DBConfig:    dbConfig,
	}

}

func GetConfig() *Config {
	if configuration == nil {
		loadConfig()
	}
	return configuration
}
