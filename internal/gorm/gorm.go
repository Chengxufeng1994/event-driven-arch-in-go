package gorm

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var gormDBOnce sync.Once
var db *gorm.DB
var dialectorMap = map[string]func(string) gorm.Dialector{
	"postgres": postgres.Open,
}

func WithDebug() GormConfigOption {
	return func(gormCfg *gorm.Config) gorm.Option {
		gormCfg.Logger = gormCfg.Logger.LogMode(logger.Info)
		return gormCfg
	}
}

func NewGormDB(config *config.Infrastructure, opts ...GormConfigOption) (*gorm.DB, error) {
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  true,          // Disable color
		},
	)

	if config.GORM.Debug {
		gormLogger = gormLogger.LogMode(logger.Info)
	}

	cfg := &gorm.Config{
		Logger:                 gormLogger,
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	dialector, ok := dialectorMap[config.GORM.DBType]
	if !ok {
		return nil, fmt.Errorf("dialector %s not found", config.GORM.DBType)
	}

	var sqlDB *sql.DB
	var err error
	gormDBOnce.Do(func() {
		db, err = gorm.Open(dialector(config.GORM.DSN), cfg)
		if err != nil {
			return
		}

		sqlDB, err = db.DB()
		if err != nil {
			return
		}

		err = sqlDB.Ping()
		if err != nil {
			return
		}

		if config.GORM.MaxLifetime > 0 {
			t, _ := time.ParseDuration(fmt.Sprintf("%ds", config.GORM.MaxLifetime))
			sqlDB.SetConnMaxLifetime(t)
		}
		if config.GORM.MaxIdleTime > 0 {
			t, _ := time.ParseDuration(fmt.Sprintf("%ds", config.GORM.MaxIdleTime))
			sqlDB.SetConnMaxIdleTime(t)
		}
		if config.GORM.MaxOpenConns > 0 {
			sqlDB.SetMaxOpenConns(config.GORM.MaxOpenConns)
		}
		if config.GORM.MaxIdleConns > 0 {
			sqlDB.SetMaxIdleConns(config.GORM.MaxIdleConns)
		}
	})

	return db, err
}
