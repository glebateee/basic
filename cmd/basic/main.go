package main

import (
	"github.com/glebateee/basic/internal/config"
	"github.com/glebateee/basic/internal/logging"
	"github.com/glebateee/basic/internal/services"
)

func writeMessage(logger logging.Logger, cfg config.Config) {
	section, ok := cfg.GetSection("main")
	if ok {
		if msg, ok := section.GetString("message"); ok {
			logger.Info(msg)
		} else {
			logger.Panic("cannot find configuration setting")
		}
	} else {
		logger.Panic("config section not found")
	}
}
func main() {
	services.RegisterDefaultServices()
	var cfg config.Config
	if err := services.GetService(&cfg); err != nil {
		panic(err)
	}
	var logger logging.Logger
	if err := services.GetService(&logger); err != nil {
		panic(err)
	}
	writeMessage(logger, cfg)
}
