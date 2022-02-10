package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	log "github.com/sirupsen/logrus"
	"sync"
)

type Config struct {
	Listen struct {
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
	JWT struct {
		Secret           string `yaml:"secret" env-required:"true"`
		AccessTokenTime  int64  `yaml:"accessTokenTime" env-default:"15"`
		RefreshTokenTime int64  `yaml:"refreshTokenTime" env-default:"86400"`
	} `yaml:"jwt"`
	Redis struct {
		URI string `yaml:"uri" env-default:"localhost:6379"`
	} `yaml:"redis"`
	Stripe struct {
		Key           string `yaml:"key" env-required:"true"`
		WebhookUrl    string `yaml:"webhookUrl" env-required:"true"`
		WebhookSecret string
	} `yaml:"stripe"`
}

var instance *Config
var once sync.Once

func GetConfig(configPath string) *Config {
	once.Do(func() {
		log.Infof("read config file: %s", configPath)
		instance = &Config{}
		if err := cleanenv.ReadConfig(configPath, instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Info(help)
			log.Fatal(err)
		}
	})
	return instance
}
