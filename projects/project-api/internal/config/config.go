package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Http struct {
	Address string `yaml:"address" address:"address" env-default:"localhost:8081" env-required:"true"`
}

type Config struct {
	Env         string `yaml:"env" address:"address" env-default:"localhost:8081" env-required:"true"`
	StoragePath string `yaml:"storage_path" address:"storage_path" env-required:"true"`
	Http        `yaml:"http_server" address:"http_server" env-required:"true"`
}

func MustLoad() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flags := flag.String("config", "", "Path to config file")
		flag.Parse()
		configPath = *flags
		if configPath == "" {
			log.Fatal("Config Path not set")
		}
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config Path not exist : %s", configPath)
	}
	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("cannot read config file : %s", err.Error())
	}
	return &cfg
}
