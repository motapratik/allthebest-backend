package postgres

type Database struct {
	Host     string
	Name     string
	Password string
	Port     int
	User     string
}

type DBConfig struct {
	Host               string `json:"host"`
	Port               int    `json:"port"`
	UserName           string `json:"userName"`
	Password           string `json:"password"`
	DBName             string `json:"dbName"`
	MinIdleConnections int    `json:"minIdleConnections"`
	MaxOpenConnections int    `json:"maxOpenConnections"`
	DebugMode          bool   `json:"debugMode"`
}
