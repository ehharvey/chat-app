package config

import "github.com/spf13/viper"

func InitializeConfig() {
	// Here lies all of the default values for configs

	// Database is currently unused
	// // Database: internal/infrastructure/db
	// viper.SetDefault("db.driver", "postgresql")
	// viper.SetDefault("db.username", "postgresql")
	// viper.SetDefault("db.password", "postgresql")
	// viper.SetDefault("db.host", "postgresql")
	// viper.SetDefault("db.port", 5432)
	// viper.SetDefault("db.database_name", "app")
	// viper.SetDefault("db.pool.max_conns", 10)

	// Open telemetry is currently unused
	// // opentelemetry: otel
	// viper.SetDefault("otel.general.service.namespace", "chat-app")
	// viper.SetDefault("otel.general.service.name", "web-backend")

	// // otel.trace
	// viper.SetDefault("otel.trace.grpc.enabled", true)
	// viper.SetDefault("otel.trace.grpc.endpoint", "otel")
	// viper.SetDefault("otel.trace.grpc.port", 4317)
	// viper.SetDefault("otel.trace.grpc.allow_insecure", true)
	// viper.SetDefault("otel.trace.stdout.enabled", false)
	// viper.SetDefault("otel.trace.batch_timeout_seconds", 1)

	// // otel.metrics
	// viper.SetDefault("otel.metric.grpc.enabled", true)
	// viper.SetDefault("otel.metric.grpc.endpoint", "otel")
	// viper.SetDefault("otel.metric.grpc.port", 4317)
	// viper.SetDefault("otel.metric.grpc.allow_insecure", true)
	// viper.SetDefault("otel.metric.stdout.enabled", false)
	// viper.SetDefault("otel.metric.polling_interval_seconds", 5)

	// // otel.logs
	// viper.SetDefault("otel.log.grpc.enabled", true)
	// viper.SetDefault("otel.log.grpc.endpoint", "otel")
	// viper.SetDefault("otel.log.grpc.port", 4317)
	// viper.SetDefault("otel.log.grpc.allow_insecure", true)
	// viper.SetDefault("otel.log.stdout.enabled", true)

	// Listen port
	viper.SetDefault("server.port", "8080")

	// Env
	viper.SetEnvPrefix("CHATAPP")

}

func LoadConfig() {
	viper.AutomaticEnv()
}
