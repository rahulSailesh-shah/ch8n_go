package config

import "os"

type DBConfig struct {
	Driver   string
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type ServerConfig struct {
	Port                int
	GracefulShutdownSec int
}

type AppConfig struct {
	DB       DBConfig
	Server   ServerConfig
	Auth     AuthConfig
	Polar    PolarConfig
	LogLevel string
	Env      string
}

type AuthConfig struct {
	JwksURL string
}

type PolarConfig struct {
	AccessToken string
}

func LoadConfig() (*AppConfig, error) {
	config := &AppConfig{
		DB: DBConfig{
			Driver:   "postgres",
			Host:     "localhost",
			Port:     5432,
			User:     "admin",
			Password: "admin",
			Name:     "ch8n",
		},
		Server: ServerConfig{
			Port:                8080,
			GracefulShutdownSec: 5,
		},
		Auth: AuthConfig{
			JwksURL: "http://localhost:3000/api/auth/jwks",
		},
		Polar: PolarConfig{
			AccessToken: os.Getenv("POLAR_ACCESS_TOKEN"),
		},
		LogLevel: "info",
		Env:      "development",
	}
	return config, nil
}
