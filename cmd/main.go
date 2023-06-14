package main

import (
	"api-public-platform/config"
	"api-public-platform/internal/db"
	"api-public-platform/pkg/model"
	"api-public-platform/pkg/routers"
	"api-public-platform/pkg/utils"
	"log"
	"strconv"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
		return
	}
	db, err := db.ConnectDatabase(config.ServerCfg)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
		return
	}
	err = db.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.API{},
		&model.Permission{},
		&model.APICallHistory{},
		&model.UserAPI{},
	)
	if err != nil {
		log.Fatalf("Could not migrate database: %v", err)
		return
	}
	err = utils.InitTrans("zh")
	if err != nil {
		log.Fatalf("Could not initialize translation: %v", err)
		return
	}
	appRouters := routers.NewRouter()
	app := appRouters.SetUpRouter()
	port := strconv.Itoa(config.ServerCfg.Port)
	err = app.Run(":" + port)
	if err != nil {
		log.Fatalf("Could not run server: %v", err)
		return
	}
}
