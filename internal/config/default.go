package config

import "github.com/spf13/viper"

const (
	DefaultLoggingLevel         = "INFO"
	DefaultLoggingAuditLogLevel = "DEBUG"
)

func setDefaults() {
	viper.SetDefault("logger.format", "text")
	viper.SetDefault("logger.level", DefaultLoggingLevel)
	viper.SetDefault("logger.output", "-")

	viper.SetDefault("logger.audit_log_level", DefaultLoggingAuditLogLevel)

	viper.SetDefault("logger.files_keep", 100)
	viper.SetDefault("logger.file_max_size_mb", (1<<10)*100) // 100MiB

	viper.SetDefault("server.mode", "debug")
	viper.SetDefault("server.enable_profiling", false)
	viper.SetDefault("server.enable_metrics", false)

	viper.SetDefault("server.http.host", "0.0.0.0")
	viper.SetDefault("server.http.port", 8080)

	viper.SetDefault("server.https.host", "0.0.0.0")
	viper.SetDefault("server.https.port", 8443)
	viper.SetDefault("server.https.tls.enabled", false)

	viper.SetDefault("server.grpc.host", "0.0.0.0")
	viper.SetDefault("server.grpc.port", 8081)

	viper.SetDefault("infrastructure.gorm.debug", true)
	viper.SetDefault("infrastructure.gorm.db_type", "postgres")
	viper.SetDefault("infrastructure.gorm.dsn", "host=localhost user=postgres password=postgres dbname=mallbots port=5432 sslmode=disable TimeZone=Asia/Taipei")
	viper.SetDefault("infrastructure.gorm.max_life_time", 7200)
	viper.SetDefault("infrastructure.gorm.max_idle_time", 7200)
	viper.SetDefault("infrastructure.gorm.max_open_conns", 150)
	viper.SetDefault("infrastructure.gorm.max_idle_conns", 50)
	viper.SetDefault("infrastructure.gorm.table_prefix", "")
	viper.SetDefault("infrastructure.gorm.enable_auto_migrate", false)
	viper.SetDefault("infrastructure.gorm.ignore_err_record_not_found", true)
	viper.SetDefault("infrastructure.gorm.parameterized_queries", true)
	viper.SetDefault("infrastructure.gorm.colorful", true)
	viper.SetDefault("infrastructure.gorm.prepare_stmt", true)
	viper.SetDefault("infrastructure.gorm.dry_run", true)
}
