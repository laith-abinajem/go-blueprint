package config

import (
	"fmt"
	"os"
)

const (
	LOG_FILE = "LOG_FILE"
)

type Config struct {
	BaseUrl string
	Logger  Logger
}
type Logger struct {
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
	LogFile           string
}

func NewConfig() *Config {
	logger := Logger{}
	c := &Config{
		Logger: logger,
	}
	logFile := os.Getenv(LOG_FILE)
	parseError := map[string]string{
		LOG_FILE: "",
	}
	if logFile != "" {
		c.Logger.LogFile = logFile
		parseError[LOG_FILE] = logFile
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
