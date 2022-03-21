package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	log "github.com/sirupsen/logrus"
	"os"
	"sync"
)

type Config struct {
	Listen struct {
		Host string `yaml:"host" env-default:"127.0.0.1"`
		Port string `yaml:"port" env-default:"8080"`
	} `yaml:"listen"`
	DB struct {
		Database string `yaml:"database" env-required:"true"`
		URI      string `yaml:"uri" env-required:"true"`
		Username string
		Password string
	} `yaml:"db"`
	JWT struct {
		Secret           string
		AccessTokenTime  int64 `yaml:"accessTokenTime" env-default:"15"`
		RefreshTokenTime int64 `yaml:"refreshTokenTime" env-default:"86400"`
	} `yaml:"jwt"`
	Redis struct {
		URI string `yaml:"uri" env-default:"127.0.0.1:6379"`
	} `yaml:"redis"`
	Stripe struct {
		Key           string
		WebhookUrl    string
		WebhookSecret string
	}
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

		instance.Stripe.WebhookUrl = "http://" + instance.Listen.Host + ":" + instance.Listen.Port + "/api/v1/payment/webhook"
		instance.Stripe.Key = os.Getenv("STRIPE_KEY")
		instance.JWT.Secret = os.Getenv("JWT_SECRET")
		instance.DB.Username = os.Getenv("DB_USERNAME")
		instance.DB.Password = os.Getenv("DB_PASSWORD")
	})
	return instance
}
