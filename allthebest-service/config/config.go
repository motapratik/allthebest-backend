package config

import (
	"allthebest-backend/allthebest-service/pkg/shared/utils"
	"fmt"

	"github.com/spf13/viper"
)

func (c *Config) AppAddress() string {
	return fmt.Sprintf(":%v", c.ServerInfo.HttpPort)
}

// LoadConfig is used to load configuration
func LoadConfig(environment string, configPath string) (cfg *Config, err error) {
	fmt.Printf("Loading Configuration for environment: %s", environment)
	if environment == "" {
		return nil, fmt.Errorf("environment not specified, got %s", environment)
	}

	// Load Configuration from JSON File
	conf, err := GetConfigurationFromJson(environment, configPath)
	if err != nil {
		panic(err)
	}
	// Update Credential Information from Env
	confFinal, err := SetCredentialInformation(environment, conf, configPath)
	if err != nil {
		panic(err)
	}
	return confFinal, nil
}

// GetConfigurationFromJson is used to Get Config from JSON file
func GetConfigurationFromJson(environment, configPath string) (cfg *Config, err error) {
	filePath := fmt.Sprintf("%s/%s.json", configPath, environment)
	viper.SetConfigFile(filePath)
	viper.SetConfigType("json")
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	conf := Config{}
	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}
	return &conf, nil
}

// GetCredentialsFromEnv used to GET credential information from ENV file
func GetCredentialsFromEnv(environment, configPath string) (cfg *envConfig, err error) {
	filePath := fmt.Sprintf("%s/%s.json", configPath, environment)
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

// SetCredentialInformation is used to SET credential information
func SetCredentialInformation(environment string, cfgInput *Config, configPath string) (cfgOutput *Config, err error) {
	// Load Credential from .ENV file
	if cfgInput.ServerInfo.LoadCredFromEnvFile {
		// Get Env file credentials
		cfgEnv, err := GetCredentialsFromEnv(environment, configPath)
		if err != nil {
			panic(err)
		}
		// SET Credential Details from .Env file
		cfgInput.JsonWebToken.SigningKeyEnv = cfgEnv.JwtSigningKey
		cfgInput.Database.Password = cfgEnv.DatabasePassword

	} else {
		// ELSE LOAD FROM ENVIRONMENT VARIABLE
		cfgInput.JsonWebToken.SigningKeyEnv, _ = utils.GetEnvironmentVariable("SIGNING_KEY_ENV")
		cfgInput.Database.Password, _ = utils.GetEnvironmentVariable("DB_PASS_ENV")
	}
	return cfgInput, nil
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
