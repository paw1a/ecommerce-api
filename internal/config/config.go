package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/paw1a/ecommerce-api/pkg/logging"
	"sync"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug" env-required:"true"`
	Listen  struct {
		Type   string `yaml:"type" env-default:"port"`
		BindIP string `yaml:"bind_ip" env-default:"127.0.0.1"`
		Port   string `yaml:"port" env-default:"8080"`
	} `yaml:"listen"`
	DB struct {
		Database string `yaml:"database" env-required:"true"`
		URI      string `yaml:"uri" env-required:"true"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
}

var instance *Config
var once sync.Once

func GetConfig(configPath string) *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("Read application configuration...")
		instance = &Config{}
		if err := cleanenv.ReadConfig(configPath, instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
