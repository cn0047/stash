package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	// both envs must be rewritten for non-dev envs
	ProjectID       string `envconfig:"GOOGLE_PROJECT_ID" default:"test-project"`
	SpannerInstance string `envconfig:"SPANNER_INSTANCE" default:"outbound-us"`
	SpannerDB       string `envconfig:"SPANNER_DB" default:"wp-dev"`

	Port               int           `envconfig:"PORT" default:"8080"`
	ServerReadTimeout  time.Duration `envconfig:"SERVER_READ_TIMEOUT" default:"10s"`
	ServerWriteTimeout time.Duration `envconfig:"SERVER_WRITE_TIMEOUT" default:"30s"`

	HTTPRetryNum       int           `envconfig:"HTTP_RETRY_NUM" default:"3"`
	HTTPRequestTimeout time.Duration `envconfig:"HTTP_REQUEST_TIMEOUT" default:"15s"`

	TSCURLTemplate        string `envconfig:"TSC_URL_TEMPLATE" default:"https://service-catalog-%s-%s.tom.to.com"`
	ISPSCategory          string `envconfig:"TSC_ISPS_CATEGORY" default:"isps"`
	ISPSEnabledConfigName string `envconfig:"TSC_ISPS_ENABLED_CONFIG_NAME" default:"ISPS_ENABLED"`
	AuthURLTemplate       string `envconfig:"AUTH_URL_TEMPLATE" default:"https://as-%s-%s.tom.to.com"`

	wpTopic string `envconfig:"WAVE_PLAN_TOPIC" default:"wave-plan.event"`

	TriggersGenerationHoursLimit int `envconfig:"TRIGGERS_GENERATION_HOURS_LIMIT" default:"5"`
}

func Load() (*Config, error) {
	var cfg Config

	err := envconfig.Process("wp", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, err
}
