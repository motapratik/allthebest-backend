package api

import (
	"allthebest-backend/allthebest-service/config"
	"allthebest-backend/allthebest-service/pkg/shared/database/postgres"
	"fmt"

	"github.com/jinzhu/copier"
)

// StartHttpService is used to initialize server
func StartHttpService(cfg *config.Config) {

	dbConfig := postgres.DBConfig{}
	err := copier.Copy(&dbConfig, cfg.Database)
	if err != nil {
		panic(err)
	}
	database, err := postgres.NewDatabase(&dbConfig)
	if err != nil {
		panic(fmt.Errorf("problem occured while creating database. %s", err.Error()))
	}
	//dbConnection
	_, err = database.ConnectDB()
	if err != nil {
		panic(fmt.Errorf("problem occured while connecting to database. %s", err.Error()))
	}

	/*
		mux := gorillamultiplexer.New(cfg.Server.Host, cfg.Server.Port)

		database, err := postgres.NewDatabase(cfg.DB)
		if err != nil {
			panic(fmt.Errorf("problem occured while creating database. %s", err.Error()))
		}

		dbConnection, err := database.Connect()
		if err != nil {
			panic(fmt.Errorf("problem occured while connecting to database. %s", err.Error()))
		}

		jwt, err := jwt.New(cfg.JWT.SigningKeyEnv, cfg.JWT.SigningAlgorithm, cfg.JWT.Duration)
		if err != nil {
			panic(fmt.Errorf("problem occured while creating JWT middleware. %s", err.Error()))
		}

		registerPublicRoutes(mux, dbConnection, *jwt)
		registerPrivateRoutes(mux, dbConnection, *jwt)

		mux.Use(jwt.MWFunc)

		srv := server.New(mux)
		srv.Serve()
	*/

}
