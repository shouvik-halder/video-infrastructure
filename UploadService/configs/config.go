package configs

import (
	"UploadService/configs/env"
)

type serverConfig struct {
	PORT string
}

type Config struct {
	Server serverConfig
}

func Load() *Config {
	env.LoadEnv()
	return &Config{
		Server: serverConfig{
			PORT: env.GetString("PORT", ":8080"),
		},
	}
}
