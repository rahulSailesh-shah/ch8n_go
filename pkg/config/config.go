package config

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
	Token    TokenConfig
	LogLevel string
	Env      string
}

type TokenConfig struct {
	AccessSecret  string
	RefreshSecret string
	AccessTTL     int // in minutes
	RefreshTTL    int // in minutes
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
		Token: TokenConfig{
			AccessSecret:  "4ciLqCmSvmtLywAsEla7RZMQpD01PqrOS2vAel2Ou/Q=",
			RefreshSecret: "cDsr0GP0OWveYQKSZ1RdPMkQUHXzgK6izs8IjyEPk60=",
			AccessTTL:     15,    // 15 minutes
			RefreshTTL:    43200, // 30 days
		},
		LogLevel: "info",
		Env:      "development",
	}
	return config, nil
}
