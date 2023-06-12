package main

import (
	"api-public-platform/internal/db"
	"api-public-platform/pkg/model"
	"api-public-platform/pkg/routers"
	"api-public-platform/pkg/utils"
	"log"
)

func main() {
	// ...
	db, err := db.ConnectDatabase()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
		return
	}
	err = db.AutoMigrate(&model.User{}, &model.Role{}, &model.API{}, &model.Permission{})
	if err != nil {
		log.Fatalf("Could not migrate database: %v", err)
		return
	}
	err = utils.InitTrans("zh")
	if err != nil {
		log.Fatalf("Could not initialize translation: %v", err)
		return
	}
	app := routers.SetUpRouter()
	err = app.Run(":8080")
	if err != nil {
		log.Fatalf("Could not run server: %v", err)
		return
	}
}
