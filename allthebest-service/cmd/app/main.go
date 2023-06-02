package main

import (
	"flag"
	"fmt"

	"allthebest-backend/allthebest-service/config"
	"allthebest-backend/allthebest-service/pkg/api"
)

var (
	environment string
)

const CONFIG_PATH string = "./../../config"

func main() {
	flag.StringVar(&environment, "env", "", "Set the environment (testing, development, production)")
	flag.Parse()
	if (environment != "development") && (environment != "testing") && (environment != "production"){
		panic(fmt.Errorf("enter valid environment name"))
	}
	// Load Configuration from JSON and ENV files
	cfg, err := config.LoadConfig(environment, CONFIG_PATH)
	if err != nil {
		panic(err)
	}
	api.StartHttpService(cfg)
}
