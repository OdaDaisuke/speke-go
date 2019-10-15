package config

import "os"

type AppConfig struct {
	KeySeed string
}

func NewConfig() *AppConfig {
	return &AppConfig{
		KeySeed: os.Getenv("KEY_SEED"),
	}
}
