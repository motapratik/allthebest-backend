package config

import (
	db "allthebest-backend/allthebest-service/pkg/shared/database/models"
	rest_client "allthebest-backend/allthebest-service/pkg/shared/rest/resty_client"
)

// Config holds data of configuration file
type Config struct {
	ServerInfo   Server      `json:"server"`
	Database     db.DBConfig `json:"database"`
	JsonWebToken JWT         `json:"jwt"`
	LogConfig    LOG         `json:"logConfig"`
	APIConfig    APIConfig   `json:"APIConfig"`
}

// Server contain Server IP/Port configuration parameters
type Server struct {
	Name     string `json:"name"`
	Version  string `json:"version"`
	HttpPort int    `json:"httpPort"`
	Host     string `json:"host"`
}

// JWT contain JWT related configuration parameters
type JWT struct {
	Duration         int    `json:"duration_minutes"`
	RefreshDuration  int    `json:"refresh_duration_minutes"`
	SigningAlgorithm string `json:"signing_algorithm"`
	SigningKeyEnv    string `json:"signing_key_env"`
}

// LOG contain logging related configurations
type LOG struct {
	Enabled  bool   `json:"enabled"`
	Filename string `json:"file_name"`
	Logpath  string `json:"log_path"`
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
