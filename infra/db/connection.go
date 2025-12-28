package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString() string {
	return "user=postgres password=password12345 host=localhost port=5432 dbname=ecommerce"

}

func NewConnection() (*sqlx.DB, error) {
	dbsource := GetConnectionString()
	db, err := sqlx.Connect("postgres", dbsource)
	if err != nil {
		return nil, err
	}
	return db, nil
}
