package postgres

import (
	dbmodel "allthebest-backend/allthebest-service/pkg/shared/database/models"
	"database/sql"
	"fmt"
	"os"
)

const (
	dbclient = "postgres"
)

// ConnectionURL is Rreturns the connection url string
func (db *DBConfig) ConnectionURL() (string, error) {
	password := os.Getenv(db.Password)
	if password == "" {
		return "", fmt.Errorf("%v env variable could not be found", db.Password)
	}

	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		db.Host, db.Port, db.UserName, password, db.DBName,
	), nil
}

// Connect open the connection with the DB
func (db *DBConfig) ConnectDB() (*dbmodel.DBConnection, error) {
	connectionURL, err := db.ConnectionURL()
	if err != nil {
		return nil, err
	}

	dbConnection, err := sql.Open(dbclient, connectionURL)
	if err != nil {
		return nil, err
	}

	if err = dbConnection.Ping(); err != nil {
		return nil, err
	}
	return &dbmodel.DBConnection{PostgreSQL:dbConnection}, nil
}
