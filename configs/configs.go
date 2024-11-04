package configs

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"time"
)

type Config struct {
	Server struct {
		AddrPort           string        `yaml:"port" env:"APP_PORT" env-default:"8283"`
		AddrHost           string        `yaml:"host" env:"APP_IP" env-default:"0.0.0.0"`
		ReadTimeout        time.Duration `yaml:"read-timeout" env:"READ_TIMEOUT" env-default:"3s"`
		WriteTimeout       time.Duration `yaml:"write-timeout" env:"WRITE_TIMEOUT" env-default:"3s"`
		IdleTimeout        time.Duration `yaml:"idle-timeout" env:"IDLE_TIMEOUT" env-default:"6s"`
		ShutdownTime       time.Duration `yaml:"shutdown-timeout" env:"SHUTDOWN_TIMEOUT" env-default:"10s"`
		CORSAllowedOrigins []string      `yaml:"cors-allowed-origins" env:"CORS_ALLOWED_ORIGINS" env-default:"localhost"`
	} `yaml:"server"`
	DataBase struct {
		ConnStr string `env:"DB_CONNECTION_STRING" env-description:"db string"`

		Host     string `yaml:"host" env:"HOST_DB" env-description:"db host"`
		User     string `yaml:"user" env:"POSTGRES_USER" env-description:"db user"`
		Password string `yaml:"password" env:"POSTGRES_PASSWORD" env-description:"db password"`
		Name     string `yaml:"db" env:"POSTGRES_DB" env-description:"db"`
		Port     string `yaml:"port" env:"PORT_DB" env-description:"port"`

		PoolMax      int           `yaml:"pool-max" env:"PG_POOL_MAX" env-default:"2"`
		ConnAttempts int           `yaml:"conn-attempts" env:"PG_CONN_ATTEMPTS" env-default:"5"`
		ConnTimeout  time.Duration `yaml:"conn-timeout" env:"PG_TIMEOUT" env-default:"2s"`
	} `yaml:"database"`
}

func NewConfig(path string) (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		return nil, err
	}
	cfg.DataBase.ConnStr = initDB(cfg)

	return &cfg, nil
}

func initDB(cfg Config) string {
	if cfg.DataBase.ConnStr != "" {
		return cfg.DataBase.ConnStr
	}
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DataBase.User,
		cfg.DataBase.Password,
		cfg.DataBase.Host,
		cfg.DataBase.Port,
		cfg.DataBase.Name,
	)
}
