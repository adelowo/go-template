package config

import (
	"github.com/caarlos0/env/v9"
)

// Config maps all dynamic app configuration values.
// This is specifically built with environment variables as the source of config
// Simply because I use Infisical for all my project needs now.
// But can always be extended to support Consul, Vault or aws Secrets manager
type Config struct {
	LogLevel string `json:"level,omitempty" env:"LOG_LEVEL"`

	OtelEndpoint string `json:"otel_endpoint,omitempty" env:"OTEL_ENDPOINT"`
	OtelUseTLS   bool   `env:"OTEL_USE_TLS" json:"otel_use_tls,omitempty"`

	PostgresDSN string `env:"POSTGRES_DSN" json:"postgres_dsn,omitempty"`

	// DISABLE IN PROD
	PostgresLogQueries bool `env:"POSTGRES_LOG_QUERIES" json:"postgres_log_queries,omitempty"`

	// Do not expose prometheus to the internet
	// Hide behind basic auth but this means your "scrapers" have to be aware of this
	PrometheusPassword string `env:"PROMETHEUS_PASSWORD" json:"prometheus_password,omitempty"`
	PrometheusUsername string `env:"PROMETHEUS_USERNAME" json:"prometheus_username,omitempty"`

	RedisDSN string `env:"REDIS_DSN" json:"redis_dsn,omitempty"`

	StripeKey           string `env:"STRIPE_KEY" json:"stripe_key,omitempty"`
	StripeWebhookSecret string `env:"STRIPE_WEBHOOK_SECRET" json:"stripe_webhook_secret,omitempty"`

	// Webhooks using Svix
	SvixAPIKey string `env:"SVIX_API_KEY" json:"svix_api_key,omitempty"`

	// AWS keys for S3 image uploading or otherwise
	AWSAccessKey string `json:"aws_access_key,omitempty" env:"AWS_ACCESS_KEY_ID"`
	AWSSecretKey string `json:"aws_secret_key,omitempty" env:"AWS_SECRET_ACCESS_KEY"`

	// domain the app is hosted
	Domain string `env:"DOMAIN" json:"domain,omitempty"`
}

func Load() (Config, error) {
	var cfg Config

	return cfg, env.Parse(&cfg)
}
