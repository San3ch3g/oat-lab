package main

import (
	"fmt"
	"oat-lab-module/internal/di"
	"oat-lab-module/internal/pkg/controllers"
	"oat-lab-module/internal/utils/config"
)

//	@title			Smart Lab
//	@version		1.0
//	@description	API server for Smart Lab
//	@host			localhost:8080
//	@BasePath		/

func main() {
	cfg := config.NewConfig()
	cfg.InitENV()

	container := di.New(cfg)
	db := container.GetDB()

	postgresDB, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("Failed to get database connection: %v", err))
	}
	if err := postgresDB.Ping(); err != nil {
		panic(fmt.Sprintf("Failed to ping database: %v", err))
	}

	storage := container.GetSQLStorage()
	server := controllers.NewServer(storage, cfg)
	err = server.Run(cfg.ServerPort)
	if err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
