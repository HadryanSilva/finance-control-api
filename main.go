package main

import (
	"github.com/HadryanSilva/finance-control-api/config"
	"github.com/HadryanSilva/finance-control-api/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errf("Error initializing config: %v", err)
		return
	}

	router.Initialize()
}
