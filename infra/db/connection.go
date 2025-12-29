package db

import (
	"fmt"
	"main/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(cnf *config.DBConfig) string {
	sslMode := "disable"
	if cnf.SSLMode {
		sslMode = "require"
	}
	connString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		cnf.User,
		cnf.Password,
		cnf.Host,
		cnf.Port,
		cnf.Name,
		sslMode,
	)
	return connString

}

func NewConnection(cnf *config.DBConfig) (*sqlx.DB, error) {
	dbsource := GetConnectionString(cnf)
	db, err := sqlx.Connect("postgres", dbsource)
	if err != nil {
		return nil, err
	}
	return db, nil
}
