package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	// Logger 日志
	Logger *zap.SugaredLogger
	// MysqlDB  mysql数据库
	MysqlDB *gorm.DB
)
