package postgres

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"
)

type PostgresConnector struct {
	DB *sql.DB
}

func NewPostgresConnector() *PostgresConnector {
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	user := viper.GetString("database.user")
	passwd := viper.GetString("database.password")
	dbName := viper.GetString("database.name")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, passwd, dbName)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return &PostgresConnector{
		DB: db,
	}
}
