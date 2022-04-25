package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	log "github.com/sirupsen/logrus"
	"os"
	"sync"
)

type Config struct {
	Listen struct {
		Host string
		Port string
	}
	DB struct {
		Database string
		URI      string
		Username string
		Password string
	}
	JWT struct {
		Secret           string
		AccessTokenTime  int64 `yaml:"accessTokenTime" env-default:"15"`
		RefreshTokenTime int64 `yaml:"refreshTokenTime" env-default:"86400"`
	} `yaml:"jwt"`
	Redis struct {
		URI string
	}
	Stripe struct {
		Key           string
		WebhookUrl    string
		WebhookSecret string
	}
	Test struct {
		Database   string
		DBURI      string
		DBUsername string
		DBPassword string
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

		instance.Listen.Host = os.Getenv("HOST")
		instance.Listen.Port = os.Getenv("PORT")

		instance.Stripe.Key = os.Getenv("STRIPE_KEY")
		instance.Stripe.WebhookUrl = "http://" + instance.Listen.Host + ":" + instance.Listen.Port + "/api/v1/payment/webhook"

		instance.JWT.Secret = os.Getenv("JWT_SECRET")

		instance.DB.Username = os.Getenv("DB_USERNAME")
		instance.DB.Password = os.Getenv("DB_PASSWORD")
		instance.DB.Database = os.Getenv("DB_NAME")
		instance.DB.URI = os.Getenv("DB_URI")

		instance.Test.DBUsername = os.Getenv("TEST_DB_USERNAME")
		instance.Test.DBPassword = os.Getenv("TEST_DB_PASSWORD")
		instance.Test.Database = os.Getenv("TEST_DB_NAME")
		instance.Test.DBURI = os.Getenv("TEST_DB_URI")

		instance.Redis.URI = os.Getenv("REDIS_URI")
	})
	return instance
}
