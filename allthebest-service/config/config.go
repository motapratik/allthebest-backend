package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func (c *Config) AppAddress() string {
	return fmt.Sprintf(":%v", c.ServerInfo.HttpPort)
}

func LoadConfig(environment string) (cfg *Config, err error) {
	fmt.Printf("Loading Configuration for environment: %s", environment)

	if environment == "" {
		return nil, fmt.Errorf("environment not specified, got %s", environment)
	}

	// Load Configuration from JSON File
	conf, err := GetConfigurationFromJson(environment)
	if err != nil {
		panic(err)
	}

	return conf, nil
}

func GetConfigurationFromJson(environment string) (cfg *Config, err error) {
	filePath := fmt.Sprintf("./../../config/%s.json", environment)
	viper.SetConfigFile(filePath)
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	conf := Config{}
	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}

	return &conf, nil
}

func GetCredentialsFromEnv(environment string) (cfg *envConfig, err error) {
	filePath := fmt.Sprintf("./../../config/%s.env", environment)
	viper.SetConfigFile(filePath)
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	conf := envConfig{}
	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}
	return &conf, nil
}

/*
import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Config holds data of configuration file
type Config struct {
	Server *Server `yaml:"server"`
	DB     *DB     `yaml:"db"`
	JWT    *JWT    `yaml:"jwt"`
	LOG
}

// Load returns a *Config with all the configuration data
// needed to start the server
func Load(environment string) (*Config, error) {
	if environment == "" {
		return nil, fmt.Errorf("Environment not specified, got %s", environment)
	}

	//fileName := fmt.Sprintf("./config/%s.yml", environment)
	fileName := fmt.Sprintf("./../../config/%s.yml", environment)

	bytesYmlfile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("Unable to read config file: %s, %s", fileName, err.Error())
	}

	cfg := &Config{}
	if err = yaml.Unmarshal(bytesYmlfile, cfg); err != nil {
		return nil, fmt.Errorf("unable to unmarshall configuration: %s", err.Error())
	}

	return cfg, nil
}
*/
