package config

import (
	"os"
	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	LogFilePath string
	Redis RedisConfig
	Sender SenderConfig
}

func ReadConfigFromFile(path string) (*Config, error) {
	byt, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	err = toml.Unmarshal(byt, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}