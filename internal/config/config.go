package config

import (
	"log"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/lib/pq"
)

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	Path       string `yaml:"path" env-required:"true"`
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	User        string        `yaml:"user" env-required:"true"`
	Password    string        `yaml:"password" env-required:"true" env:"HTTP_SERVER_PASSWORD"`
}
type Client struct {
	Address      string        `yaml:"address"`
	Timeout      time.Duration `yaml:"timeout"`
	RetriesCount int           `yaml:"retriesCount"`
}
type ClientsConfig struct {
	Auth Client `yaml:"auth"`
}

func MustLoad() *Config {
	configPath := "./config/prod.yaml"
	// if configPath == "" {
	// 	log.Fatal("CONFIG_PATH is not set")
	// }

	// // check if file exists
	// if _, err := os.Stat(configPath); os.IsNotExist(err) {
	// 	log.Fatalf("config file does not exist: %s", configPath)
	// }

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
