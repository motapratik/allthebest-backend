package config

// Config holds data of configuration file
type Config struct {
	Server *Server         `yaml:"server"`
	DB     *DB             `yaml:"db"`
	JWT    *JWT            `yaml:"jwt"`
	LOG    *LOG            `yaml:"log"`
	REST   *RestAPIOptions `yaml:"rest_api_options"`
	RPC    *RPCOptions     `yaml:"rpc_options"`
}

// Server contain Server IP/Port configuration parameters
type Server struct {
	Name     string `yaml:"name"`
	Version  string `yaml:"version"`
	HttpPort int    `yaml:"http_port"`
	GrpcPort int    `yaml:"grpc_port"`
	Host     string `yaml:"host"`
}

// DB contain Databse related configuration parameters
type DB struct {
	Host               string `yaml:"host"`
	Port               int    `yaml:"port"`
	UserName           string `yaml:"user_name"`
	Password           string `yaml:"password"`
	DBName             string `yaml:"db_name"`
	MinIdleConnections int    `yaml:"min_idle_connections"`
	MaxOpenConnections int    `yaml:"max_open_connections"`
	DebugMode          bool   `yaml:"debug_mode"`
}

// JWT contain JWT related configuration parameters
type JWT struct {
	Duration         int    `yaml:"duration_minutes"`
	RefreshDuration  int    `yaml:"refresh_duration_minutes"`
	SigningAlgorithm string `yaml:"signing_algorithm"`
	SigningKeyEnv    string `yaml:"signing_key_env"`
}

// LOG contain logging related configurations
type LOG struct {
	Enabled  bool   `yaml:"enabled"`
	Filename string `yaml:"file_name"`
	Logpath  string `yaml:"log_path"`
}

// RestAPIOptions contain REST API related configurations
type RestAPIOptions struct {
	Address           string `yaml:"address"`
	Timeout           int    `yaml:"timeout"`
	DebugMode         bool   `yaml:"debug_mode"`
	WithProxy         bool   `yaml:"with_proxy"`
	SkipTLS           bool   `yaml:"skip_tls"`
	SkipCheckRedirect bool   `yaml:"skip_check_redirect"`
}

// RPCOptions contain GRPC related options
type RPCOptions struct {
	Address string `yaml:"address"`
	Timeout int    `yaml:"timeout"`
}
