package main

import (
	"NCO-Chat-Bot/config"
	"NCO-Chat-Bot/controllers"
	"NCO-Chat-Bot/database/database"
	"NCO-Chat-Bot/database/repository"
	"NCO-Chat-Bot/logger"
	"NCO-Chat-Bot/routers"
	"NCO-Chat-Bot/services"
	"log"

	_ "modernc.org/sqlite"
)

func main() {
	config := config.DefaultConfig()

	err := logger.Init(config.LogFile)
	if err != nil {
		log.Fatalf("Не удалось инициализировать логгер: %v", err)
	}
	defer logger.Close()

	db, err := database.InitSQLite(config.DBPath)
	if err != nil {
		logger.Fatal("Ошибка инициализации БД", err)
	}
	defer db.Close()

	repo := repository.NewSQLiteRepository(db.GetDB(), &config)

	getService := services.NewGetService(repo)
	getController := controllers.NewGetController(getService)

	postService := services.NewPostService(repo)
	postController := controllers.NewPostController(postService)

	srv := routers.NewServer(config.Address, config.Port, getController, postController)
	err = srv.Start()
	if err != nil {
		panic(err)
	}
}
