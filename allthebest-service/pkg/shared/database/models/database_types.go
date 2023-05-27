package database

import "database/sql"

// Database holds connection data
type Database struct {
	Host     string
	Name     string
	Password string
	Port     int
	User     string
}

type DBConnection struct {
	PostgreSQL *sql.DB
}

type DBConfig struct {
	Host               string `json:"host"`
	Port               int    `json:"port"`
	UserName           string `json:"user_name"`
	Password           string `json:"password"`
	DBName             string `json:"db_name"`
	MinIdleConnections int    `json:"min_idle_connections"`
	MaxOpenConnections int    `json:"max_open_connections"`
	DebugMode          bool   `json:"debug_mode"`
}
