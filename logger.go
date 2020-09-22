package gormrus

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
)

type GormLogger struct {
	Logger *logrus.Logger
}

func New() logger.Interface {
	return &GormLogger{
		Logger: logrus.New(),
	}
}

func NewWithLogger(logger *logrus.Logger) logger.Interface {
	return &GormLogger{
		Logger: logger,
	}
}

// LogMode log mode
func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.Logger.SetLevel(convertLevel(level))
	return l
}

// Info print info
func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	l.Logger.WithField("data", data).Info(msg)
}

// Warn print warn messages
func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.Logger.WithField("data", data).Warn(msg)
}

// Error print error messages
func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.Logger.WithField("data", data).Error(msg)
}

// Trace print sql message
func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, rows := fc()
	elapsed := time.Since(begin)

	l.Logger.WithFields(logrus.Fields{
		"sql":     sql,
		"rows":    rows,
		"elapsed": elapsed,
	}).Trace(err)
}
