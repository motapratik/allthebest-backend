package main

import (
	"flag"
	"fmt"

	"github.com/joho/godotenv"
)

var (
	environment string
)

func main() {
	flag.StringVar(&environment, "env", "", "Set the environment (testing, development, production)")
	flag.Parse()

	envVariablesFileName := fmt.Sprintf("../../config/%s.env", environment)
	if err := godotenv.Load(envVariablesFileName); err != nil {
		panic(fmt.Errorf("no .env file found: %s", err.Error()))
	}
	//Http.StartHttpService(container.Setup())
}
