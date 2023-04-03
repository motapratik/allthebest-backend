package postgres

type Database struct {
	Host     string
	Name     string
	Password string
	Port     int
	User     string
}

type DBConfig struct {
	Host               string `yaml:"host"`
	Port               int    `yaml:"port"`
	UserName           string `yaml:"user_name"`
	Password           string `yaml:"password"`
	DBName             string `yaml:"db_name"`
	MinIdleConnections int    `yaml:"min_idle_connections"`
	MaxOpenConnections int    `yaml:"max_open_connections"`
	DebugMode          bool   `yaml:"debug_mode"`
}