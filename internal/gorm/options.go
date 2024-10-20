package gorm

import "gorm.io/gorm"

type GormConfigOption func(gormConfig *gorm.Config) gorm.Option
