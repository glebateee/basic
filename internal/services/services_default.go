package services

import (
	"github.com/glebateee/basic/internal/config"
	"github.com/glebateee/basic/internal/logging"
)

func RegisterDefaultServices() {
	err := AddSingleton(func() config.Config {
		cfg, err := config.New("./config/config.json")
		if err != nil {
			panic(err)
		}
		return cfg
	})
	if err != nil {
		panic(err)
	}

	err = AddSingleton(func(cfg config.Config) logging.Logger {
		return logging.New(cfg)
	})
	if err != nil {
		panic(err)
	}
}
