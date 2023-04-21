package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

// Config contains all configuration values.
type Config struct {
	Env  string `json:"env" mapstructure:"env"`
	Host string `json:"host" mapstructure:"host"`
	Port string `json:"port" mapstructure:"port"`

	SpannerDatabase string `json:"spanner_database" mapstructure:"spanner_database"`
	DBLogEnabled    bool   `json:"db_log_enabled" mapstructure:"db_log_enabled"`

	// MFCConfigAutoRefreshTime holds time in seconds to autorefresh MFC configs.
	MFCConfigAutoRefreshTime int `json:"mfc_config_auto_refresh_time" mapstructure:"mfc_config_auto_refresh_time"`

	// BuildCommitHash contains hash for GIT commit which was used to build app.
	BuildCommitHash string
}

// New returns new config instance.
func New() (*Config, error) {
	configPath := flag.String("config-path", "", "Path to the configuration file")
	flag.Parse()

	cfg := &Config{}
	v := viper.New()
	v.AutomaticEnv()

	if *configPath == "" {
		*configPath = os.Getenv("CONFIG_PATH")
	}
	if *configPath == "" {
		loadOnlyFromEnv(v, cfg)
	} else {
		v.SetConfigType("json")
		v.SetConfigFile(*configPath)
		err := v.ReadInConfig()
		if err != nil {
			return nil, fmt.Errorf("failed to read config, err: %w", err)
		}

		err = v.Unmarshal(cfg)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal config, err: %w", err)
		}
	}

	return cfg, nil
}

func loadOnlyFromEnv(v *viper.Viper, cfg *Config) {
	v.SetTypeByDefaultValue(true)

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

	spannerDatabase := v.Get("spanner_database")
	if spannerDatabase != nil {
		cfg.SpannerDatabase = spannerDatabase.(string)
	}

	dbLogEnabled := v.Get("db_log_enabled")
	if dbLogEnabled != nil {
		cfg.DBLogEnabled = dbLogEnabled.(string) == "true"
	}

	mfcConfigAutoRefreshTime := v.Get("mfc_config_auto_refresh_time")
	if mfcConfigAutoRefreshTime != nil {
		v, err := strconv.Atoi(mfcConfigAutoRefreshTime.(string))
		if err == nil {
			cfg.MFCConfigAutoRefreshTime = v
		}
	}
}
