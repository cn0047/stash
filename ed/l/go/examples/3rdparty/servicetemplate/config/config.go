package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Config contains all configuration values.
type Config struct {
	Env  string `json:"env" mapstructure:"env"`
	Host string `json:"host" mapstructure:"host"`
	Port int    `json:"port" mapstructure:"port"`

	// BuildCommitHash contains hash for GIT commit which was used to build app.
	BuildCommitHash string
}

// New returns new config instance.
func New() (*Config, error) {
	configPath := flag.String("config-path", "", "Path to the configuration file")
	flag.Parse()

	if *configPath == "" {
		*configPath = os.Getenv("CONFIG_PATH")
	}
	if *configPath == "" {
		return nil, fmt.Errorf("config path cannot be empty")
	}

	v := viper.New()
	v.AutomaticEnv()
	v.SetConfigType("json")
	v.SetConfigFile(*configPath)
	err := v.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read config, err: %w", err)
	}
	cfg := &Config{}
	err = v.Unmarshal(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config, err: %w", err)
	}

	return cfg, nil
}

