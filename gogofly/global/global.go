package global

import (
	"gogofly/conf"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Logger      *zap.SugaredLogger
	DB          *gorm.DB
	RedisClient *conf.RedisClient
)
