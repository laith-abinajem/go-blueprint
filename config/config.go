package config

import (
	"fmt"
	"os"
)

const (
	LOG_FILE       = "LOG_FILE"
	MYSQL_HOST     = "MYSQL_HOST"
	MYSQL_PORT     = "MYSQL_PORT"
	MYSQL_USER     = "MYSQL_USER"
	MYSQL_PASSWORD = "MYSQL_PASSWORD"
	MYSQL_DB       = "MYSQL_DB"
)

type Config struct {
	BaseUrl string
	Logger  Logger
	MySQL   MySQL
}
type Logger struct {
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
	LogFile           string
}
type MySQL struct {
	MysqlHost     string
	MysqlPort     string
	MysqlUser     string
	MysqlPassword string
	MysqlDBName   string
}

func NewConfig() *Config {
	logger := Logger{}
	mysql := MySQL{}
	c := &Config{
		Logger: logger,
		MySQL:  mysql,
	}
	logFile := os.Getenv(LOG_FILE)
	parseError := map[string]string{
		LOG_FILE:       "",
		MYSQL_HOST:     "",
		MYSQL_PORT:     "",
		MYSQL_USER:     "",
		MYSQL_PASSWORD: "",
		MYSQL_DB:       "",
	}
	if logFile != "" {
		c.Logger.LogFile = logFile
		parseError[LOG_FILE] = logFile
	}

	mysqlHost := os.Getenv(MYSQL_HOST)

	if mysqlHost != "" {
		c.MySQL.MysqlHost = mysqlHost
		parseError[MYSQL_HOST] = mysqlHost
	}

	mysqlPort := os.Getenv(MYSQL_PORT)
	if mysqlPort != "" {
		c.MySQL.MysqlPort = mysqlPort
		parseError[MYSQL_PORT] = mysqlPort

	}

	mysqlUser := os.Getenv(MYSQL_USER)

	if mysqlUser != "" {
		c.MySQL.MysqlUser = mysqlUser
		parseError[MYSQL_USER] = mysqlUser
	}

	mysqlPassword := os.Getenv(MYSQL_PASSWORD)

	if mysqlPassword != "" {
		c.MySQL.MysqlPassword = mysqlPassword
		parseError[MYSQL_PASSWORD] = mysqlPassword
	}
	MysqlDBName := os.Getenv(MYSQL_DB)

	if MysqlDBName != "" {
		c.MySQL.MysqlDBName = MysqlDBName
		parseError[MYSQL_DB] = MysqlDBName
	}
	exitParse := false
	for k, v := range parseError {
		if v == "" {
			exitParse = true
			fmt.Printf("%s = %s\n", k, v)
		}
	}

	// one faild
	if exitParse {
		panic("Env vars not set see list")
	}
	return c
}
