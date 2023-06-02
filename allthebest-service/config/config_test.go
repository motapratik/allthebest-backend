package config

import (
	"fmt"
	"testing"
)

// TestLoadConfig is used to test LoadConfig function
func TestLoadConfig(t *testing.T) {
	environment := "development"
	configPath := "."
	cfg, err := LoadConfig(environment, configPath)
	if err != nil {
		fmt.Printf("error to load config for %s environment, err: %v", environment, err)
	}
	t.Log(cfg)
}
