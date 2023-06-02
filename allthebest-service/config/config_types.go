package config

import (
	db "allthebest-backend/allthebest-service/pkg/shared/database/models"
	rest_client "allthebest-backend/allthebest-service/pkg/shared/rest/resty_client"
)

// Config holds data of configuration file
type Config struct {
	ServerInfo   Server      `json:"serverInfo"`
	Database     db.DBConfig `json:"database"`
	JsonWebToken JWT         `json:"jwt"`
	LogConfig    LOG         `json:"logConfig"`
	APIConfig    APIConfig   `json:"APIConfig"`
}

// envConfig is struct to map env values
type envConfig struct {
	DatabasePassword string `mapstructure:"DB_PASS_ENV"`
	JwtSigningKey    string `mapstructure:"JWT_SIGNING_KEY_ENV"`
}

// Server contain Server IP/Port configuration parameters
type Server struct {
	Name                string `json:"name"`
	Version             string `json:"version"`
	HttpPort            int    `json:"httpPort"`
	Host                string `json:"host"`
	LoadCredFromEnvFile bool   `json:"loadServerCredentialFromEnvFile"`
}

// JWT contain JWT related configuration parameters
type JWT struct {
	Duration         int    `json:"durationMinutes"`
	RefreshDuration  int    `json:"refreshDurationMinutes"`
	SigningAlgorithm string `json:"signingAlgorithm"`
	SigningKeyEnv    string `json:"signingKeyEnv"`
}

// LOG contain logging related configurations
type LOG struct {
	Enabled  bool   `json:"enabled"`
	Filename string `json:"fileName"`
	LogPath  string `json:"logPath"`
}

type APIConfig struct {
	RestOptions rest_client.Options `json:"restAPIOptions"`
}

// Config holds data of configuration file
// type Config struct {
// 	Server *Server         `yaml:"server"`
// 	DB     *DB             `yaml:"db"`
// 	JWT    *JWT            `yaml:"jwt"`
// 	LOG    *LOG            `yaml:"log"`
// 	REST   *RestAPIOptions `yaml:"rest_api_options"`
// 	RPC    *RPCOptions     `yaml:"rpc_options"`
// }

// DB contain Databse related configuration parameters
// type DB struct {
// 	Host               string `yaml:"host"`
// 	Port               int    `yaml:"port"`
// 	UserName           string `yaml:"user_name"`
// 	Password           string `yaml:"password"`
// 	DBName             string `yaml:"db_name"`
// 	MinIdleConnections int    `yaml:"min_idle_connections"`
// 	MaxOpenConnections int    `yaml:"max_open_connections"`
// 	DebugMode          bool   `yaml:"debug_mode"`
// }

// JWT contain JWT related configuration parameters
// type JWT struct {
// 	Duration         int    `yaml:"duration_minutes"`
// 	RefreshDuration  int    `yaml:"refresh_duration_minutes"`
// 	SigningAlgorithm string `yaml:"signing_algorithm"`
// 	SigningKeyEnv    string `yaml:"signing_key_env"`
// }

// // LOG contain logging related configurations
// type LOG struct {
// 	Enabled  bool   `yaml:"enabled"`
// 	Filename string `yaml:"file_name"`
// 	Logpath  string `yaml:"log_path"`
// }

// RestAPIOptions contain REST API related configurations
// type RestAPIOptions struct {
// 	Address           string `yaml:"address"`
// 	Timeout           int    `yaml:"timeout"`
// 	DebugMode         bool   `yaml:"debug_mode"`
// 	WithProxy         bool   `yaml:"with_proxy"`
// 	SkipTLS           bool   `yaml:"skip_tls"`
// 	SkipCheckRedirect bool   `yaml:"skip_check_redirect"`
// }

// // RPCOptions contain GRPC related options
// type RPCOptions struct {
// 	Address string `yaml:"address"`
// 	Timeout int    `yaml:"timeout"`
// }
