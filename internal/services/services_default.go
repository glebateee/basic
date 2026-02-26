package services

import (
	"github.com/glebateee/basic/internal/config"
	"github.com/glebateee/basic/internal/logging"
	"github.com/glebateee/basic/internal/templates"
)

func RegisterDefaultServices(configPath string) {
	err := AddSingleton(func() config.Config {
		cfg, err := config.New(configPath)
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
	err = AddSingleton(func(cfg config.Config) templates.TemplateExecutor {
		err = templates.LoadTemplates(cfg)
		return &templates.LayoutTemplateProcessor{}
	})
	if err != nil {
		panic(err)
	}
}
