package gormrus

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
)

func convertLevel(level logger.LogLevel) logrus.Level {
	switch level {
	case logger.Silent:
		return logrus.PanicLevel // No silent equivalent in logrus
	case logger.Error:
		return logrus.ErrorLevel
	case logger.Warn:
		return logrus.WarnLevel
	case logger.Info:
		return logrus.InfoLevel
	default:
		return logrus.InfoLevel
	}
}
