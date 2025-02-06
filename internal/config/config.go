package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)


type Config struct {
	Env string `yaml:"env" env-required:"true"`
	Dsn string `yaml:"dsn" env-required:"true"`
	Server SettingServer `yaml:"server"  env-required:"true"`
}



type SettingServer struct {
	Addr string `yaml:"address" env-required:"true"`
}


func Deal() (*Config, error) {
	configPath := os.Getenv("configPath")
	
	if configPath == "" {
		return nil, fmt.Errorf("config path empty")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, fmt.Errorf("failed read config: %w", err)
	}

	return &cfg, nil
}