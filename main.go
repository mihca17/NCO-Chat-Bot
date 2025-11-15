package main

import (
	"NCO-Chat-Bot/config"
	"NCO-Chat-Bot/database/database"
	"NCO-Chat-Bot/database/repository"
	"NCO-Chat-Bot/logger"
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

	err = StartServer("localhost", "8080", repo)
	if err != nil {
		panic(err)
	}
}
