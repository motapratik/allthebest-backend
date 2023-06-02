package postgres

import (
	dbinterface "allthebest-backend/allthebest-service/pkg/shared/database"
	"errors"
)

// NewDatabase is used to Creates a new DB instance
func NewDatabase(cfg *DBConfig) (dbinterface.IDatabase, error) {
	if cfg == nil {
		return nil, errors.New("unable to init Database with no configuration data")
	}

	return &DBConfig{
		Host:               cfg.Host,
		Port:               cfg.Port,
		UserName:           cfg.UserName,
		Password:           cfg.Password,
		DBName:             cfg.DBName,
		MinIdleConnections: cfg.MinIdleConnections,
		MaxOpenConnections: cfg.MaxOpenConnections,
		DebugMode:          cfg.DebugMode,
	}, nil
}
