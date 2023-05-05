package config

import (
	"github.com/spf13/viper"
)

// Config contains all configuration values.
type Config struct {
	Env string

	Host string
	Port string

	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int
}

// New returns new config instance.
func New() (*Config, error) {
	cfg := &Config{}
	v := viper.New()
	v.AutomaticEnv()
	loadFromEnv(v, cfg)

	return cfg, nil
}

func loadFromEnv(v *viper.Viper, cfg *Config) {
	env := v.Get("env")
	if env != nil {
		cfg.Env = env.(string)
	}

	host := v.Get("host")
	if host != nil {
		cfg.Host = host.(string)
	}

	port := v.Get("port")
	if port != nil {
		cfg.Port = port.(string)
	}

	redisHost := v.Get("redis_host")
	if host != nil {
		cfg.RedisHost = redisHost.(string)
	}

	redisPort := v.Get("redis_port")
	if port != nil {
		cfg.RedisPort = redisPort.(string)
	}
}
