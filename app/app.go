package app

import (
	"GoProject/config"
	"GoProject/pkg/logger"
	"fmt"
)

func Start() {

	cfg := config.NewConfig()
	fmt.Println(cfg)
	// pass to logger handler instant
	log, _ := logger.NewLogger(cfg)
	fmt.Println(log)
	log.Logger.Info("Application started")

}
