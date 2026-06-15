package configs

import (
	"AuthenticationService/configs/env"
)

type serverConfig struct {
	PORT string
}

type dbConfig struct {
	DBUSER string
	DBPASS string
	DBNET  string
	DBADDR string
	DBNAME string
}

type AuthenticationConfig struct {
	TOKENSECRET string
}

type Config struct {
	Server serverConfig
	Auth   AuthenticationConfig
	DB     dbConfig
}

var config *Config

func Load() *Config {
	env.LoadEnv()
	config = &Config{
		Server: serverConfig{
			PORT: env.GetString("PORT", ":8080"),
		},
		DB: dbConfig{
			DBUSER: env.GetString("DBUSER", ""),
			DBPASS: env.GetString("DBPASS", ""),
			DBNET:  env.GetString("DBNET", "tcp"),
			DBADDR: env.GetString("DBADDR", ""),
			DBNAME: env.GetString("DBNAME", ""),
		},
		Auth: AuthenticationConfig{
			TOKENSECRET: env.GetString("TOKENSECRET", ""),
		},
	}

	return config
}

func GetConfig() *Config {
	return config
}
