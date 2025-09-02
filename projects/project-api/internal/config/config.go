package config

import (
	"flag"
	"log/slog"
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
			slog.Error("Config Path not set")
		}
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		slog.Error("Config Path not exist : ", slog.String("configpath", configPath))
	}
	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		slog.Error("cannot read config file : %s", slog.String("error", err.Error()))
	}
	return &cfg
}
