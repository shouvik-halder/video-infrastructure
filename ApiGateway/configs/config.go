package configs

import "ApiGateway/configs/env"

type serverConfig struct {
	PORT string
}

type serviceConfig struct {
	AUTH_SERVICE_URL   string
	UPLOAD_SERVICE_URL string
}
type Config struct {
	Server  serverConfig
	Service serviceConfig
}

var config *Config

func Load() *Config {
	env.LoadEnv()
	config = &Config{
		Server: serverConfig{
			PORT: env.GetString("PORT", ":8000"),
		},
		Service: serviceConfig{
			AUTH_SERVICE_URL:   env.GetString("AUTH_SERVICE_URL", ""),
			UPLOAD_SERVICE_URL: env.GetString("UPLOAD_SERVICE_URL", ""),
		},
	}

	return config
}

func GetConfig() *Config {
	return config
}
