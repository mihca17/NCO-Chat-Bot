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

	logger, err := logger.Init(config.LogFile)

	if err != nil {
		log.Fatalf("Не удалось инициализировать логгер: %v", err)
	}
	defer logger.Close()

	db, err := database.InitSQLite(config.DBPath, logger)
	if err != nil {
		logger.Fatal("Ошибка инициализации БД", err)
	}
	defer db.Close()

	repo := repository.NewSQLiteRepository(db.GetDB(), &config)

	getService := services.NewGetService(repo, logger)
	getController := controllers.NewGetController(getService, logger)

	postService := services.NewPostService(repo, logger)
	postController := controllers.NewPostController(postService, logger)

	srv := routers.NewServer(config.Address, config.Port, getController, postController, logger)
	err = srv.Start()
	if err != nil {
		logger.Fatal("Ошибка запуска сервера", err)
		panic(err)
	}
}
