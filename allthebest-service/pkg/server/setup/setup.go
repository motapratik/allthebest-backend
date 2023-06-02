package setup

import (
	"allthebest-backend/allthebest-service/config"
	//Database "allthebest-backend/allthebest-service/pkg/shared/database"
)

type Container struct {
	Config *config.Config
	//UserSvc userSvc.Service
	//AuthSvc  authSvc.Service
	//Validate *validator.Validate
}

func Setup(environment, configPath string) *Container {
	// ====== Construct Config
	cfg, err := config.LoadConfig(environment, configPath)
	if err != nil {
		panic(err)
	}

	// ====== Construct Database
	//_ = Database.ConnectDB(cfg.Database)

	//userRepo := gorm.UserSetup(db)

	//crudUserWrapper := goarch.SetupCrudUserWrapper(&cfg.GoarchGrpc)
	//goArchAPIWrapper := goarchApi.NewWrapper(&cfg.GoarchAPIConfig)

	//userSvc := userSvc.NewService(userRepo, crudUserWrapper, goArchAPIWrapper)
	//authSvc := authSvc.NewService(userRepo)

	return &Container{
		Config: cfg,
		//UserSvc:  userSvc,
		//AuthSvc:  authSvc,
		//Validate: validator.New(),
	}
}
