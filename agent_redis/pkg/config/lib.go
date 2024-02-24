package config

import (
	"github.com/pelletier/go-toml/v2"
	"path/filepath"
	"os"
)

type Config struct {
	LogFilePath string
	LogLevel    string
	Redis       RedisConfig
	Sender      SenderConfig
}

type onfigFileStruct struct {
	Config *Config
}


func ReadConfigFromFile(path string) (*Config, error) {
	byt, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	ptr := onfigFileStruct{
		Config: cfg,
	}
	err = toml.Unmarshal(byt, &ptr)
	if err != nil {
		return nil, err
	}
	ptr.Config = nil
	cfg.LogFilePath, err = filepath.Abs(cfg.LogFilePath)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
