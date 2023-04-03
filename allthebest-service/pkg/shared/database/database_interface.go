package database

import (
	dbmodel "allthebest-backend/allthebest-service/pkg/shared/database/models"
)

// IDatabase interface is used to create New Database and connect it
type IDatabase interface {
	ConnectDB() (*dbmodel.DBConnection, error)
}