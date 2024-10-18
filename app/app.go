package app

import (
	"GoProject/config"
	"GoProject/pkg/db"
	"GoProject/pkg/logger"
	"fmt"
)

var (
	// this name is one time only
	service = "Blueprint"
	version = "v1.0.0"
)

func Start() {
	// config instant
	cfg := config.NewConfig()
	fmt.Println(cfg)

	// pass to logger handler instant
	log, _ := logger.NewLogger(cfg)
	fmt.Println(log)

	log.Logger.Info("Logging started for service: ", service+"@"+version)
	_, err := db.NewMysqlDB(cfg)
	if err != nil {
		fmt.Printf("We have a problem connected to database %v", err)
		panic(0)
	}
	log.Logger.Infof("Connected to database %s", cfg.MySQL.MysqlHost+":"+cfg.MySQL.MysqlPort)

	// when painc receover
	if err := recover(); err != nil {
		log.Logger.Fatalf("some panic ...:", err)
	}
	log.Logger.Info("Server Started")

}
