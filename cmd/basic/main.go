package main

import (
	"github.com/glebateee/basic/internal/config"
	"github.com/glebateee/basic/internal/placeholder"
	"github.com/glebateee/basic/internal/services"
)

func main() {
	configPath := config.MustLoadConfig()
	services.RegisterDefaultServices(configPath)
	placeholder.Start()
}
