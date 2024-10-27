package config

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

type Config struct {
	Environment    string          `mapstructure:"environment"`
	Logger         *Logger         `mapstructure:"logger"`
	Server         *Server         `mapstructure:"server"`
	Infrastructure *Infrastructure `mapstructure:"infrastructure"`
}

type Logger struct {
	Format        string   `mapstructure:"format"`
	Level         string   `mapstructure:"level"`
	Output        []string `mapstructure:"output"`
	FileMaxSizeMB int      `mapstructure:"file_max_size_mb"`
	FilesKeep     int      `mapstructure:"files_keep"`
}

type Server struct {
	Mode            string `mapstructure:"mode"`
	EnableProfiling bool   `mapstructure:"enable_profiling"`
	EnableMetrics   bool   `mapstructure:"enable_metrics"`
	GPPC            *GPPC  `mapstructure:"grpc"`
	HTTP            *HTTP  `mapstructure:"http"`
	HTTPS           *HTTPS `mapstructure:"https"`
}

// insecureServer
type HTTP struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

// secureServer
type HTTPS struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	TLS  struct {
		Enabled  bool   `mapstructure:"enabled"`
		CertFile string `mapstructure:"cert_file"`
		KeyFile  string `mapstructure:"key_file"`
	} `mapstructure:"tls"`
}

type GPPC struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type Infrastructure struct {
	GORM *GORM `mapstructure:"gorm"`
}

type GORM struct {
	// Whether to enable debugging mode, default false
	Debug bool `mapstructure:"debug"`
	// Database type (currently supported database type: mysql/sqlite3/postgres), default: postgres
	DBType string `mapstructure:"db_type"`
	// Database connection string
	DSN string `mapstructure:"dsn"`
	// Set the maximum time that the connection can be reused (unit: second), default: 7200
	MaxLifetime int `mapstructure:"max_life_time"`
	// Set the maximum time that the idle can be reused (unit: second), default: 7200
	MaxIdleTime int `mapstructure:"max_idle_time"`
	// Set the maximum number of open connections for the database, default: 150
	MaxOpenConns int `mapstructure:"max_open_conns"`
	// Set the maximum number of connections in the free connection pool, default: 50
	MaxIdleConns int `mapstructure:"max_idle_conns"`
	// Database table name prefix, default: ""
	TablePrefix string `mapstructure:"table_prefix"`
	// Whether to enable automatic mapping of database table structure
	EnableAutoMigrate bool `mapstructure:"enable_auto_migrate"`
	// Whether to ignore ErrRecordNotFound
	IgnoreErrRecordNotFound bool `mapstructure:"ignore_err_record_not_found"`
	// Whether to enable parameterized queries
	ParameterizedQueries bool `mapstructure:"parameterized_queries"`
	// Whether to enable colorful logs
	Colorful bool `mapstructure:"colorful"`
	//  Whether to enable prepare statement
	PrepareStmt bool `mapstructure:"prepare_stmt"`
}

const (
	// RecommendedHomeDir defines the default directory used to place all iam service configurations.
	RecommendedHomeDir = ".eda"

	// RecommendedEnvPrefix defines the ENV prefix used by all iam service.
	RecommendedEnvPrefix = "EDA"
)

func LoadConfig(cfgFile string) (*Config, error) {
	return initConfig(cfgFile)
}

func initConfig(cfgFile string) (*Config, error) {
	logger := logger.ContextUnavailable().WithField("phase", "startup")
	if cfgFile != "" {
		logger.WithField("file", cfgFile).Info("Configuration file")
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.AddConfigPath(path.Join(getHomeDir(), RecommendedHomeDir))
		viper.AddConfigPath("/etc/eda")
	}

	viper.SetConfigType("yaml")
	viper.SetEnvPrefix(RecommendedEnvPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	// read in environment variables
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	logger = logger.WithField("file", viper.ConfigFileUsed()) // should be called after SetConfigFile
	var errFileNotFound viper.ConfigFileNotFoundError
	if err != nil && !errors.As(err, &errFileNotFound) {
		logger.WithError(err).Fatal("Failed to find a config file")
	}

	cfg, err := newConfig()
	if err != nil {
		logger.WithError(err).Fatal("Load config")
	} else {
		logger.Info("Config loaded")
	}

	return cfg, nil
}

func newConfig() (*Config, error) {
	c := &Config{}

	setDefaults()

	err := viper.Unmarshal(c)
	if err != nil {
		return nil, err
	}

	// setup logging package
	logger.SetOutputFormat(c.Logger.Format)
	if err := logger.SetOutputs(c.Logger.Output, c.Logger.FileMaxSizeMB, c.Logger.FilesKeep); err != nil {
		return nil, err
	}

	logger.SetLevel(c.Logger.Level)
	return c, nil
}

func getHomeDir() string {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println("Get home directory -", err)
		os.Exit(1)
	}
	return home
}
